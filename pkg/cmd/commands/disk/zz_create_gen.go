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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.StringVarP(&p.Name, "name", "", p.Name, "")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.StringVarP(&p.DiskPlan, "disk-plan", "", p.DiskPlan, "options: [ssd/hdd]")
	fs.StringVarP(&p.Connection, "connection", "", p.Connection, "options: [virtio/ide]")
	fs.VarP(core.NewIDFlag(&p.SourceDiskID, &p.SourceDiskID), "source-disk-id", "", "")
	fs.VarP(core.NewIDFlag(&p.SourceArchiveID, &p.SourceArchiveID), "source-archive-id", "", "")
	fs.VarP(core.NewIDFlag(&p.ServerID, &p.ServerID), "server-id", "", "")
	fs.IntVarP(&p.SizeGB, "size", "", p.SizeGB, "")
	fs.VarP(core.NewIDSliceFlag(&p.DistantFrom, &p.DistantFrom), "distant-from", "", "")
	fs.StringVarP(&p.OSType, "os-type", "", p.OSType, "options: [centos/centos8/centos7/ubuntu/ubuntu2004/ubuntu1804/ubuntu1604/debian/debian10/debian9/coreos/rancheros/k3os/kusanagi/freebsd/windows2016/windows2016-rds/windows2016-rds-office/windows2016-sql-web/windows2016-sql-standard/windows2016-sql-standard-all/windows2016-sql2017-standard/windows2016-sql2017-enterprise/windows2016-sql2017-standard-all/windows2019/windows2019-rds/windows2019-rds-office2019/windows2019-sql2017-web/windows2019-sql2019-web/windows2019-sql2017-standard/windows2019-sql2019-standard/windows2019-sql2017-enterprise/windows2019-sql2019-enterprise/windows2019-sql2017-standard-all/windows2019-sql2019-standard-all]")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.FormatFile, "format-file", "", p.FormatFile, "Output format in Go templates(from file)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.QueryFile, "query-file", "", p.QueryFile, "JMESPath query(from file)")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *createParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *createParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-plan"))
		fs.AddFlag(cmd.LocalFlags().Lookup("connection"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-archive-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("server-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("size"))
		fs.AddFlag(cmd.LocalFlags().Lookup("distant-from"))
		fs.AddFlag(cmd.LocalFlags().Lookup("os-type"))
		sets = append(sets, &core.FlagSet{
			Title: "Disk options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
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
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("disk-plan", util.FlagCompletionFunc("ssd", "hdd"))
	cmd.RegisterFlagCompletionFunc("connection", util.FlagCompletionFunc("virtio", "ide"))
	cmd.RegisterFlagCompletionFunc("os-type", util.FlagCompletionFunc("centos", "centos8", "centos7", "ubuntu", "ubuntu2004", "ubuntu1804", "ubuntu1604", "debian", "debian10", "debian9", "coreos", "rancheros", "k3os", "kusanagi", "freebsd", "windows2016", "windows2016-rds", "windows2016-rds-office", "windows2016-sql-web", "windows2016-sql-standard", "windows2016-sql-standard-all", "windows2016-sql2017-standard", "windows2016-sql2017-enterprise", "windows2016-sql2017-standard-all", "windows2019", "windows2019-rds", "windows2019-rds-office2019", "windows2019-sql2017-web", "windows2019-sql2019-web", "windows2019-sql2017-standard", "windows2019-sql2019-standard", "windows2019-sql2017-enterprise", "windows2019-sql2019-enterprise", "windows2019-sql2017-standard-all", "windows2019-sql2019-standard-all"))

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
