// Interactive TUI main

package tui

import (
	"subteez/config"
	"subteez/subteez"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Interactive TUI context
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

// push a view to view stack and show it
func (c *Context) pushView(view tview.Primitive) {
	c.viewStack = append(c.viewStack, view)
	c.activateViewStack()
}

// remove top view from stack and show its previous view, if possible
func (c *Context) popView() bool {
	length := len(c.viewStack) - 1
	if length <= 0 {
		return false
	}
	c.viewStack = c.viewStack[:length]
	c.activateViewStack()
	return true
}

// show top view of stack
func (c *Context) activateViewStack() {
	current := c.viewStack[len(c.viewStack)-1]
	c.app.SetRoot(current, true).SetFocus(current)
}

// initialize interactive TUI's context
func (c *Context) Initialize(cfg config.Config) {
	c.config = cfg
	c.client = subteez.NewClient(cfg.GetServer())
	c.app = tview.NewApplication().
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			// go back if possible, or show a confirm to exit dialog
			if event.Key() == tcell.KeyESC {
				if !c.popView() {
					c.showExitDialog()
				}
				return nil
			}

			return event
		})

	c.showHome()
}

// run interactive TUI
func (c *Context) Run() error {
	if err := c.app.Run(); err != nil {
		return err
	}
	return c.error
}

// set search query and replace current view with search result view, if query is not empty
// must not be called from main loop and event handlers
func (c *Context) NavigateToSearchResult(query string) {
	if query == "" {
		return
	}
	c.query = query
	c.viewStack = c.viewStack[1:]
	c.showSearchResultList()
}

// set movie ID and replace current view with subtitle details view, if ID is not empty
// must not be called from main loop and event handlers
func (c *Context) NavigateToDetails(id string) {
	if id == "" {
		return
	}
	c.movieId = id
	c.viewStack = c.viewStack[:0]
	c.showItemDetails()
}

// set file ID and replace current view with file download dialog, if ID is not empty
// must not be called from main loop and event handlers
func (c *Context) NavigateToDownload(id string) {
	if id == "" {
		return
	}
	c.fileId = id
	c.viewStack = c.viewStack[:0]
	c.showDownloadDialog()
}

// replace current view with config form
func (c *Context) NavigateToConfig() {
	c.viewStack = c.viewStack[:0]
	c.showConfigForm()
}
