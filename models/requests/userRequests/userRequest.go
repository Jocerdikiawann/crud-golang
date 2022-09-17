package userrequests

type UserRequest struct {
	Email       string `json:"email" bson:"email"`
	Password    []byte `json:"password" bson:"password"`
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	Address     string `json:"address" bson:"address"`
	AccessToken string `json:"accessToken" bson:"accessToken"`
}

type UserBody struct {
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Address   string `json:"address" bson:"address"`
}

type AuthSignInRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
