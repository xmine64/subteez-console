package commands

import (
	"fmt"
	"log"
	"subteez/config"
	"subteez/subteez"
)

func SearchMain(cfg config.Config, query string) {
	client := subteez.NewClient(cfg.GetServer())

	request := subteez.SearchRequest{
		Query:           query,
		LanguageFilters: cfg.GetLanguageFilters(),
	}
	response, err := client.Search(request)
	if err != nil {
		log.Fatal(err)
	}

	if len(response.Result) < 1 {
		fmt.Println("No movies/series found.")
		return
	}

	for _, item := range response.Result {
		fmt.Printf("%s, %s, %d\n", item.ID, item.Name, item.Count)
	}
}
