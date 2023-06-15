package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func etherueumHashedMessage(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
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

// validateTimestamp checks if a timestamp is valid and not expired
func validateTimestamp(message []byte) error {
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

	return nil
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

	if err := validateTimestamp(message); err != nil {
		return err
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

// validateERC1271Signature validates a token followed the ERC1271 spec
func validateERC1271Signature(c context.Context, walletAddress, token string) (bool, error) {
	tokenBytes, err := hex.DecodeString(token)
	if err != nil {
		return false, err
	}

	tokenBytesLen := len(tokenBytes)

	signature := tokenBytes[0 : tokenBytesLen-13]
	message := tokenBytes[tokenBytesLen-13:]

	if err := validateTimestamp(message); err != nil {
		return false, err
	}

	rpcClient, err := ethclient.Dial(viper.GetString("ethereum.rpc_url"))
	if err != nil {
		return false, err
	}

	contract, err := NewERC1271(common.HexToAddress(walletAddress), rpcClient)
	if err != nil {
		return false, err
	}

	var hash [32]byte
	copy(hash[0:32], etherueumHashedMessage(crypto.Keccak256(message)))

	r, err := contract.IsValidSignature(&bind.CallOpts{Context: c}, [32]byte(hash), signature)
	if err != nil {
		return false, err
	}

	return hex.EncodeToString(r[:]) == "1626ba7e", nil
}

// authorization is a middleware to validate requests
func authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		logisticID := c.GetHeader("LogisticID")
		// TODO: remove the default value of logistic id
		if logisticID == "" {
			logisticID = "refik-001"
		}

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
		if strings.ToLower(requester[0:2]) != "0x" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "only support ethereum address",
			})
			return
		}

		if err := validateEthereumSignature(requester, authToken); err != nil {
			logrus.WithError(err).WithField("requester", requester).WithField("authToken", authToken).Error("invalid ethereum signature")

			// fallback check ERC1271
			contractSign, err := validateERC1271Signature(c, requester, authToken)
			if err != nil {
				logrus.WithError(err).WithField("requester", requester).WithField("authToken", authToken).Error("fail to check erc1272 signature")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid signature",
				})
				return
			}

			if !contractSign {
				logrus.WithError(err).WithField("requester", requester).WithField("authToken", authToken).Error("invalid erc1271 signature")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid signature",
				})
				return
			}
		}

		logrus.WithField("requester", requester).Debug("authentication")
		c.Set("requester", requester)
		c.Set("logisticID", logisticID)

		c.Next()
	}
}
