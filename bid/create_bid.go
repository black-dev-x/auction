package bid

type CreateBidUseCase struct {
	repo *BidRepository
}

var createBidUseCase *CreateBidUseCase

func init() {
	bidRepository := GetBidRepository()
	createBidUseCase = &CreateBidUseCase{
		repo: bidRepository,
	}
}

func GetCreateBidUseCase() *CreateBidUseCase {
	return createBidUseCase
}

func (uc *CreateBidUseCase) CreateBid(bid *BidDTO) (string, error) {
	return uc.repo.CreateBid(bid)
}
