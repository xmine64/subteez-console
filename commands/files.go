package commands

import (
	"fmt"
	"log"
	"subteez/config"
	"subteez/subteez"
)

func FilesMain(cfg config.Config, id string) {
	client := subteez.NewClient(cfg.GetServer())

	request := subteez.SubtitleDetailsRequest{
		ID:              id,
		LanguageFilters: cfg.GetLanguageFilters(),
	}
	response, err := client.GetDetails(request)
	if err != nil {
		log.Fatal(err)
	}

	if len(response.Files) < 1 {
		fmt.Println("No any file found.")
		return
	}

	for _, item := range response.Files {
		fmt.Printf("%s, %s, %s, %s\n", item.ID, item.Language.GetTitle(), item.Title, item.Author)
	}
}
