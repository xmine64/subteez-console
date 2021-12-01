// interactive TUI home view

package tui

import (
	"subteez/messages"

	"github.com/rivo/tview"
)

func (c *Context) showHome() {
	input := tview.NewInputField().
		SetLabel(messages.QueryPrompt).
		SetText(c.query).
		SetFieldWidth(40).
		SetChangedFunc(func(text string) {
			c.query = text
		}).
		SetPlaceholder(messages.QueryPlaceholder)

	form := tview.NewForm().AddFormItem(input).AddButton(messages.ButtonOK, func() {
		if len(c.query) > 3 {
			go c.showSearchResultList()
		}
	}).AddButton(messages.ButtonCancel, func() {
		c.showExitDialog()
	}).AddButton(messages.ButtonConfig, func() {
		c.showConfigForm()
	})
	form.SetBorder(true).SetTitle(messages.AppTitle)

	colFlex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(form, 60, 1, true).
		AddItem(nil, 0, 1, false)

	rowFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(colFlex, 7, 1, true).
		AddItem(nil, 0, 1, false)

	c.pushView(rowFlex)
}
