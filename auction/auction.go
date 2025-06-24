package auction

import "time"

type AuctionDTO struct {
	ID            string    `json:"id"`
	ProductName   string    `json:"product_name"`
	Category      string    `json:"category"`
	Description   string    `json:"description"`
	Condition     string    `json:"condition"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
	StartingPrice float64   `json:"starting_price"`
}

type AuctionEntity struct {
	ID            string  `bson:"_id"`
	ProductName   string  `bson:"product_name"`
	Category      string  `bson:"category"`
	Description   string  `bson:"description"`
	Condition     string  `bson:"condition"`
	Status        string  `bson:"status"`
	Timestamp     int64   `bson:"timestamp"`
	StartingPrice float64 `bson:"starting_price"`
}

const StatusActive = "Active"
const StatusInactive = "Inactive"

const ConditionNew = "New"
const ConditionUsed = "Used"
const ConditionRefurbished = "Refurbished"

type AuctionStatus int

func (u *AuctionDTO) ToEntity() AuctionEntity {
	return AuctionEntity{
		ID:            u.ID,
		ProductName:   u.ProductName,
		Category:      u.Category,
		Description:   u.Description,
		Condition:     u.Condition,
		Status:        u.Status,
		Timestamp:     u.Timestamp.Unix(),
		StartingPrice: u.StartingPrice,
	}
}

func (u *AuctionEntity) ToDTO() *AuctionDTO {
	return &AuctionDTO{
		ID:            u.ID,
		ProductName:   u.ProductName,
		Category:      u.Category,
		Description:   u.Description,
		Condition:     u.Condition,
		Status:        u.Status,
		StartingPrice: u.StartingPrice,
		Timestamp:     time.Unix(u.Timestamp, 0),
	}
}
