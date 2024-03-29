package auth

import (
	"bookstorage_web/server/dao"
	"bookstorage_web/server/model"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestSignUp(t *testing.T) {
	dao.DBInit()
	var user model.User
	demoUserJson := []byte(`{"user_name":"Alice","email_address":"alice@gmail.com","password":"passdesuyo"}`)

	if err := json.Unmarshal(demoUserJson, &user); err != nil {
		t.Error(err.Error())
	}

	if user.Name == "" {
		t.Error("error user name is blank")
	} else {
		fmt.Println("name: " + user.Name)
	}

	if user.EmailAddress == "" {
		t.Error("error user email_address is blank")
	} else {
		fmt.Println("emailAddress: " + user.EmailAddress)
	}

	if user.Password == "" {
		t.Error("error user password is blank")
	} else {
		fmt.Println("password: " + user.Password)
	}

	hashedPassWord := hashStringPassWord(user.Password)
	user.Password = hashedPassWord

	dao.Create(user)

	fmt.Println("hashed password: " + user.Password)

	jwtAccessToken := GetJwtAccessToken(user)

	fmt.Println("jwtAccessToken: " + jwtAccessToken)

}

func TestLogin(t *testing.T) {

}

func TestValidateUser(t *testing.T) {
	targetTrue := model.User{
		Name:         "s",
		EmailAddress: "s@gmail.com",
		Password:     "oh"}

	registered := model.User{
		Name:         "s",
		EmailAddress: "s@gmail.com",
		Password:     hashStringPassWord("oh"),
	}

	result1 := bcrypt.CompareHashAndPassword([]byte(registered.Password), []byte(targetTrue.Password))
	if result1 == nil {
		fmt.Println("Validation1 Success")
	} else {
		t.Error(result1.Error())
	}

	targetFlase := model.User{
		Name:         "s",
		EmailAddress: "s@gmail.com",
		Password:     "o",
	}

	result2 := bcrypt.CompareHashAndPassword([]byte(registered.Password), []byte(targetFlase.Password))
	if result2 == nil {
		t.Error("Validation2 Expected false but got true")
	} else {
		fmt.Println("Validation2 Success")
	}

}

func TestHashStringPassWord(t *testing.T) {
	demoPassword := "oh"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(demoPassword), 10)
	if err != nil {
		t.Error(err.Error())
	}
	hashedStringPass := string(hashedPassword)
	fmt.Println("パスワード: ", demoPassword)
	fmt.Println("ハッシュ化されたパスワード", hashedPassword)
	fmt.Println("コンバート後のパスワード: ", hashedStringPass)
}
