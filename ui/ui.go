package ui

import (
	"github.com/rivo/tview"
)

var (
	people_count int
	app          *tview.Application
	people       *tview.List
)

func Init() {
	people_count = 0
	app = tview.NewApplication()
	people = tview.NewList()
	quitItem()
	addPersonItem(people_count)

	showList()
}
