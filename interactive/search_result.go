package interactive

import (
	"errors"
	"fmt"
	"subteez/messages"
	"subteez/subteez"

	"github.com/rivo/tview"
)

// must not be called from main loop and event handlers
func (c *Context) showSearchResultList() {
	if c.query == "" {
		c.error = errors.New(messages.EmptyQuery)
		c.app.Stop()
		return
	}

	c.app.QueueUpdateDraw(func() {
		c.showStatusDialog(messages.FetchingData)
	})

	request := subteez.SearchRequest{
		Query:           c.query,
		LanguageFilters: c.config.GetLanguageFilters(),
	}
	response, err := c.client.Search(request)
	if err != nil {
		c.error = err
		c.app.Stop()
		return
	}

	if len(response.Result) < 1 {
		c.error = errors.New(messages.NoSearchResult)
		c.app.Stop()
		return
	}

	list := tview.NewList().SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		c.movieId = response.Result[index].ID
	})

	for _, item := range response.Result {
		list.AddItem(item.Name, fmt.Sprintf(messages.FilesCount, item.Count), 0, func() {
			go c.showItemDetails()
		})
	}

	list.SetBorder(true).SetTitle(fmt.Sprintf(messages.SearchResultTitle, c.query))

	c.app.QueueUpdateDraw(func() {
		c.pushView(list)
	})
}
