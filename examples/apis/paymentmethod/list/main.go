package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	paymentMethodClient := paymentmethod.NewClient(cfg)
	paymentMethods, err := paymentMethodClient.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range paymentMethods {
		fmt.Println(v)
	}
}
