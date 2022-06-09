package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/goombaio/namegenerator"
	manager "go_fyne/src"
	"log"
	"strconv"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	logger := manager.NewLogger("Student System")
	db := manager.InitDatabase(logger)
	studentManager := manager.StudentManager(db, logger)
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	s := studentManager.CreateUser(&manager.Student{
		FName:        nameGenerator.Generate(),
		LName:        nameGenerator.Generate(),
		PhoneNumber:  "05315313131",
		FPhoneNumber: "05696969696",
		Adress:       nameGenerator.Generate(),
	})

	r := studentManager.GetStudentWithUUID(s.UUID.String())
	r.DevamsizlikEkle("ozursuz")

	var data []*manager.Student
	db.Find(&data)

	var activeUser *manager.Student
	activeUser = data[0]

	var list *widget.List

	list = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("Will be replaced")
			btn1 := widget.NewButton("Update", nil)
			btn2 := widget.NewButton("Remove", nil)
			return container.New(
				layout.NewGridLayout(6),
				label,
				layout.NewSpacer(),
				layout.NewSpacer(),
				layout.NewSpacer(),
				btn1,
				btn2,
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(strconv.Itoa(int(data[id].ID)) + " " + data[id].GetStudentName())
			item.(*fyne.Container).Objects[4].(*widget.Button).OnTapped = func() {

				FNameEntry := widget.NewEntry()
				LNameEntry := widget.NewEntry()
				FamilyPhoneEntry := widget.NewEntry()
				StudentPhoneEntry := widget.NewEntry()
				AdressEntry := widget.NewMultiLineEntry()
				Dozs := widget.NewEntry()
				Dozl := widget.NewEntry()

				FNameEntry.Text = data[id].FName
				LNameEntry.Text = data[id].LName
				StudentPhoneEntry.Text = data[id].PhoneNumber
				FamilyPhoneEntry.Text = data[id].FPhoneNumber
				AdressEntry.Text = data[id].Adress
				Dozs.Text = strconv.Itoa(int(data[id].Dozs))
				Dozl.Text = strconv.Itoa(int(data[id].Dozl))

				updateUserForm := &widget.Form{
					Items: []*widget.FormItem{ // we can specify items in the constructor
						{Text: "First Name", Widget: FNameEntry},
						{Text: "Last Name", Widget: LNameEntry},
						{Text: "Family Phone", Widget: FamilyPhoneEntry},
						{Text: "Student Phone", Widget: StudentPhoneEntry},
						{Text: "Adress", Widget: AdressEntry},
						{Text: "Özürsüz Devamsızlık", Widget: Dozs},
						{Text: "Özürlü Devamsızlık", Widget: Dozl},
					},
					OnSubmit: func() {
						dozsv, err := strconv.Atoi(Dozs.Text)
						dozlv, err := strconv.Atoi(Dozl.Text)
						if err != nil {
							panic(err)
						}
						studentManager.UpdateUser(&manager.Student{
							UUID:         data[id].UUID,
							FName:        FNameEntry.Text,
							LName:        LNameEntry.Text,
							PhoneNumber:  StudentPhoneEntry.Text,
							FPhoneNumber: FamilyPhoneEntry.Text,
							Adress:       AdressEntry.Text,
							Dozs:         dozsv,
							Dozl:         dozlv,
						})
						db.Find(&data)
						list.Refresh()
					},
				}

				dialog.ShowCustom("Update", "Close", container.NewVBox(widget.NewLabel(strconv.Itoa(int(data[id].ID))+"  -  "+data[id].UUID.String()), updateUserForm), w)
				fmt.Println("I am button " + data[id].GetStudentName())
			}
			item.(*fyne.Container).Objects[5].(*widget.Button).OnTapped = func() {
				d := dialog.NewConfirm("Confirm", "Are you sure", func(t bool) {
					if t == true {
						studentManager.RemoveUserWithID(data[id].ID)
						db.Find(&data)
						list.Refresh()
					}
				}, w)
				d.Show()
				fmt.Println("I am button " + data[id].GetStudentName())
			}
		},
	)

	FNameEntry := widget.NewEntry()
	LNameEntry := widget.NewEntry()
	FamilyPhoneEntry := widget.NewEntry()
	StudentPhoneEntry := widget.NewEntry()
	AdressEntry := widget.NewMultiLineEntry()

	setEntry(FNameEntry, "Omer")
	setEntry(LNameEntry, "Faruk")
	setEntry(FamilyPhoneEntry, "5316987123")
	setEntry(StudentPhoneEntry, "5321694789")
	setEntry(AdressEntry, "Bursa/osmangazi ısık sk. elixxrade apart.")

	addUserForm := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "First Name", Widget: FNameEntry},
			{Text: "Last Name", Widget: LNameEntry},
			{Text: "Family Phone", Widget: FamilyPhoneEntry},
			{Text: "Student Phone", Widget: StudentPhoneEntry},
			{Text: "Adress", Widget: AdressEntry},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", FNameEntry.Text)
			log.Println("multiline:", AdressEntry.Text)

			studentManager.CreateUser(&manager.Student{
				FName:        FNameEntry.Text,
				LName:        LNameEntry.Text,
				PhoneNumber:  StudentPhoneEntry.Text,
				FPhoneNumber: FamilyPhoneEntry.Text,
				Adress:       AdressEntry.Text,
			})
			db.Find(&data)
			list.Refresh()
		},
	}

	box3 := container.NewVBox(widget.NewLabel("Add User Form"), addUserForm)

	name := binding.NewString()
	familyPhone := binding.NewString()
	phone := binding.NewString()
	adress := binding.NewString()
	dozs := binding.NewString()
	dozl := binding.NewString()
	toplam := binding.NewString()

	setDatas(activeUser, name, familyPhone, phone, adress, dozs, dozl, toplam)

	box2 := container.NewVBox(
		widget.NewLabelWithData(name),
		widget.NewLabelWithData(familyPhone),
		widget.NewLabelWithData(phone),
		widget.NewLabelWithData(adress),
		widget.NewLabelWithData(dozs),
		widget.NewLabelWithData(dozl),
		widget.NewLabelWithData(toplam),
	)
	box1 := container.NewVSplit(box2, box3)

	list.OnSelected = func(id widget.ListItemID) {
		activeUser = data[id]
		setDatas(activeUser, name, familyPhone, phone, adress, dozs, dozl, toplam)
		box2.Refresh()
	}

	w.SetContent(container.NewHSplit(list, box1))
	w.Resize(fyne.NewSize(1600, 800))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func setEntry(entry *widget.Entry, placeholder string) {
	entry.PlaceHolder = placeholder
	entry.Validator = func(str string) error {
		if len(str) < 2 {
			return errors.New("Too short")
		}
		return nil
	}
	entry.Refresh()
}

func setDatas(usr *manager.Student, namex binding.String, familyPhonex binding.String, phonex binding.String, adressx binding.String, dozsx binding.String, dozlx binding.String, toplamx binding.String) {
	namex.Set("Name : " + usr.GetStudentName())
	familyPhonex.Set("Family Phone : " + usr.FPhoneNumber)
	phonex.Set("Phone : " + usr.PhoneNumber)
	adressx.Set("Adress : " + usr.Adress)
	dozsx.Set("Özürlü Devamsızlık : " + strconv.Itoa(int(usr.Dozs)))
	dozlx.Set("Özürsüz Devamsızlık : " + strconv.Itoa(int(usr.Dozs)))
	toplamx.Set("Toplam Devamsızlık : " + usr.GetDevamsizlikSTR())
}
