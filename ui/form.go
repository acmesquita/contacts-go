package ui

import (
	"strconv"

	"github.com/acmesquita/contacts/models"
	"github.com/acmesquita/contacts/services"
	"github.com/rivo/tview"
)

func showForm(form *tview.Form) {
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func createForm(person *models.Person, people_count *int) *tview.Form {
	person.ID = []byte(strconv.Itoa(*people_count))
	return initForm(person)
}

func initForm(person *models.Person) *tview.Form {
	form := tview.NewForm()

	firstNameField(form, person)
	lastNameField(form, person)
	contactField(form, person)
	saveButton(form, person)
	quitButton(form)

	return configForm(form)
}

func firstNameField(form *tview.Form, person *models.Person) *tview.Form {
	return form.AddInputField("First name", "", 20, nil, func(firstName string) {
		person.FirstName = firstName
	})
}

func lastNameField(form *tview.Form, person *models.Person) *tview.Form {
	return form.AddInputField("Last name", "", 20, nil, func(lastName string) {
		person.LastName = lastName
	})
}

func contactField(form *tview.Form, person *models.Person) *tview.Form {
	return form.AddInputField("Contact", "", 20, nil, func(contact string) {
		person.Contact = contact
	})
}

func saveButton(form *tview.Form, person *models.Person) *tview.Form {
	return form.AddButton("Save", func() {
		services.SavePerson(person)
		addPersonToList(person)
		showList()
	})
}

func quitButton(form *tview.Form) *tview.Form {
	return form.AddButton("Quit", func() {
		showList()
	})
}

func configForm(form *tview.Form) *tview.Form {
	form.SetBorder(true).SetTitle("Create new contact").SetTitleAlign(tview.AlignLeft)

	return form
}
