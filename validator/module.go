package validator

import (
	"github.com/capcom6/go-core-fx/fxutil"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"validator",
		fxutil.WithNamedLogger("validator"),
		fx.Provide(New),
	)
}
