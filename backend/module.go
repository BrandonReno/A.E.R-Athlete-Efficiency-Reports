package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BrandonReno/A.E.R/config"
	"github.com/BrandonReno/A.E.R/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewConfig,
		config.NewDatabaseClient,
		config.NewHTTPClient,
		NewRouter,
	),
	fx.Invoke(
		RunServer,
		MountRoutes,
	),
)

type MountRouteParams struct {
	fx.In
	Router         chi.Router
	WorkoutHandler *handler.WorkoutHandler
}

func MountRoutes(params MountRouteParams) error {
	params.WorkoutHandler.MountRoutes(params.Router)
	return nil
}

func NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Mount("/debug", middleware.Profiler())
	return router
}

func RunServer(lc fx.Lifecycle, cfg *config.Config, router chi.Router) error {
	srv := http.Server{}
	srv.Addr = fmt.Sprintf(":%d", cfg.ServerPort)
	srv.Handler = router
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil {
					logrus.Error("server can not listen and serve")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return nil
}
