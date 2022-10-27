package user

type User struct {
	Id          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"-" bson:"password"`
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	Address     string `json:"address" bson:"address"`
	AccessToken string `json:"access_token,omitempty" bson:"access_token,omitempty"`
}

type AuthSignIn struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password,omitempty"`
}

type AuthSignUp struct {
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Address   string `json:"address" bson:"address"`
}
