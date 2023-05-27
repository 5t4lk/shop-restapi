package types

type CreateProduct struct {
	Id          string  `json:"_id" bson:"_id"`
	Name        string  `json:"name" bson:"name" binding:"required"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
}

type UsersProducts struct {
	Id        int
	UserId    int
	ProductId int
}
