// main of "files" command

package files_command

import (
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
	id := args[1]

	if id == "" {
		return errors.ErrEmptyID
	}

	// generate full address if given address is not a full address
	if !strings.HasPrefix(id, "/subtitles/") {
		id = "/subtitles/" + id
	}

	// run in interactive mode if it's enabled
	if cfg.IsInteractive() {
		context := tui.Context{}
		context.Initialize(cfg)
		go context.NavigateToDetails(id)
		return context.Run()
	}

	if !cfg.IsScriptMode() {
		log.Printf(messages.GettingPage, id)
	}

	// send request
	client := subteez.NewClient(cfg.GetServer())
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

	// print result

	if cfg.IsScriptMode() {
		for _, item := range response.Files {
			fmt.Printf("%s,%s,%s,%s\n", item.ID, item.Language, item.Title, item.Author)
		}
		return nil
	}

	for _, item := range response.Files {
		idParts := strings.SplitAfterN(item.ID, "/", 5)
		fmt.Printf(messages.FileItem, idParts[4], item.Language.GetTitle(), item.Title, item.Author)
	}
	return nil
}
