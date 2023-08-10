// Package server provides a server for the AeroSquirrel API.
package server

import (
	"context"
	"fmt"

	v1 "github.com/oidc-soma/aerosquirrel/api/v1"
)

// Exec executes the server.
func Exec(address string, port int) error {
	err := runServer(address, port)
	if err != nil {
		return err
	}

	return nil
}

func runServer(address string, port int) error {
	router := v1.Endpoints(context.Background())

	if err := router.Run(fmt.Sprintf("%s:%d", address, port)); err != nil {
		return err
	}

	return nil
}
