package auction

type CreateAuctionUseCase struct {
	repo *AuctionRepository
}

var createAuctionUseCase *CreateAuctionUseCase

func init() {
	auctionRepository := GetAuctionRepository()
	createAuctionUseCase = &CreateAuctionUseCase{
		repo: auctionRepository,
	}
}

func GetCreateAuctionUseCase() *CreateAuctionUseCase {
	return createAuctionUseCase
}

func (uc *CreateAuctionUseCase) CreateAuction(auction *AuctionDTO) (string, error) {
	return uc.repo.CreateAuction(auction)
}
