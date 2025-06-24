package bid

import "time"

type BidDTO struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type BidEntity struct {
	ID        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type AuctionStatus int

func (u *BidDTO) ToEntity() BidEntity {
	return BidEntity{
		ID:        u.ID,
		UserId:    u.UserId,
		AuctionId: u.AuctionId,
		Amount:    u.Amount,
		Timestamp: u.Timestamp.Unix(),
	}
}

func (u *BidEntity) ToDTO() *BidDTO {
	return &BidDTO{
		ID:        u.ID,
		UserId:    u.UserId,
		AuctionId: u.AuctionId,
		Amount:    u.Amount,
		Timestamp: time.Unix(u.Timestamp, 0),
	}
}
