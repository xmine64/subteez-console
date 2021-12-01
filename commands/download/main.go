package download

import (
	"fmt"
	"subteez/config"
	"subteez/errors"
	"subteez/subteez"
)

func Main(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	if len(args) < 2 {
		return errors.ErrNotEnoughArguments
	}

	request := subteez.SubtitleDownloadRequest{
		ID: args[1],
	}
	response, err := client.Download(request)
	if err != nil {
		return err
	}

	return fmt.Errorf("file saving not implemented but %d bytes downloaded", len(response))
}
