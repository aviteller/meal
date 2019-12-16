package models

import (
	"fmt"
	"time"

	u "../utils"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   int    `json:"deleted"`
	Children  []Child
}

func (user *User) Validate() (map[string]interface{}, bool) {
	// if !strings.Contains(account.Email, "@") {
	// 	return u.Message(false, "Please enter valid email address"), false
	// }

	if len(user.Password) < 6 {
		return u.Message(false, "Password needs to be greater then 6 chars"), false
	}

	temp := &User{}
	err := GetDB().QueryRow("SELECT name FROM `users` WHERE name = ? ", user.Name).Scan(&temp.Name)
	if err != nil {
		fmt.Println(err)
	}

	if temp.Name != "" {
		return u.Message(false, "Username has already been raken"), false
	}
	return u.Message(false, "Requirement passed"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	fmt.Println(user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `users` (`name`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?)")
	t := time.Now()
	result, _ := statement.Exec(user.Name, user.Password, t.Format("2006-01-02T15:04:05Z07:00"), t.Format("2006-01-02T15:04:05Z07:00"))

	lastid, _ := result.LastInsertId()

	tk := &Token{UserId: uint(lastid)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	user.Token = tokenString

	user.Password = ""
	res := u.Message(true, "User has been created")
	res["user"] = user
	return res
}

func Login(name, password string) map[string]interface{} {
	user := &User{}

	database := GetDB()
	err := database.QueryRow("SELECT id, name, email, password FROM users WHERE name = ?", name).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {

		fmt.Println(err)
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		fmt.Println(err)
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	user.Password = ""

	tk := &Token{UserId: uint(user.ID)}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, _ := token.SignedString([]byte("secret"))

	user.Token = tokenString

	res := u.Message(true, "Logged In")

	res["user"] = user
	return res
}

func GetUser(u int) User {
	var user User

	database := GetDB()
	err := database.QueryRow("SELECT id, name FROM users WHERE id = ? AND deleted = 0 LIMIT 0, 1", u).Scan(&user.ID, &user.Name)

	if err != nil {
		fmt.Println(err)
	}

	user.Children = GetChildrenByUser(u)

	return user

}
