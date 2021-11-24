package commands

import (
	"fmt"
	"log"
	"subteez/config"
	"subteez/subteez"
)

func DownloadMain(cfg config.Config, id string) {
	client := subteez.NewClient(cfg.GetServer())

	request := subteez.SubtitleDownloadRequest{
		ID: id,
	}
	response, err := client.Download(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes downloaded.\n", len(response))
}
