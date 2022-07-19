package categorydomain

type Category struct {
	Id             string `json:"_id" bson:"_id,omitempty"`
	NameOfCategory string `json:"nameOfCategory" bson:"nameOfCategory,omitempty"`
}
