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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package services

import (
	service "github.com/sacloud/libsacloud/v2/helper/service/serverplan"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/conv"
)

func init() {
	setDefaultServiceFunc("server-plan", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.FindRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.FindWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil

		},
	)
	setDefaultListAllFunc("server-plan", "list",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
	setDefaultServiceFunc("server-plan", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())

			req := &service.ReadRequest{}
			if err := conv.ConvertTo(parameter, req); err != nil {
				return nil, err
			}
			if err := req.Validate(); err != nil {
				return nil, err
			}

			res, err := svc.ReadWithContext(ctx, req)
			if err != nil {
				return nil, err
			}

			return []interface{}{res}, nil

		},
	)
	setDefaultListAllFunc("server-plan", "read",
		func(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
			svc := service.New(ctx.Client())
			res, err := svc.FindWithContext(ctx, &service.FindRequest{Zone: (parameter.(cflag.ZoneParameterValueHandler)).ZoneFlagValue()})
			if err != nil {
				return nil, err
			}

			var results []interface{}
			for _, v := range res {
				results = append(results, v)
			}
			return results, nil
		},
	)
}
