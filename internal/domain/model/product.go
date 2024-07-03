package model

type Product struct {
    ID          string  `json:"id" bson:"_id,omitempty" gorm:"primary_key"`
    Name        string  `json:"name" bson:"name" gorm:"name"`
    Description string  `json:"description" bson:"description" gorm:"description"`
    Price       float64 `json:"price" bson:"price" gorm:"price"`
}
