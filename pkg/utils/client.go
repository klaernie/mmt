package utils

import (
	"github.com/spf13/viper"
	"github.com/hashicorp/go-retryablehttp"
)

func timeoutFromConfig() int {
	key := "network_timeout"
	viper.SetDefault(key, 4)
	return viper.GetInt(key)
}

var Client = retryablehttp.NewClient().StandardClient()
