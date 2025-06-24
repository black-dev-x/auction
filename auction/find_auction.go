package auction

type FindAuctionUseCase struct {
	repo *AuctionRepository
}

var findAuctionUseCase *FindAuctionUseCase

func init() {
	auctionRepository := GetAuctionRepository()
	findAuctionUseCase = &FindAuctionUseCase{
		repo: auctionRepository,
	}
}

func GetFindAuctionUseCase() *FindAuctionUseCase {
	return findAuctionUseCase
}

func (uc *FindAuctionUseCase) FindAuctionById(id string) (*AuctionDTO, error) {
	auction, err := uc.repo.FindAuctionById(id)
	if err != nil {
		return nil, err
	}
	return auction, nil
}
