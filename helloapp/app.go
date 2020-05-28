package helloapp

import (
	"encoding/json"
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(App{})
}

type Gadgeter struct {
}

type App struct {
	GadgetRaw json.RawMessage `json:"gadget,omitempty" caddy:"namespace=caddyworld.gadgets inline_key=gadgeter"`

	Gadget Gadgeter `json:"-"`

	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (App) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddyworld",
		New: func() caddy.Module { return new(App) },
	}
}

// Provision sets up the module.
func (a *App) Provision(ctx caddy.Context) error {
	// TODO: Provision guest modules
	a.logger = ctx.Logger(a) // a.logger is a *zap.Logger https://github.com/uber-go/zap
	a.logger.Info("Caddyworld app Provisioning.",
		zap.String("User", "kousha"),
	)
	if a.GadgetRaw != nil {
		val, err := ctx.LoadModule(a, "GadgetRaw")
		if err != nil {
			return fmt.Errorf("loading gadget module: %v", err)
		}
		a.Gadget = val.(Gadgeter)
	}

	return nil
}

func (a *App) Start() error {
	a.logger.Info("Caddyworld app start.",
		zap.String("User", "kousha"),
	)
	return nil
}

func (a *App) Stop() error {
	return nil
}

// Interface guards
var (
	_ caddy.App         = (*App)(nil)
	_ caddy.Provisioner = (*App)(nil)
)
