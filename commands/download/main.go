package download

import (
	"fmt"
	"log"
	"os"
	"strings"
	"subteez/config"
	"subteez/errors"
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

	request := subteez.SubtitleDownloadRequest{
		ID: id,
	}
	name, data, err := client.Download(request)
	if err != nil {
		return err
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	count, err := file.Write(data)
	if err != nil {
		return err
	}
	log.Printf("%d bytes written to file %s", count, name)
	return nil
}
