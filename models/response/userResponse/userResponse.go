package userresponse

type UserResponse struct {
	Id          string `json:"_id" bson:"_id"`
	Email       string `json:"email" bson:"email"`
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	Address     string `json:"address" bson:"address"`
	AccessToken string `json:"access_token" bson:"access_token,omitempty"`
}
