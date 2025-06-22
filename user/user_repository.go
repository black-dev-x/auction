package user

import (
	"github.com/black-dev-x/auction/database"
	errors "github.com/black-dev-x/auction/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	db, _ := database.DBConnection()
	collection := db.Collection("users")
	return &UserRepository{collection: collection}
}

func (r *UserRepository) FindUserById(id string) (*User, error) {
	var user UserEntity
	err := r.collection.FindOne(nil, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.NotFoundError("User not found")
	}
	return user.ToModel(), err
}
