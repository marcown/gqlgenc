package main

import (
	"context"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/marcown/gqlgenc/clientgen"
	"github.com/marcown/gqlgenc/config"
	"github.com/marcown/gqlgenc/generator"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cfg: %+v", err.Error())
		os.Exit(2)
	}

	clientGen := api.AddPlugin(clientgen.New(cfg))

	if err := generator.Generate(ctx, cfg, clientGen); err != nil {
		fmt.Fprintf(os.Stderr, "generate: %+v", err.Error())
		os.Exit(4)
	}
}
