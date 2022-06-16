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

package server

import (
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "(*required) ")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format options: [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "Query for JSON output")
	fs.StringVarP(&p.QueryDriver, "query-driver", "", p.QueryDriver, "Name of the driver that handles queries to JSON output options: [jmespath/jq]")
	fs.StringVarP(&p.Name, "name", "", p.Name, "(*required) ")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.IntVarP(&p.CPU, "cpu", "", p.CPU, "(*required) (aliases: --core)")
	fs.IntVarP(&p.Memory, "memory", "", p.Memory, "(*required) ")
	fs.IntVarP(&p.GPU, "gpu", "", p.GPU, "")
	fs.StringVarP(&p.Commitment, "commitment", "", p.Commitment, "(*required) options: [standard/dedicatedcpu]")
	fs.StringVarP(&p.Generation, "generation", "", p.Generation, "(*required) options: [default/g100/g200]")
	fs.StringVarP(&p.InterfaceDriver, "interface-driver", "", p.InterfaceDriver, "(*required) options: [virtio/e1000]")
	fs.BoolVarP(&p.BootAfterCreate, "boot-after-create", "", p.BootAfterCreate, "")
	fs.VarP(core.NewIDFlag(&p.CDROMID, &p.CDROMID), "cdrom-id", "", "(aliases: --iso-image-id)")
	fs.VarP(core.NewIDFlag(&p.PrivateHostID, &p.PrivateHostID), "private-host-id", "", "")
	fs.StringVarP(&p.NetworkInterface.Upstream, "network-interface-upstream", "", p.NetworkInterface.Upstream, "options: [shared/disconnected/(switch-id)]")
	fs.VarP(core.NewIDFlag(&p.NetworkInterface.PacketFilterID, &p.NetworkInterface.PacketFilterID), "network-interface-packet-filter-id", "", "")
	fs.StringVarP(&p.NetworkInterface.UserIPAddress, "network-interface-user-ip-address", "", p.NetworkInterface.UserIPAddress, "")
	fs.StringVarP(&p.NetworkInterfaceData, "network-interfaces", "", p.NetworkInterfaceData, "")
	fs.VarP(core.NewIDFlag(&p.Disk.ID, &p.Disk.ID), "disk-id", "", "")
	fs.StringVarP(&p.Disk.Name, "disk-name", "", p.Disk.Name, "")
	fs.StringVarP(&p.Disk.Description, "disk-description", "", p.Disk.Description, "")
	fs.StringSliceVarP(&p.Disk.Tags, "disk-tags", "", p.Disk.Tags, "")
	fs.VarP(core.NewIDFlag(&p.Disk.IconID, &p.Disk.IconID), "disk-icon-id", "", "")
	fs.StringVarP(&p.Disk.DiskPlan, "disk-disk-plan", "", p.Disk.DiskPlan, "options: [ssd/hdd]")
	fs.StringVarP(&p.Disk.Connection, "disk-connection", "", p.Disk.Connection, "options: [virtio/ide]")
	fs.VarP(core.NewIDFlag(&p.Disk.SourceDiskID, &p.Disk.SourceDiskID), "disk-source-disk-id", "", "")
	fs.VarP(core.NewIDFlag(&p.Disk.SourceArchiveID, &p.Disk.SourceArchiveID), "disk-source-archive-id", "", "")
	fs.IntVarP(&p.Disk.SizeGB, "disk-size", "", p.Disk.SizeGB, "(aliases: --size-gb)")
	fs.VarP(core.NewIDSliceFlag(&p.Disk.DistantFrom, &p.Disk.DistantFrom), "disk-distant-from", "", "")
	fs.StringVarP(&p.Disk.OSType, "disk-os-type", "", p.Disk.OSType, "options: [almalinux/rockylinux/miraclelinux/centos8stream/ubuntu/debian/rancheros/k3os/...]")
	fs.StringVarP(&p.Disk.EditDisk.HostName, "disk-edit-host-name", "", p.Disk.EditDisk.HostName, "")
	fs.StringVarP(&p.Disk.EditDisk.Password, "disk-edit-password", "", p.Disk.EditDisk.Password, "")
	fs.StringVarP(&p.Disk.EditDisk.IPAddress, "disk-edit-ip-address", "", p.Disk.EditDisk.IPAddress, "")
	fs.IntVarP(&p.Disk.EditDisk.NetworkMaskLen, "disk-edit-netmask", "", p.Disk.EditDisk.NetworkMaskLen, "(aliases: --network-mask-len)")
	fs.StringVarP(&p.Disk.EditDisk.DefaultRoute, "disk-edit-gateway", "", p.Disk.EditDisk.DefaultRoute, "(aliases: --default-route)")
	fs.BoolVarP(&p.Disk.EditDisk.DisablePWAuth, "disk-edit-disable-pw-auth", "", p.Disk.EditDisk.DisablePWAuth, "")
	fs.BoolVarP(&p.Disk.EditDisk.EnableDHCP, "disk-edit-enable-dhcp", "", p.Disk.EditDisk.EnableDHCP, "")
	fs.BoolVarP(&p.Disk.EditDisk.ChangePartitionUUID, "disk-edit-change-partition-uuid", "", p.Disk.EditDisk.ChangePartitionUUID, "")
	fs.StringSliceVarP(&p.Disk.EditDisk.SSHKeys, "disk-edit-ssh-keys", "", p.Disk.EditDisk.SSHKeys, "")
	fs.VarP(core.NewIDSliceFlag(&p.Disk.EditDisk.SSHKeyIDs, &p.Disk.EditDisk.SSHKeyIDs), "disk-edit-ssh-key-ids", "", "")
	fs.BoolVarP(&p.Disk.EditDisk.IsSSHKeysEphemeral, "disk-edit-make-ssh-keys-ephemeral", "", p.Disk.EditDisk.IsSSHKeysEphemeral, "")
	fs.VarP(core.NewIDSliceFlag(&p.Disk.EditDisk.NoteIDs, &p.Disk.EditDisk.NoteIDs), "disk-edit-note-ids", "", "")
	fs.StringVarP(&p.Disk.EditDisk.NotesData, "disk-edit-notes", "", p.Disk.EditDisk.NotesData, "")
	fs.BoolVarP(&p.Disk.EditDisk.IsNotesEphemeral, "disk-edit-make-notes-ephemeral", "", p.Disk.EditDisk.IsNotesEphemeral, "")
	fs.BoolVarP(&p.Disk.NoWait, "disk-no-wait", "", p.Disk.NoWait, "")
	fs.VarP(core.NewIDSliceFlag(&p.DiskIDs, &p.DiskIDs), "disk-ids", "", "")
	fs.StringVarP(&p.DisksData, "disks", "", p.DisksData, "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
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
	case "core":
		name = "cpu"
	case "iso-image-id":
		name = "cdrom-id"
	case "size-gb":
		name = "disk-size"
	case "network-mask-len":
		name = "disk-edit-netmask"
	case "default-route":
		name = "disk-edit-gateway"
	}
	return pflag.NormalizedName(name)
}

