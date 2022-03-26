package details

import "fmt"

type UserDetail struct {
	Id      int
	Fname   string
	City    string
	Phone   string
	Height  int
	Married bool
}

var users []UserDetail

func CreateUser() *UserDetail {
	return &UserDetail{}
}
func GetUser() *[]UserDetail {
	return &users
}

func ListAllUser() []UserDetail {
	var listUser []UserDetail
	for _, user := range users {
		listUser = append(listUser, user)
	}
	return listUser
}

//set user details
func (setUser *UserDetail) SetUserDetail(getUser UserDetail) {
	//setting user details
	setUser.Id = getUser.Id
	setUser.Fname = getUser.Fname
	setUser.City = getUser.City
	setUser.Phone = getUser.Phone
	setUser.Height = getUser.Height
	setUser.Married = getUser.Married
}

func AddUserToList(user *UserDetail) {

	//appneding user to users list
	users = append(users, *user)
}

//get user details from user list and return user
func GetUserDetail(id int) (UserDetail, error) {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}
	return UserDetail{}, fmt.Errorf("user not found!")
}

func DeleteUser(id int) error {
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found!")
}

func DeleteAllUsers() error {
	users = nil
	if users == nil {
		return nil
	}
	return fmt.Errorf("User is not deleted")
}

func UpdateUser(id int, fname, city, phone string, height int, marrried bool) error {
	for i, user := range users {
		if user.Id == id {
			fmt.Println("hit")
			users[i].Fname = fname
			users[i].City = city
			users[i].Phone = phone
			users[i].Height = height
			users[i].Married = marrried
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

func GetLastId() int {
	if len(users) == 0 {
		return 0
	}
	id := users[len(users)-1].Id

	return id
}
