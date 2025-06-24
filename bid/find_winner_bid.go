package bid

type FindWinnerBidUseCase struct {
	repo *BidRepository
}

var findWinnerBidUseCase *FindWinnerBidUseCase

func init() {
	bidRepository := GetBidRepository()
	findWinnerBidUseCase = &FindWinnerBidUseCase{
		repo: bidRepository,
	}
}

func GetFindWinnerBidUseCase() *FindWinnerBidUseCase {
	return findWinnerBidUseCase
}

func (uc *FindWinnerBidUseCase) FindWinnerBid(auctionId string) (*BidDTO, error) {
	return uc.repo.FindWinningBidByAuctionId(auctionId)
}
