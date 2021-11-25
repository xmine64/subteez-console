package commands

import (
	"errors"
	"fmt"
	"subteez/config"
	"subteez/subteez"
)

func mainSearch(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	if len(args) < 2 {
		return errors.New("not enough arguments")
	}

	request := subteez.SearchRequest{
		Query:           args[1],
		LanguageFilters: cfg.GetLanguageFilters(),
	}
	response, err := client.Search(request)
	if err != nil {
		return err
	}

	if len(response.Result) < 1 {
		return errors.New("no any movie/series found")
	}

	for _, item := range response.Result {
		fmt.Printf("%s, %s, %d\n", item.ID, item.Name, item.Count)
	}
	return nil
}
