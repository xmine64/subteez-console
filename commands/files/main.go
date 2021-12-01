package files

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

	id := strings.Join(flag.Args()[1:], " ")

	if id == "" {
		return errors.ErrEmptyID
	}

	if cfg.IsInteractive() {
		context := interactive.Context{}
		context.Initialize(cfg)
		go context.NavigateToDetails(id)
		return context.Run()
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

	for _, item := range response.Files {
		fmt.Printf(messages.FileRow, item.ID, item.Language.GetTitle(), item.Title, item.Author)
	}

	return nil
}
