package commands

import (
	"errors"
	"fmt"
	"subteez/config"
	"subteez/subteez"
)

func mainDownload(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	if len(args) < 2 {
		return errors.New("not enough arguments")
	}

	request := subteez.SubtitleDownloadRequest{
		ID: args[1],
	}
	response, err := client.Download(request)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes downloaded.\n", len(response))
	return errors.New("not implemented, yet")
}
