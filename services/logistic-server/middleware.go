package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/sha3"

	"github.com/bitmark-inc/bitmark-sdk-go/account"
)

func etherueumHashedMessage(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func bitmarkHashedMessage(messages string) []byte {
	hash := sha3.New256()
	hash.Write([]byte(fmt.Sprintf("\x01\xe9Bitmark Signed Message:\n%d%s", len(messages), messages)))
	return hash.Sum(nil)
}

func ecrecover(message, signature []byte) (string, error) {
	hash := etherueumHashedMessage(crypto.Keccak256(message))

	if len(signature) != 65 {
		return "", fmt.Errorf("invalid length of signature")
	}

	// Version of signature should be 27 or 28, but 0 and 1 are also possible versions
	if signature[64] >= 27 {
		signature[64] -= 27
	}

	pubKey, err := crypto.SigToPub(hash, signature[:])
	if err != nil {
		return "", err
	}

	ethAddress := crypto.PubkeyToAddress(*pubKey)
	return ethAddress.String(), nil
}

// validateEthereumSignature validate the token by requester
// and check the timestamp
// TODO: figure out how to deal with the token attached message in an abstract way
func validateEthereumSignature(requester, token string) error {
	tokenBytes, err := hex.DecodeString(token)
	if err != nil {
		return err
	}

	signature := tokenBytes[0:65]
	message := tokenBytes[65:]

	t, err := strconv.Atoi(string(message))
	if err != nil {
		return fmt.Errorf("fail to parse timestamp. error: %s", err.Error())
	}
	sec := t / 1000
	nsec := t & 1000
	reqTime := time.Unix(int64(sec), int64(nsec))

	if time.Since(reqTime) > 5*time.Minute {
		return fmt.Errorf("token expired")
	}

	addr, err := ecrecover(message, signature)
	if err != nil {
		return err
	}
	if addr != requester {
		return fmt.Errorf("signature is mismatched with the requester address")
	}

	return nil
}

// validateBitmarkSignature validate a token is signed by the requester
// and check the message is a specific string
// TODO: figure out how to deal with the token attached message in an abstract way
func validateBitmarkSignature(requester, token string) error {
	tokenBytes, err := hex.DecodeString(token)
	if err != nil {
		return err
	}

	signature := tokenBytes[0:64]

	message := string(tokenBytes[64:])
	if message != viper.GetString("feralfile.preset_message") {
		return fmt.Errorf("invalid payload")
	}

	return account.Verify(requester, bitmarkHashedMessage(message), signature)
}

// authorization is a middleware to validate requests
func authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.GetHeader("Requester")
		authStrings := strings.Split(c.GetHeader("Authorization"), " ")

		if len(authStrings) != 2 && authStrings[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization information",
			})
			return
		}

		authToken := authStrings[1]
		// requester starts with "0x" would be treated as an ethereum account
		if requester[0:2] == "0x" {
			if err := validateEthereumSignature(requester, authToken); err != nil {
				logrus.WithError(err).WithField("requester", requester).WithField("authToken", authToken).Error("invalid ethereum signature")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid signature",
				})
				return
			}
		} else {
			if err := validateBitmarkSignature(requester, authToken); err != nil {
				logrus.WithError(err).WithField("requester", requester).WithField("authToken", authToken).Error("invalid bitmark signature")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid signature",
				})
				return
			}
		}
		logrus.WithField("requester", requester).Debug("authentication")
		c.Set("requester", requester)

		c.Next()
	}
}
