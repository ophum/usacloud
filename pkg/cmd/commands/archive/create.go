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

package archive

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/validate"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
	ValidateFunc: validateCreateParameter,
}

type createParameter struct {
	cflag.ZoneParameter `cli:",squash" mapconv:",squash"`

	Name        string   `cli:",category=archive" validate:"required"`
	Description string   `cli:",category=archive" validate:"description"`
	Tags        []string `cli:",category=archive" validate:"tags"`
	IconID      types.ID `cli:",category=archive"`
	SizeGB      int      `cli:"size,category=archive" validate:"required_with=SourceFile"`

	SourceFile      string   `cli:",category=archive" mapconv:"SourceReader,filters=path_to_reader" validate:"omitempty,file"` // TODO 標準入力(パイプも)への対応
	SourceDiskID    types.ID `cli:",category=archive"`
	SourceArchiveID types.ID `cli:",category=archive"`

	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func validateCreateParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*createParameter)

	var errs []error
	// conflict with
	targets := []*validate.Target{
		{FlagName: "--source-file", Value: p.SourceFile},
		{FlagName: "--source-archive-id", Value: p.SourceArchiveID},
		{FlagName: "--source-disk-id", Value: p.SourceDiskID},
	}
	if err := validate.ConflictWith(targets...); err != nil {
		errs = append(errs, err)
	}

	return validate.NewValidationError(errs...)
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}