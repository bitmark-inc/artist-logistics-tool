package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	bitmarksdk "github.com/bitmark-inc/bitmark-sdk-go"
	"github.com/bitmark-inc/config-loader"
	"github.com/bitmark-inc/logistics"
)

func main() {
	config.LoadConfig("LOGISTIC_SERVER")

	bitmarksdk.Init(&bitmarksdk.Config{
		Network: bitmarksdk.Network(viper.GetString("network")),
	})

	db := logistics.NewLogisticStore(viper.GetString("store.dsn"))

	router := gin.New()

	if viper.GetBool("development") {
		corsConf := cors.DefaultConfig()
		corsConf.AllowAllOrigins = true
		corsConf.AddAllowHeaders("Authorization", "Requester")
		router.Use(cors.New(corsConf))
	}

	apiRouter := router.Group("/api")
	apiRouter.GET("/owned/:requester", func(c *gin.Context) {
		requester := c.Param("requester")
		artworks, err := db.QueryOwnedArtworks(requester)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"artworks": artworks})
	})

	apiRouter.HEAD("/claim", func(c *gin.Context) {
		var query struct {
			Requester string `form:"requester" binding:"required"`
		}

		if err := c.Bind(&query); err != nil {
			logrus.WithError(err).Error("invalid request parameter")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		info, err := db.GetShipmentInformation("refik-001", query.Requester)

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
	apiRouter.POST("/claim", func(c *gin.Context) {
		requester := c.GetString("requester")

		var req struct {
			Information logistics.ShipmentInformation `json:"information"`
			Tokens      []string                      `json:"tokens"`
		}

		if err := c.Bind(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
			return
		}

		// requester = "0x37519C9BdE96f1Af2e498E3Dd3E87C1fBBAE84Be"

		count, err := db.QueryOwnedArtworkCounts(requester)
		if err != nil {
			logrus.WithError(err).Error("can not get requester data from db")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "can not get requester data from db"})
			return
		}

		requiredSelectingTokenCounts := 3 * int(count/3)
		logrus.WithField("requiredSelectingTokenCounts", requiredSelectingTokenCounts).Info("check selected token counts")
		if len(req.Tokens) != requiredSelectingTokenCounts {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid number of token selected"})
			return
		}

		ok, err := db.ValidateOwnedArtworks(requester, req.Tokens)
		if err != nil {
			logrus.WithError(err).Error("can not get requester data from db")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "can not get requester data from db"})
			return
		}

		if !ok {
			logrus.WithField("requester", requester).WithField("tokens", req.Tokens).Error("invalid NFT ownership detected")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid NFT ownership detected"})
			return
		}

		b, _ := json.Marshal(req.Tokens)

		if err := db.SaveShipmentInformation("refik-001", requester, req.Information, b); err != nil {
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

	router.Run(":8087")
}
