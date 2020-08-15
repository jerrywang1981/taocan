// Package main provides ...
package main

import (
	"log"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/gin-gonic/gin"
	"github.com/jerrywang1981/taocan/routers"
	"github.com/jerrywang1981/watson/assistant"
	"github.com/joho/godotenv"
)

var configData *assistant.WAConfig = &assistant.WAConfig{}

func loadConfig() {
	// load from default env first
	configData.ApiKey = os.Getenv("WA_API_KEY")
	configData.ApiUrl = os.Getenv("WA_API_URL")
	configData.AssistantId = os.Getenv("WA_ASSISTANT_ID")
	configData.Version = os.Getenv("WA_VERSION")
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file does not exist")
	} else {
		configData.ApiKey = os.Getenv("WA_API_KEY")
		configData.ApiUrl = os.Getenv("WA_API_URL")
		configData.AssistantId = os.Getenv("WA_ASSISTANT_ID")
		configData.Version = os.Getenv("WA_VERSION")
	}

	appEnv, _ := cfenv.Current()

	if appEnv != nil {
		waService, _ := appEnv.Services.WithLabel("conversation")
		if len(waService) > 0 {
			configData.ApiKey = waService[0].Credentials["apikey"].(string)
			configData.ApiUrl = waService[0].Credentials["url"].(string)
		}
	}
}

func main() {
	loadConfig()
	r := gin.Default()

	// connect to WA
	bot := assistant.ConnectWA(configData)
	defer bot.Close()

	v1 := r.Group("/api/v1")
	{
		routers.LoadWATaoCan(v1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //Local
	}
	r.Run(":" + port)
}
