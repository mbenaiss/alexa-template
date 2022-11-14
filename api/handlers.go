package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HelloWorldIntent = "HelloWorldIntent"
)

func dispatchIntent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request AlexaRequest

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusOK, SimpleResponse("une erreur est survenue", true))

			return
		}

		if GetIntentType(request) == IntentTypeLaunch {
			c.JSON(http.StatusOK, SimpleResponse("Bienvenue ", false))

			return
		}

		switch GetIntentName(request) {
		case HelloWorldIntent:
			c.JSON(http.StatusOK, SimpleResponse("Hello !", true))

			return
		case StopIntent:
			c.JSON(http.StatusOK, SimpleResponse("Goodbye", true))

			return
		default:
			c.JSON(http.StatusOK, SimpleResponse("Bonjour", false))

			return
		}
	}
}
