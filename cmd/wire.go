//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/AllPaste/paste/config"
	"github.com/AllPaste/paste/internal/app"
	"github.com/AllPaste/paste/internal/data/repository"
	"github.com/AllPaste/paste/internal/driver"
	"github.com/AllPaste/paste/internal/service"
	"github.com/AllPaste/paste/pkg/log"
	"github.com/google/wire"
)

// InitializeEvent 声明injector的函数签名
func server(c *config.Config, logger *log.Logger) (*app.Server, func(), error) {
	panic(
		wire.Build(
			driver.ProviderSet,
			repository.ProviderSet,
			service.ProviderSet,
			app.ProviderSet,
		),
	)
}
