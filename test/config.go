package test

import (
	"os"

	"github.com/plevym/sdk-go/pkg/config"
)

func Config() *config.Config {
	cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}
	return cfg
}
