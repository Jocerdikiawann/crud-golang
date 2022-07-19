package productsdomain

import categorydomain "belajar-golang-rest-api/models/domain/categoryDomain"

type Products struct {
	Id            string                  `json:"_id" bson:"_id,omitempty"`
	NameOfProduct string                  `json:"nameOfProduct" bson:"nameOfProduct,omitempty"`
	Stock         int64                   `json:"stock" bson:"stock,omitempty"`
	Category      categorydomain.Category `json:"category" bson:"category,omitempty"`
}
