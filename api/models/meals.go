package models

import (
	"fmt"
	"time"

	u "../utils"
)

type Ingredient struct {
	ID      int    `json:"id"`
	MealID  int    `json:"meal_id"`
	Name    string `json:"name"`
	Calorie int    `json:"calorie"`
	// CreatedAt string `json:"created_at"`
	// UpdatedAt string `json:"updated_at"`
	// Deleted   int    `json:"deleted"`
}

type Feedback struct {
	ID        int     `json:"id"`
	MealID    int     `json:"meal_id"`
	ChildID   int     `json:"child_id"`
	ChildName string  `json:"child_name"`
	Rating    float32 `json:"rating"`
	DidEat    bool    `json:"did_eat"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Deleted   int     `json:"deleted"`
}

//Ingredient
func (i *Ingredient) Validate() (map[string]interface{}, bool) {
	if i.Name == "" {
		return u.Message(false, "Please enter meal name"), false
	}

	if i.MealID <= 0 {
		return u.Message(false, "Please login"), false
	}

	return u.Message(true, "success"), true
}

func (i *Ingredient) Create() map[string]interface{} {
	if res, ok := i.Validate(); !ok {
		return res
	}
	// fmt.Println(i)
	database := GetDB()
	statement, err := database.Prepare("INSERT INTO `ingredients` (`meal_id`,`name`,`calorie`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")

	if err != nil {
		fmt.Println(err)
	}

	t := time.Now()
	result, err := statement.Exec(i.MealID, i.Name, i.Calorie, t.Format("2006-01-02T15:04:05Z07:00"), t.Format("2006-01-02T15:04:05Z07:00"))
	if err != nil {
		fmt.Println(err)
	}
	lastid, _ := result.LastInsertId()

	i.ID = int(lastid)

	res := u.Message(true, "success")
	res["ingredient"] = i
	return res
}

func GetIngredientsByMeal(meal_id int) []Ingredient {
	var ingredients []Ingredient

	database := GetDB()
	rows, err := database.Query("SELECT id, meal_id, name, calorie FROM ingredients WHERE meal_id = ? AND deleted = 0", meal_id)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var i Ingredient
		_ = rows.Scan(&i.ID, &i.MealID, &i.Name, &i.Calorie)
		ingredients = append(ingredients, i)
	}

	rows.Close()

	return ingredients

}

func DeleteIngredient(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("update `ingredients` SET `deleted` = 1, `updated_at` = ? where id=?")
	t := time.Now()
	stmt.Exec(t.Format("2006-01-02T15:04:05Z07:00"), id)

	res2 := u.Message(true, "success")

	return res2
}

//MEALS

type Meal struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	MealType string `json:"meal_type"` // lunch, breakfast etc
	MealDate string `json:"meal_date"` // lunch, breakfast etc
	// CreatedAt   string `json:"created_at"`
	// UpdatedAt   string `json:"updated_at"`
	// Deleted     int    `json:"deleted"`
	Ingredients []Ingredient
	Feedback    []Feedback
}

func (meal *Meal) Validate() (map[string]interface{}, bool) {
	if meal.Name == "" || meal.MealType == "" {
		return u.Message(false, "Please enter meal name"), false
	}

	if meal.UserID <= 0 {
		return u.Message(false, "Please login"), false
	}

	return u.Message(true, "success"), true
}

func (meal *Meal) Create() map[string]interface{} {
	if res, ok := meal.Validate(); !ok {
		return res
	}

	database := GetDB()
	statement, _ := database.Prepare("INSERT INTO `meals` (`user_id`,`name`,`meal_type`,`meal_date`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")
	t := time.Now()
	result, _ := statement.Exec(meal.UserID, meal.Name, meal.MealType, meal.MealDate, t.Format("2006-01-02T15:04:05Z07:00"), t.Format("2006-01-02T15:04:05Z07:00"))
	lastid, _ := result.LastInsertId()

	meal.ID = int(lastid)

	res := u.Message(true, "success")
	res["meal"] = meal
	return res
}

func GetMealsByUser(user int) []Meal {
	var meals []Meal

	database := GetDB()
	rows, err := database.Query("SELECT id, user_id, name, meal_type, meal_date FROM meals WHERE user_id = ? AND deleted = 0", user)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var m Meal
		_ = rows.Scan(&m.ID, &m.UserID, &m.Name, &m.MealType, &m.MealDate)
		meals = append(meals, m)
	}

	rows.Close()

	return meals

}

func GetMeal(id string) Meal {
	var m Meal

	database := GetDB()
	err := database.QueryRow("SELECT id, user_id, name, meal_type, meal_date FROM meals WHERE id = ? AND deleted = 0 LIMIT 0, 1", id).Scan(&m.ID, &m.UserID, &m.Name, &m.MealType, &m.MealDate)

	if err != nil {
		fmt.Println(err)
	}

	m.Ingredients = GetIngredientsByMeal(m.ID)

	return m
}

func DeleteMeal(id string) map[string]interface{} {

	database := GetDB()
	stmt, _ := database.Prepare("update `meals` SET `deleted` = 1, `updated_at` = ? where id=?")
	t := time.Now()
	stmt.Exec(t.Format("2006-01-02T15:04:05Z07:00"), id)
	stmt, _ = database.Prepare("update `ingredients` SET `deleted` = 1, `updated_at` = ? where meal_id=?")
	stmt.Exec(t.Format("2006-01-02T15:04:05Z07:00"), id)

	res2 := u.Message(true, "success")

	return res2
}
