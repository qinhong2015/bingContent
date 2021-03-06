package example

import (
	"github.com/qinhong2015/bingContent"
	bingClient "github.com/qinhong2015/bingContent/client"
)

func InitSessionWithRefreshToken() *bingContent.Session {
	bingConfig := bingClient.NewConfig()
	bingConfig.ClientId = "clientIdString"
	bingConfig.ClientSecret = "clientSecretString"
	bingConfig.RequireLiveConnect = false
	bingConfig.MerchantId = "MerchantId"
	bingConfig.RefreshToken = "RefreshToken"
	bingConfig.CustomerId = "CustomerId" //optional, required by manager account
	bingConfig.AdsId = "AdsId" //optional, required by manager account
	bingConfig.DeveloperToken = "developerToken"

	bingSession := bingContent.NewSession(bingConfig)
	return bingSession
}
