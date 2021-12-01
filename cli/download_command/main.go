// main of "download" command

package download_command

import (
	"fmt"
	"log"
	"os"
	"strings"
	"subteez/config"
	"subteez/errors"
	"subteez/messages"
	"subteez/subteez"
	"subteez/tui"
)

func Main(args []string, cfg config.Config) error {
	if len(args) < 2 {
		return errors.ErrNotEnoughArguments
	}

	id := args[1]

	// generate full address, if given address is not a full address
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

	// run interactive mode if it's enabled
	client := subteez.NewClient(cfg.GetServer())
	if cfg.IsInteractive() {
		context := tui.Context{}
		context.Initialize(cfg)
		go context.NavigateToDownload(id)
		return context.Run()
	}

	if !cfg.IsScriptMode() {
		log.Printf(messages.GettingPage, id)
	}

	// send request
	request := subteez.SubtitleDownloadRequest{
		ID: id,
	}
	name, data, err := client.Download(request)
	if err != nil {
		return err
	}

	// write downloaded file directly to stdout if script mode is enabled
	if cfg.IsScriptMode() {
		_, err := os.Stdout.Write(data)
		return err
	}

	// write downloaded file to disk
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(data); err != nil {
		return err
	}
	log.Printf(messages.FileWritten, name)

	return nil
}