func (p *createParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-icon-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Common options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("plan", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("cpu"))
		fs.AddFlag(cmd.LocalFlags().Lookup("memory"))
		fs.AddFlag(cmd.LocalFlags().Lookup("commitment"))
		fs.AddFlag(cmd.LocalFlags().Lookup("gpu"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generation"))
		sets = append(sets, &core.FlagSet{
			Title: "Plan options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("server", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("boot-after-create"))
		fs.AddFlag(cmd.LocalFlags().Lookup("cdrom-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("interface-driver"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-host-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Server-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("disk", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-connection"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-disk-plan"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-distant-from"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-no-wait"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-os-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-size"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-source-archive-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-source-disk-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disks"))
		sets = append(sets, &core.FlagSet{
			Title: "Disk options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("diskedit", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-host-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-password"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-netmask"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-gateway"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-disable-pw-auth"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-enable-dhcp"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-change-partition-uuid"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-ssh-keys"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-ssh-key-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-make-ssh-keys-ephemeral"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-note-ids"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-notes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("disk-edit-make-notes-ephemeral"))
		sets = append(sets, &core.FlagSet{
			Title: "Edit disk options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("network", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("network-interface-packet-filter-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("network-interface-upstream"))
		fs.AddFlag(cmd.LocalFlags().Lookup("network-interface-user-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("network-interfaces"))
		sets = append(sets, &core.FlagSet{
			Title: "Network options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("zone", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		sets = append(sets, &core.FlagSet{
			Title: "Zone options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("wait", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("no-wait"))
		sets = append(sets, &core.FlagSet{
			Title: "Wait options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("input", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-driver"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("example", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("example"))
		sets = append(sets, &core.FlagSet{
			Title: "Parameter example",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("commitment", util.FlagCompletionFunc("standard", "dedicatedcpu"))
	cmd.RegisterFlagCompletionFunc("generation", util.FlagCompletionFunc("default", "g100", "g200"))
	cmd.RegisterFlagCompletionFunc("interface-driver", util.FlagCompletionFunc("virtio", "e1000"))
	cmd.RegisterFlagCompletionFunc("network-interface-upstream", util.FlagCompletionFunc("shared", "disconnected"))
	cmd.RegisterFlagCompletionFunc("disk-disk-plan", util.FlagCompletionFunc("ssd", "hdd"))
	cmd.RegisterFlagCompletionFunc("disk-connection", util.FlagCompletionFunc("virtio", "ide"))
	cmd.RegisterFlagCompletionFunc("disk-os-type", util.FlagCompletionFunc("centos", "centos8stream", "centos7", "almalinux", "rockylinux", "miracle", "miraclelinux", "ubuntu", "ubuntu2204", "ubuntu2004", "ubuntu1804", "debian", "debian10", "debian11", "rancheros", "k3os", "kusanagi"))

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}