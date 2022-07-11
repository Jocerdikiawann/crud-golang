package requests

type UserRequest struct {
	Email       string `json:"email" bson:"email,omitempty"`
	Password    []byte `json:"password" bson:"password,omitempty"`
	FirstName   string `json:"firstName" bson:"firstName,omitempty"`
	LastName    string `json:"lastName" bson:"lastName,omitempty"`
	Address     string `json:"address" bson:"address,omitempty"`
	AccessToken string `json:"accessToken" bson:"accessToken,omitempty"`
}

type UserBody struct {
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Address   string `json:"address" bson:"address"`
}
