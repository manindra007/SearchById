package userops

import (
	"fmt"
	"strings"

	UD "github.com/details"
)

var User UD.UserDetail

func CreateUser() {
	var id int
	id = UD.GetLastId() + 1
	fname, city, phone, height, married := getUserDetail()
	newuser := UD.CreateUser()
	newuser.SetUserDetail(UD.UserDetail{
		Id:      id,
		Fname:   fname,
		City:    city,
		Phone:   phone,
		Height:  height,
		Married: married,
	},
	)
	UD.AddUserToList(newuser)
}

func DeleteUser() error {
	fmt.Println("enter userId to delete")
	var id int
	fmt.Scanln(&id)
	err := UD.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllUsers() error {

	return UD.DeleteAllUsers()
}

func UpdateUserDetail() error {
	var id int
	fmt.Println("enter userId")
	fmt.Scanln(&id)
	fname, city, phone, height, married := getUserDetail()
	err := UD.UpdateUser(id, fname, city, phone, height, married)
	return err
}

func getUserDetail() (string, string, string, int, bool) {
	var height int
	married := false
	var fname, city, phone string
	fmt.Println("enter fname")
	fmt.Scanln(&fname)
	fmt.Println("enter city")
	fmt.Scanln(&city)
	fmt.Println("enter phone")
	fmt.Scanln(&phone)
	fmt.Println("enter height")
	fmt.Scanln(&height)
	fmt.Println("Married? Y/N")
	var m string
	fmt.Scanln(&m)
	if "yes" == strings.ToLower(m) || "y" == strings.ToLower(m) {
		married = true
	}

	return fname, city, phone, height, married

}
