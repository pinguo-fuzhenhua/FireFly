//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/pinguo-icc/April/internal/application"
	"github.com/pinguo-icc/April/internal/domain"
	infra "github.com/pinguo-icc/April/internal/infrastructure"
	"github.com/pinguo-icc/April/internal/infrastructure/conf"
)

func initApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		infra.ProviderSet,

		domain.ProviderSet,
		application.ProviderSet,
		newApp,
	))
}
