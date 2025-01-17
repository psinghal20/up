// Copyright 2021 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

import (
	"context"

	cp "github.com/upbound/up-sdk-go/service/controlplanes"
	"github.com/upbound/up/internal/cloud"
)

// CreateCmd creates a hosted control plane on Upbound Cloud.
type CreateCmd struct {
	Name string `arg:"" required:"" help:"Name of control plane."`

	Description string `short:"d" help:"Description for control plane."`
}

// Run executes the create command.
func (c *CreateCmd) Run(client *cp.Client, cloudCtx *cloud.Context) error {
	_, err := client.Create(context.Background(), &cp.ControlPlaneCreateParameters{
		Account:     cloudCtx.Account,
		Name:        c.Name,
		Description: c.Description,
	})
	return err
}
