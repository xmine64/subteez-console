package download

import (
	"errors"
	"fmt"
	"subteez/config"
	"subteez/messages"
	"subteez/subteez"
)

func Main(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	if len(args) < 2 {
		return errors.New(messages.NotEnoughArguments)
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
