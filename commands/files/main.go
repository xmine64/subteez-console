package files

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"subteez/config"
	"subteez/errors"
	"subteez/interactive"
	"subteez/messages"
	"subteez/subteez"
)

func Main(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	id := strings.Join(flag.Args()[1:], " ")
	if !strings.HasPrefix(id, "/subtitles/") {
		id = "/subtitles/" + id
	}

	if id == "" {
		return errors.ErrEmptyID
	}

	if cfg.IsInteractive() {
		context := interactive.Context{}
		context.Initialize(cfg)
		go context.NavigateToDetails(id)
		return context.Run()
	}

	if !cfg.IsScriptMode() {
		log.Printf(messages.GettingPage, id)
	}

	request := subteez.SubtitleDetailsRequest{
		ID:              id,
		LanguageFilters: cfg.GetLanguageFilters(),
	}
	response, err := client.GetDetails(request)
	if err != nil {
		return err
	}

	if len(response.Files) < 1 {
		return errors.ErrNoFileFound
	}

	if cfg.IsScriptMode() {
		for _, item := range response.Files {
			fmt.Printf("%s,%s,%s,%s\n", item.ID, item.Language, item.Title, item.Author)
		}
	} else {
		for _, item := range response.Files {
			idParts := strings.SplitAfterN(item.ID, "/", 5)
			fmt.Printf(messages.FileItem, idParts[4], item.Language.GetTitle(), item.Title, item.Author)
		}
	}

	return nil
}
