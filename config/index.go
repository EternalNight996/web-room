// config/index.go
package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

var Viper *viper.Viper

func init() {
	once.Do(func() {
		Viper = viper.New()
		Viper.AddConfigPath("./")
		Viper.SetConfigName("config")
		if err := Viper.ReadInConfig(); err == nil {
			log.Println("Read config successfully: ", Viper.ConfigFileUsed())
		} else {
			log.Printf("Read failed: %s \x0A", err)
			panic(err)
		}
	})
}
