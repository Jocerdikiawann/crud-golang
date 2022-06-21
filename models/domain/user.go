package domain

type User struct {
	Id        string `json:"_id" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Address   string `json:"address" bson:"address"`
}
