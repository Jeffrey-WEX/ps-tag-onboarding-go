package model

type User struct {
	ID               string   `json:"id" bson:"_id,omitempty"`
	FirstName        string   `json:"first_name" bson:"firstName"`
	LastName         string   `json:"last_name" bson:"lastName"`
	Email            string   `json:"email" bson:"email"`
	Age              int      `json:"age" bson:"age"`
	ValidationErrors []string `json:"-" bson:"-"`
}
