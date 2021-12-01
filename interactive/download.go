package interactive

import (
	"fmt"
	"os"
	"subteez/messages"
	"subteez/subteez"

	"github.com/rivo/tview"
)

// must not be called from main loop and event handlers
func (c *Context) showDownloadDialog() {
	c.app.QueueUpdateDraw(func() {
		c.showStatusDialog(messages.DownloadingStatus)
	})

	request := subteez.SubtitleDownloadRequest{
		ID: c.fileId,
	}
	name, data, err := c.client.Download(request)
	if err != nil {
		c.error = err
		c.app.Stop()
		return
	}

	c.app.QueueUpdateDraw(func() {
		c.showStatusDialog(messages.WritingStatus)
	})

	file, err := os.Create(name)
	if err != nil {
		c.error = err
		c.app.Stop()
		return
	}
	defer file.Close()

	count, err := file.Write(data)
	if err != nil {
		c.error = err
		c.app.Stop()
		return
	}

	modal := tview.NewModal().
		SetText(fmt.Sprintf(messages.FileWrittenDialogText, count, name)).
		AddButtons([]string{messages.ButtonOK}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if len(c.viewStack)-1 <= 0 {
				c.app.Stop()
				return
			}
			c.popView()
		})

	c.app.QueueUpdateDraw(func() {
		c.pushView(modal)
	})
}
