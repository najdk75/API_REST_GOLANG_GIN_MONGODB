package models

type User struct {
	ID         string   `bson:"_id,omitempty" json:"id,omitempty"`
	Password   string   `bson:"password,omitempty" json:"password,omitempty"`
	Name       string   `bson:"name,omitempty" json:"name,omitempty"`
	Isactive   bool     `bson:"isactive,omitempty" json:"isactive,omitempty"`
	Balance    string   `bson:"balance,omitempty" json:"balance,omitempty"`
	Age        uint     `bson:"age,omitempty" json:"age,omitempty"`
	Gender     string   `bson:"gender,omitempty" json:"genr,omitempty"`
	Company    string   `bson:"company,omitempty" json:"company,omitempty"`
	Email      string   `bson:"email,omitempty" json:"email,omitempty"`
	Phone      string   `bson:"phone,omitempty" json:"phone,omitempty"`
	Address    string   `bson:"address,omitempty" json:"address,omitempty"`
	About      string   `bson:"about,omitempty" json:"about,omitempty"`
	Registered string   `bson:"registered,omitempty" json:"registered,omitempty"`
	Latitude   float64  `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude  float64  `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Tags       []string `bson:"tags,omitempty" json:"tags,omitempty"`
	Friends    []struct {
		Id   int    `bson:"_id,omitempty" json:"id,omitempty"`
		Name string `bson:"name,omitempty" json:"name,omitempty"`
	} ` json:"friends,omitempty" bson:"friends,omitempty"`
	Data string `bson:"data,omitempty" json:"data,omitempty"`
}
