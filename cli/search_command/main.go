// main of "search" command

package search_command

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"subteez/config"
	"subteez/errors"
	"subteez/messages"
	"subteez/subteez"
	"subteez/tui"
)

func Main(args []string, cfg config.Config) error {
	query := strings.Join(flag.Args()[1:], " ")

	// run in interactive mode if it's enabled
	if cfg.IsInteractive() {
		context := tui.Context{}
		context.Initialize(cfg)
		go context.NavigateToSearchResult(query)
		return context.Run()
	}

	if query == "" {
		return errors.ErrEmptyQuery
	}

	if !cfg.IsScriptMode() {
		log.Printf(messages.Searching, query)
	}

	// send request
	client := subteez.NewClient(cfg.GetServer())
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

	// print result

	if cfg.IsScriptMode() {
		for _, item := range response.Result {
			fmt.Printf("%s,%s,%d\n", item.ID, item.Name, item.Count)
		}
		return nil
	}

	for _, item := range response.Result {
		id := strings.SplitAfterN(item.ID, "/", 3)[2]
		fmt.Printf(messages.SearchItem, item.Name, item.Count, id)
	}
	return nil
}
