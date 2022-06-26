package domain

type User struct {
	Id        string `json:"_id" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Password  string `json:"password" bson:"password,omitempty"`
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	Address   string `json:"address" bson:"address,omitempty"`
}
