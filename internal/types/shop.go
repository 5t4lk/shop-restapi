package types

import "errors"

type CreateProduct struct {
	Id          string  `json:"_id" bson:"_id"`
	Name        string  `json:"name" bson:"name" binding:"required"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
}

type GetProducts struct {
	Id          string  `json:"_id" bson:"_id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
	UserId      string  `json:"user_id" bson:"user_id"`
}

type Products struct {
	Products []GetProducts `json:"products" bson:"products"`
}

type UpdateProduct struct {
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Price       *float64 `json:"price" bson:"price"`
	Stock       *int     `json:"stock" bson:"stock"`
}

func (u *UpdateProduct) Validate() error {
	if u.Name == "" && u.Description == "" && u.Price == nil && u.Stock == nil {
		return errors.New("update request is empty")
	}
	return nil
}

type AddToCart struct {
	ProductID string `json:"product_id" bson:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" bson:"quantity" binding:"required"`
}

type RemoveFromCart struct {
	ProductID string `json:"product_id" binding:"required"`
}

type CartProduct struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	UserId    string `json:"user_id" bson:"user_id"`
}

type ShoppingCart struct {
	UserID string        `json:"user_id" bson:"user_id"`
	Items  []CartProduct `json:"items" bson:"items"`
}
