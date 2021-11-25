package interactive

import (
	"subteez/messages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (c *Context) showSearchQueryPrompt() {
	input := tview.NewInputField().
		SetLabel(messages.QueryPrompt).
		SetPlaceholder(messages.QueryPlaceholder).
		SetText(c.query).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				go c.showSearchResultList()
				return
			}
		}).
		SetChangedFunc(func(text string) {
			c.query = text
		})
	input.SetBorder(true).SetTitle(messages.AppTitle)

	colFlex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(input, 60, 1, true).
		AddItem(nil, 0, 1, false)

	rowFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(colFlex, 3, 1, true).
		AddItem(nil, 0, 1, false)

	c.pushView(rowFlex)
}
