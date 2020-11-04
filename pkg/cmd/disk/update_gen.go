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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-params-v1'; DO NOT EDIT

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/base"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *UpdateParameter) BuildFlags(fs *pflag.FlagSet) {
	fs.StringVarP(p.Name, "name", "", *p.Name, "")
	fs.StringVarP(p.Description, "description", "", *p.Description, "")
	fs.StringSliceVarP(p.Tags, "tags", "", *p.Tags, "")
	fs.VarP(base.NewIDFlag(p.IconID, p.IconID), "icon-id", "", "")
	fs.StringVarP(p.Connection, "connection", "", *p.Connection, "")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: )")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: )")
	fs.StringVarP(&p.FormatFile, "format-file", "", p.FormatFile, "Output format in Go templates(from file)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.QueryFile, "query-file", "", p.QueryFile, "JMESPath query(from file)")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *UpdateParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *UpdateParameter) CategorizedFlagSets(cmd *cobra.Command) []*base.FlagSet {
	var sets []*base.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("connection"))
		sets = append(sets, &base.FlagSet{
			Title: "Common options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &base.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}

	return sets
}
