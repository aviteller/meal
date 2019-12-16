package models

import (
	"time"

	u "../utils"
)

type Child struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   int    `json:"deleted"`
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
	statement, _ := database.Prepare("INSERT INTO `children` (`user_id`,`name`,`age`,`gender`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")
	t := time.Now()
	result, _ := statement.Exec(child.UserID, child.Name, child.Age, child.Gender, t.Format("2006-01-02T15:04:05Z07:00"), t.Format("2006-01-02T15:04:05Z07:00"))
	lastid, _ := result.LastInsertId()

	child.ID = int(lastid)

	res := u.Message(true, "success")
	res["child"] = child
	return res
}

func GetChildrenByUser(user int) []Child {
	var children []Child

	database := GetDB()
	rows, _ := database.Query("SELECT id, user_id, name, age, gender FROM children WHERE user_id = ? AND deleted = 0", user)
	for rows.Next() {
		var child Child
		_ = rows.Scan(&child.ID, &child.UserID, &child.Name, &child.Age, &child.Gender)
		children = append(children, child)
	}

	rows.Close()

	return children

}
