package config

import (
	"github.com/golobby/container/v3"
	"http-endpoint/application"
	"http-endpoint/application/contracts"
)

func BuildDI() error {
	if err := container.Singleton(func() contracts.PayService {
		return application.NewStupidPayService()
	}); err != nil {
		return err
	}

	return nil
}
