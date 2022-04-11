// Copyright 2017-2022 The Usacloud Authors
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

package config

import (
	"github.com/sacloud/usacloud/pkg/commands/iaas"
	"github.com/sacloud/usacloud/pkg/core"
)

var Resource = &core.Resource{
	Name:               "config",
	Usage:              "Management commands for Configuration file/Profile",
	Aliases:            []string{"profile"},
	DefaultCommandName: "edit",
	Category:           iaas.ResourceCategoryConfig,
	IsGlobalResource:   true,
	SkipLoadingProfile: true,
}
