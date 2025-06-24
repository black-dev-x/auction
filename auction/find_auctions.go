package auction

type FindAuctionsUseCase struct {
	repo *AuctionRepository
}

var findAuctionsUseCase *FindAuctionsUseCase

func init() {
	auctionRepository := GetAuctionRepository()
	findAuctionsUseCase = &FindAuctionsUseCase{
		repo: auctionRepository,
	}
}

func GetFindAuctionsUseCase() *FindAuctionsUseCase {
	return findAuctionsUseCase
}

func (uc *FindAuctionsUseCase) FindAuctions(status string, category string, productName string) ([]*AuctionDTO, error) {
	auctions, err := uc.repo.FindAuctions(status, category, productName)
	if err != nil {
		return nil, err
	}
	return auctions, nil
}
