To install sdk

```
go get github.com/qinhong2015/bingContent
```

Init Session With RefreshToken

```
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
```

example
* [Catalog](https://github.com/qinhong2015/bingContent/blob/master/example/catalogExample.go)