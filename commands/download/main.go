package download

import (
	"fmt"
	"log"
	"os"
	"strings"
	"subteez/config"
	"subteez/errors"
	"subteez/interactive"
	"subteez/messages"
	"subteez/subteez"
)

func Main(args []string, cfg config.Config) error {
	client := subteez.NewClient(cfg.GetServer())

	if len(args) < 2 {
		return errors.ErrNotEnoughArguments
	}

	id := args[1]

	if !strings.HasPrefix(id, "/subtitles/") {
		if len(args) < 4 {
			return errors.ErrNotEnoughArguments
		}
		language, err := subteez.ParseLanguage(args[2])
		if err != nil {
			return err
		}
		languageDownloadPathPart, err := language.GetDownloadPathPart()
		if err != nil {
			return err
		}
		id = fmt.Sprintf("/subtitles/%s/%s/%s", args[1], languageDownloadPathPart, args[3])
	}

	if cfg.IsInteractive() {
		context := interactive.Context{}
		context.Initialize(cfg)
		go context.NavigateToDownload(id)
		return context.Run()
	}

	if !cfg.IsScriptMode() {
		log.Printf(messages.GettingPage, id)
	}

	request := subteez.SubtitleDownloadRequest{
		ID: id,
	}
	name, data, err := client.Download(request)
	if err != nil {
		return err
	}

	if cfg.IsScriptMode() {
		if _, err := os.Stdout.Write(data); err != nil {
			return err
		}
	} else {
		file, err := os.Create(name)
		if err != nil {
			return err
		}
		if _, err = file.Write(data); err != nil {
			return err
		}
		log.Printf(messages.FileWritten, name)
	}
	return nil
}
