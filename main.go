package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Schedule Meeting")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter Name")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter Email")

	dayEntry := widget.NewEntry()
	dayEntry.SetPlaceHolder("Enter day")

	monthEntry := widget.NewEntry()
	monthEntry.SetPlaceHolder("Enter month")

	hourEntry := widget.NewEntry()
	hourEntry.SetPlaceHolder("Enter hour")
	minuteEntry := widget.NewEntry()
	minuteEntry.SetPlaceHolder("Enter minute")

	button := widget.NewButton("Schedule Meeting", func() {
		name := nameEntry.Text
		email := emailEntry.Text
		date := datePicker.Date
		fmt.Printf("Name: %s, Email: %s, Date: %s\n", name, email, date.Format("2006-01-02"))
		// Here you can call your function to schedule a meeting
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Email", Widget: emailEntry},
			{Text: "Date", Widget: datePicker},
		},
		OnSubmit: func() {
			// This function will be called when the user hits the return key
			button.OnTapped()
		},
		SubmitText: "Schedule Meeting",
	}

	window.SetContent(container.NewVBox(form, button))
	window.Resize(fyne.NewSize(300, 200))
	window.ShowAndRun()
}
