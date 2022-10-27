package category

type Category struct {
	Id   string `json:"id" bson:"email,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
}
