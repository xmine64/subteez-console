package interactive

import (
	"subteez/config"
	"subteez/subteez"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Context struct {
	config    config.Config
	query     string
	error     error
	app       *tview.Application
	client    subteez.ISubteezAPI
	movieId   string
	fileId    string
	viewStack []tview.Primitive
}

func (c *Context) pushView(view tview.Primitive) {
	c.viewStack = append(c.viewStack, view)
	c.activateViewStack()
}

func (c *Context) popView() bool {
	length := len(c.viewStack) - 1
	if length <= 0 {
		return false
	}
	c.viewStack = c.viewStack[:length]
	c.activateViewStack()
	return true
}

func (c *Context) activateViewStack() {
	current := c.viewStack[len(c.viewStack)-1]
	c.app.SetRoot(current, true).SetFocus(current)
}

func (c *Context) Initialize(cfg config.Config) {
	c.config = cfg
	c.client = subteez.NewClient(cfg.GetServer())
	c.app = tview.NewApplication()
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			if !c.popView() {
				c.showExitDialog()
			}
			return nil
		}
		return event
	})
	c.showSearchQueryPrompt()
}

func (c *Context) Run() error {
	if err := c.app.Run(); err != nil {
		return err
	}
	return c.error
}

// must not be called from main loop and event handlers
func (c *Context) NavigateToSearchResult(query string) {
	if query == "" {
		return
	}
	c.query = query
	c.viewStack = c.viewStack[:0]
	c.showSearchResultList()
}

// must not be called from main loop and event handlers
func (c *Context) NavigateToDetails(id string) {
	if id == "" {
		return
	}
	c.movieId = id
	c.viewStack = c.viewStack[:0]
	c.showItemDetails()
}

// must not be called from main loop and event handlers
func (c *Context) NavigateToDownload(id string) {
	if id == "" {
		return
	}
	c.fileId = id
	c.viewStack = c.viewStack[:0]
	c.showDownloadDialog()
}

func (c *Context) NavigateToConfig() {
	c.viewStack = c.viewStack[:0]
	c.showConfigForm()
}
