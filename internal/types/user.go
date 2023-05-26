package types

type User struct {
	Id       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name" binding:"required"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type UserSignIn struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
