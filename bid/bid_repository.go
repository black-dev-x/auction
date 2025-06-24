package bid

import (
	"github.com/black-dev-x/auction/auction"
	"github.com/black-dev-x/auction/database"
	errors "github.com/black-dev-x/auction/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BidRepository struct {
	collection *mongo.Collection
}

var bidRepository *BidRepository

func init() {
	db, _ := database.DBConnection()
	collection := db.Collection("bids")
	bidRepository = &BidRepository{collection: collection}
}

func GetBidRepository() *BidRepository {
	return bidRepository
}

func (r *BidRepository) FindBidById(id string) (*BidDTO, error) {
	var bid BidEntity
	err := r.collection.FindOne(nil, bson.M{"_id": id}).Decode(&bid)
	if err == mongo.ErrNoDocuments {
		return nil, errors.NotFoundError("Bid not found")
	}
	return bid.ToDTO(), err
}

func (r *BidRepository) CreateBid(bid *BidDTO) (string, error) {
	bidEntity := bid.ToEntity()
	auctionFound, err := auction.GetAuctionRepository().FindAuctionById(bidEntity.AuctionId)
	if err != nil || auctionFound.Status != auction.AuctionStatusActive {
		return "", errors.BadRequestError("Auction not found")
	}
	data, err := r.collection.InsertOne(nil, bidEntity)
	if err != nil {
		return "", errors.InternalServerError("Failed to create bid")
	}
	return data.InsertedID.(string), nil
}

func (r *BidRepository) FindBidByAuctionId(auctionId string) ([]*BidDTO, error) {
	cursor, err := r.collection.Find(nil, bson.M{"auction_id": auctionId})
	if err != nil {
		return nil, errors.InternalServerError("Failed to find bids")
	}
	defer cursor.Close(nil)

	var bids []*BidEntity
	cursor.All(nil, &bids)

	var result []*BidDTO
	for _, bid := range bids {
		result = append(result, bid.ToDTO())
	}
	return result, nil
}

func (r *BidRepository) FindWinningBidByAuctionId(auctionId string) (*BidDTO, error) {

	var bid BidEntity
	findOne := options.FindOne().SetSort(bson.D{primitive.E{Key: "amount", Value: -1}})
	err := r.collection.FindOne(nil, bson.M{"auction_id": auctionId}, findOne).Decode(&bid)

	if err == mongo.ErrNoDocuments {
		return nil, errors.NotFoundError("No winning bid found")
	}
	return bid.ToDTO(), nil
}
