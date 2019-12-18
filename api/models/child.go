package models

import (
	"fmt"
	"time"

	u "../utils"
)

type Child struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	// CreatedAt   string `json:"created_at"`
	// UpdatedAt   string `json:"updated_at"`
	// Deleted     int    `json:"deleted"`
}

type Children struct {
	children []Child
}

func (child *Child) Validate() (map[string]interface{}, bool) {
	if child.Name == "" {
		return u.Message(false, "Please enter child name"), false
	}

	if child.UserID <= 0 {
		return u.Message(false, "Please login"), false
	}

	return u.Message(true, "success"), true
}

func (child *Child) Create() map[string]interface{} {
	if res, ok := child.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `children` (`user_id`,`name`,`date_of_birth`,`gender`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")
	t := time.Now()
	result, _ := statement.Exec(child.UserID, child.Name, child.DateOfBirth, child.Gender, t.Format("2006-01-02T15:04:05Z07:00"), t.Format("2006-01-02T15:04:05Z07:00"))
	lastid, _ := result.LastInsertId()

	child.ID = int(lastid)

	res := u.Message(true, "success")
	res["child"] = child
	return res
}

func GetChildrenByUser(user int) []Child {
	var children []Child

	database := GetDB()
	rows, err := database.Query("SELECT id, user_id, name, date_of_birth, gender FROM children WHERE user_id = ? AND deleted = 0", user)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var child Child
		_ = rows.Scan(&child.ID, &child.UserID, &child.Name, &child.DateOfBirth, &child.Gender)
		children = append(children, child)
	}

	rows.Close()

	return children

}

func GetChildren() []Child {
	var children []Child
	database := GetDB()
	rows, _ := database.Query("SELECT id, user_id, name, date_of_birth, gender FROM `children` where deleted = 0")

	for rows.Next() {
		var child Child
		_ = rows.Scan(&child.ID, &child.UserID, &child.Name, &child.DateOfBirth, &child.Gender)
		children = append(children, child)
	}

	rows.Close() //good habit to close

	return children
}

func DeleteChild(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("update `children` SET `deleted` = 1, `updated_at` = ? where id=?")
	t := time.Now()
	stmt.Exec(t.Format("2006-01-02T15:04:05Z07:00"), id)

	res2 := u.Message(true, "success")

	return res2
}

func (child *Child) UpdateChild(id string) map[string]interface{} {
	if res, ok := child.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("UPDATE `children` SET `name`= ?, `date_of_birth`=?, `gender`=?, `updated_at` = ? WHERE id =?")
	t := time.Now()
	result, _ := statement.Exec(child.Name, child.DateOfBirth, child.Gender, t.Format("2006-01-02T15:04:05Z07:00"), id)

	res := u.Message(true, "success")
	lastid, _ := result.LastInsertId()
	child.ID = int(lastid)
	res["child"] = child
	return res

}
