package interactive

import (
	"subteez/messages"

	"github.com/rivo/tview"
)

func (c *Context) showStatusDialog(message string) {
	modal := tview.NewModal().SetText(message)
	modal.SetTitle(messages.AppTitle)
	c.app.SetRoot(modal, false).SetFocus(modal)
}

func (c *Context) showExitDialog() {
	exitDialog := tview.NewModal().
		SetText(messages.ExitDialogText).
		AddButtons([]string{messages.ButtonYes, messages.ButtonNo}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				c.app.Stop()
			} else {
				c.popView()
			}
		})
	exitDialog.SetTitle(messages.AppTitle)

	c.pushView(exitDialog)
}
