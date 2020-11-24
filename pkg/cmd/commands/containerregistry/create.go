// Copyright 2017-2020 The Usacloud Authors
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

package containerregistry

import (
	"github.com/sacloud/libsacloud/v2/helper/service/containerregistry"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	AccessLevel    string `cli:",options=container_registry_access_levels" mapconv:",filters=container_registry_access_levels_to_value" validate:"required,container_registry_access_levels"`
	SubDomainLabel string `validate:"required"`
	VirtualDomain  string `validate:"omitempty,fqdn"`

	UsersData string                    `cli:"users" mapconv:"-"`
	Users     []*containerregistry.User `cli:"-"` // --parametersでファイルからパラメータ指定する場合向け
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize() error {
	var users []*containerregistry.User
	if p.UsersData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.UsersData, &users); err != nil {
			return err
		}
	}

	p.Users = append(p.Users, users...)
	return nil
}