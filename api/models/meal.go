package models

type Meal struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Deleted     int    `json:"deleted"`
	Ingredients []Ingredient
	Feedback    []Feedback
}

type Ingredient struct {
	ID        int    `json:"id"`
	Meal      int    `json:"meal_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   int    `json:"deleted"`
}

type Feedback struct {
	ID        int     `json:"id"`
	Meal      int     `json:"meal_id"`
	ChildID   int     `json:"child_id"`
	ChildName string  `json:"child_name"`
	Rating    float32 `json:"rating"`
	DidEat    bool    `json:"did_eat"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Deleted   int     `json:"deleted"`
}
