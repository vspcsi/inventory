package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vspcsi/inventory/internal"
	"github.com/vspcsi/inventory/internal/caches"
	"github.com/vspcsi/inventory/internal/providers"
	"log"
	"net/http"
	"strings"
)

func main() {
	cache := caches.NewLocal([]internal.Provider{
		//providers.NewLinode(),
		providers.NewLocal(),
		providers.NewHetzner(),
	})

	router := gin.Default()
	router.POST("/register", func(context *gin.Context) {
		var request struct {
			Address string
			Label   string
			Tags    []string
		}

		if err := context.BindJSON(&request); err != nil {
			return
		}

		cache.Create(request.Address, request.Label, request.Tags)
		context.String(http.StatusCreated, "")
	})
	router.DELETE("/delete/:address", func(context *gin.Context) {
		address := context.Param("address")

		cache.Delete(address)
	})
	router.GET("/tags", func(context *gin.Context) {
		requestIp := context.ClientIP()
		tags := cache.GetTags(requestIp)

		context.String(200, "%s", strings.Join(tags, ","))
	})
	router.GET("/label", func(context *gin.Context) {
		requestIp := context.ClientIP()
		label := cache.GetLabel(requestIp)

		context.String(200, "%s", label)
	})

	err := router.Run("0.0.0.0:8500")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
