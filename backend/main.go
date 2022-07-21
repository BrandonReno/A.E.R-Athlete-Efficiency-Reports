package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(Module, fx.Logger(logrus.StandardLogger()))
	app.Run()
}
