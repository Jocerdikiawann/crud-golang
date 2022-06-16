package domain

type User struct {
	Id        string `json:"_id" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Address   string `json:"address" bson:"address"`
}
