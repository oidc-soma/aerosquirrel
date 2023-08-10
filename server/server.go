/*
 * Copyright 2023 The AeroSquirrel Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
