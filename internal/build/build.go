// +build packaging

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

package build

// NOTE(hasheddan): See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

//go:generate go run github.com/goreleaser/nfpm/v2/cmd/nfpm pkg --config $CACHE_DIR/nfpm.yaml --packager $PACKAGER --target $OUTPUT_DIR/$PACKAGER/$PLATFORM/up.$PACKAGER

import (
	_ "github.com/goreleaser/nfpm/v2/cmd/nfpm" //nolint:typecheck
	_ "github.com/spf13/cobra/doc"             //nolint:typecheck
)
