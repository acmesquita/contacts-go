package ui

import (
	"fmt"
	"os"

	"github.com/acmesquita/contacts/models"
	"github.com/rivo/tview"
)

func showList() {
	if err := app.SetRoot(people, true).SetFocus(people).Run(); err != nil {
		panic(err)
	}
}

func addPersonToList(person *models.Person) {
	fullName := fmt.Sprintf("%v, %v", person.LastName, person.FirstName)

	people.AddItem(fullName, person.Contact, rune(person.ID[0]), nil)
}

func quitItem() *tview.List {
	return people.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
		os.Exit(0)
	})
}

func addPersonItem(people_count int) *tview.List {
	return people.AddItem("Add", "Press to adding new contact", 'a', func() {
		people_count += 1
		form := createForm(&models.Person{}, &people_count)
		showForm(form)
	})
}
