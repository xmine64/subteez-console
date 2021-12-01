package interactive

import (
	"subteez/errors"
	"subteez/messages"
	"subteez/subteez"

	"github.com/rivo/tview"
)

func (c *Context) showAddLanguageFilterForm() {
	var selected subteez.Language

	dropdown := tview.NewDropDown().SetLabel(messages.LanguageLabel)

Languages:
	for _, language := range subteez.Languages {
		for _, filter := range c.config.GetLanguageFilters() {
			if language == filter {
				continue Languages
			}
		}
		item := language
		dropdown.AddOption(item.GetTitle(), func() {
			selected = item
		})
	}

	form := tview.NewForm().
		AddFormItem(dropdown).
		AddButton(messages.ButtonAdd, func() {
			if string(selected) == "" {
				return
			}
			c.config.AddLanguageFilter(selected)
			c.viewStack = c.viewStack[:len(c.viewStack)-2]
			c.showSetLanguageFilterForm()
		}).
		AddButton(messages.ButtonCancel, func() {
			c.popView()
		})

	form.SetBorder(true).SetTitle(messages.AddFilterTitle)

	c.pushView(form)
}

func (c *Context) showRemoveLanguageFilterDialog(language subteez.Language) {
	confimModal := tview.NewModal().
		SetText(messages.DeleteFilterConfirmText).
		AddButtons([]string{messages.ButtonYes, messages.ButtonNo})
	confimModal.SetBorder(true).SetTitle(messages.ConfirmTitle)
	confimModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == messages.ButtonYes {
			c.config.RemoveLanguageFilter(language)
			c.viewStack = c.viewStack[:len(c.viewStack)-2]
			c.showSetLanguageFilterForm()
		} else {
			c.popView()
		}
	})
	c.pushView(confimModal)
}

func (c *Context) showSetLanguageFilterForm() {
	list := tview.NewList().AddItem(messages.ButtonAdd, "", 'a', func() {
		c.showAddLanguageFilterForm()
	})

	for _, language := range c.config.GetLanguageFilters() {
		item := language
		list.AddItem(item.GetTitle(), "", 0, func() {
			c.showRemoveLanguageFilterDialog(item)
		})
	}

	list.SetBorder(true).SetTitle(messages.LanguageFiltersTitle)
	c.pushView(list)
}

func (c *Context) showConfigForm() {
	form := tview.NewForm().
		AddInputField(messages.ServerLabel, c.config.GetServer(), 50, nil, func(text string) {
			c.config.SetServer(text)
		}).
		AddCheckbox(messages.InteractiveLabel, c.config.IsInteractive(), func(checked bool) {
			c.config.SetInteractive(checked)
		}).
		AddButton(messages.ButtonSave, func() {
			c.error = errors.ErrConfigChanged
			c.app.Stop()
		}).
		AddButton(messages.ButtonCancel, func() {
			c.app.Stop()
		}).
		AddButton(messages.ButtonSetFilter, func() {
			c.showSetLanguageFilterForm()
		})

	form.SetBorder(true).SetTitle(messages.ConfigTitle)

	c.pushView(form)
}
