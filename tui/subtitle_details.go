// interactive TUI for showing subtitle files

package tui

import (
	"fmt"
	"subteez/errors"
	"subteez/messages"
	"subteez/subteez"

	"github.com/rivo/tview"
)

// must not be called from main loop and event handlers
func (c *Context) showItemDetails() {
	if c.movieId == "" {
		c.error = errors.ErrEmptyID
		c.app.Stop()
		return
	}

	c.app.QueueUpdateDraw(func() {
		c.showStatusDialog(messages.FetchingData)
	})

	request := subteez.SubtitleDetailsRequest{
		ID:              c.movieId,
		LanguageFilters: c.config.GetLanguageFilters(),
	}
	response, err := c.client.GetDetails(request)
	if err != nil {
		c.error = err
		c.app.Stop()
		return
	}

	if len(response.Files) < 1 {
		c.error = errors.ErrNoFileFound
		c.app.Stop()
		return
	}

	header := tview.NewTextView().SetText(fmt.Sprintf(
		messages.MovieInfo,
		response.Name,
		len(response.Files),
		response.Year,
	))
	header.SetBorder(true).SetTitle(messages.MovieInfoTitle)

	details := tview.NewTextView().SetWrap(true).SetWordWrap(true)
	details.SetBorder(true).SetTitle(messages.FileDetailsTitle)

	list := tview.NewList().SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		c.fileId = response.Files[index].ID
		details.SetText(fmt.Sprintf(
			messages.FileDetails,
			response.Files[index].Name,
			response.Files[index].Author,
			response.Files[index].Comment,
		))
	})

	list.SetBorder(true).SetTitle(messages.FilesTitle)

	for _, file := range response.Files {
		item := file
		list.AddItem(item.Title, item.Language.GetTitle(), 0, func() {
			go c.showDownloadDialog()
		})
	}

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 5, 0, false).
		AddItem(list, 0, 1, true).
		AddItem(details, 8, 0, false)

	c.app.QueueUpdateDraw(func() {
		c.pushView(flex)
	})
}
