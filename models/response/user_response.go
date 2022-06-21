package response

type UserResponse struct {
	Id        string `json:"_id" bson:"_id"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Address   string `json:"address" bson:"address"`
}
