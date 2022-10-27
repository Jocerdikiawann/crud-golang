package roles

type Roles struct {
	Id   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

type RolesReq struct {
	Name string `json:"name" bson:"name"`
}
