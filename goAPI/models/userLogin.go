package models

type UserLogin struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}
