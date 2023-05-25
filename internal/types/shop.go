package types

type ProductCreateInput struct {
	productId   string  `json:"id" bson:"productId"`
	name        string  `json:"name" bson:"name"`
	description string  `json:"description" bson:"description"`
	price       float64 `json:"price" bson:"price"`
	stock       int     `json:"stock" bson:"stock"`
}
