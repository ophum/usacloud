package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func DiskReinstallFromDisk(ctx Context, params *ReinstallFromDiskDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallFromDisk is failed: %s", e)
	}

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewSpinner(
		"Installing...",
		"Reinstall disk from disk is complete.\n",
		internal.CharSetProgress,
		GlobalOption.Progress)
	spinner.Start()

	go func() {
		_, err := api.ReinstallFromDisk(params.Id, params.SourceDiskId, params.DistantFrom...)
		if err != nil {
			errChan <- err
			return
		}
		err = api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			errChan <- err
			return
		}
		compChan <- true
	}()

copy:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break copy
		case err := <-errChan:
			return fmt.Errorf("DiskReinstallFromDisk is failed: %s", err)
		}
	}

	return nil

}
