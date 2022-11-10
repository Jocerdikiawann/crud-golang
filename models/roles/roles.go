package roles

type Roles struct {
	ID   int `json:"id" gorm:"<-:false;primaryKey"`
	Name string
}
