package search

import (
	"flag"
	"fmt"
	"strings"
	"subteez/config"
	"subteez/errors"
	"subteez/interactive"
	"subteez/messages"
	"subteez/subteez"
)

func Main(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	query := strings.Join(flag.Args()[1:], " ")

	if cfg.IsInteractive() {
		context := interactive.Context{}
		context.Initialize(cfg)
		go context.NavigateToSearchResult(query)
		return context.Run()
	}

	if query == "" {
		return errors.ErrEmptyQuery
	}

	request := subteez.SearchRequest{
		Query:           query,
		LanguageFilters: cfg.GetLanguageFilters(),
	}
	response, err := client.Search(request)
	if err != nil {
		return err
	}

	if len(response.Result) < 1 {
		return errors.ErrNoSearchResult
	}

	for _, item := range response.Result {
		fmt.Printf(messages.SearchRow, item.ID, item.Name, item.Count)
	}
	return nil
}
