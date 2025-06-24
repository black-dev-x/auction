package auction

import (
	"github.com/black-dev-x/auction/database"
	errors "github.com/black-dev-x/auction/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionRepository struct {
	collection *mongo.Collection
}

var auctionRepository *AuctionRepository

func init() {
	db, _ := database.DBConnection()
	collection := db.Collection("auctions")
	auctionRepository = &AuctionRepository{collection: collection}
}

func GetAuctionRepository() *AuctionRepository {
	return auctionRepository
}

func (r *AuctionRepository) FindAuctionById(id string) (*AuctionDTO, error) {
	var auction AuctionEntity
	err := r.collection.FindOne(nil, bson.M{"_id": id}).Decode(&auction)
	if err == mongo.ErrNoDocuments {
		return nil, errors.NotFoundError("Auction not found")
	}
	return auction.ToDTO(), err
}

func (r *AuctionRepository) CreateAuction(auction *AuctionDTO) (string, error) {
	auctionEntity := auction.ToEntity()

	data, err := r.collection.InsertOne(nil, auctionEntity)
	if err != nil {
		return "", errors.InternalServerError("Failed to create auction")
	}
	return data.InsertedID.(string), nil
}

func (r *AuctionRepository) FindAuctions(status string, category string, name string) ([]*AuctionDTO, error) {
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}
	if category != "" {
		filter["category"] = category
	}
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}

	var auctions []*AuctionEntity
	cursor, err := r.collection.Find(nil, filter)
	if err != nil {
		return nil, errors.InternalServerError("Failed to find auctions")
	}
	defer cursor.Close(nil)

	for cursor.Next(nil) {
		var auction AuctionEntity
		if err := cursor.Decode(&auction); err != nil {
			return nil, errors.InternalServerError("Failed to decode auction")
		}
		auctions = append(auctions, &auction)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.InternalServerError("Failed to find auctions")
	}

	var result []*AuctionDTO
	for _, auction := range auctions {
		result = append(result, auction.ToDTO())
	}
	return result, nil
}
