package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/bitmark-inc/config-loader"
	"github.com/bitmark-inc/logistics"
)

func main() {
	config.LoadConfig("LOGISTIC_SERVER")

	db := logistics.NewLogisticStore(viper.GetString("store.dsn"))

	router := gin.New()

	if viper.GetBool("development") {
		corsConf := cors.DefaultConfig()
		corsConf.AllowAllOrigins = true
		corsConf.AddAllowHeaders("Authorization", "Requester")
		router.Use(cors.New(corsConf))
	}

	apiRouter := router.Group("/api")

	// validate if an address has already submitted for a logistic campaign before
	apiRouter.HEAD("/claim/:address", func(c *gin.Context) {
		logisticID := c.GetHeader("LogisticID")
		claimAddress := c.Param("address")

		info, err := db.GetShipmentInformation(logisticID, claimAddress)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusNotFound)
				return
			}
			logrus.WithError(err).Error("fail to get shipment information")
			c.Status(http.StatusInternalServerError)
			return
		}

		if info.OwnerAddress == "" {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusNoContent)
	})

	apiRouter.Use(authorization())
	// submits the shipmenet information of an owner if he is qualified
	apiRouter.POST("/claim", func(c *gin.Context) {
		requester := c.GetString("requester")
		logisticID := c.GetHeader("LogisticID")
		if logisticID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "LogisticID from header is required"})
			return
		}

		var req struct {
			Information logistics.ShipmentInformation `json:"information"`
			TokenID     string                        `json:"tokenID"`
		}

		if err := c.Bind(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
			return
		}

		// validate ownership here. Here is an example of validate from a db snapshot
		// ok, err := db.ValidateOwnedArtworks(requester, logisticID, []string{req.TokenID})
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "fail to validate token ownership"})
		// 	return
		// }
		// if !ok {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "token not belong to this user"})
		// 	return
		// }

		// TODO: or validate the ownership from infura or other sources

		if err := db.SaveShipmentInformation(logisticID, requester, req.Information, nil); err != nil {
			pqErr := err.(*pgconn.PgError)
			if pqErr.Code == "23505" {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "information has already submitted"})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "fail to upload shipment information"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": 1})
	})

	router.Use(static.Serve("/", static.LocalFile(viper.GetString("ui_path"), false)))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "this is not what you are looking for",
		})
	})

	if err := router.Run(":8087"); err != nil {
		logrus.WithError(err).Panic("server stopped with error")
	}
}
