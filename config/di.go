package config

import (
	"github.com/golobby/container/v3"
	"payervice-http-endpoint/application"
	"payervice-http-endpoint/application/contracts"
)

func BuildDI() error {
	if err := container.Singleton(func() contracts.PayService {
		return application.NewStupidPayService()
	}); err != nil {
		return err
	}

	return nil
}
