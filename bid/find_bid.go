package bid

type FindBidsUseCase struct {
	repo *BidRepository
}

var findBidsUseCase *FindBidsUseCase

func init() {
	bidRepository := GetBidRepository()
	findBidsUseCase = &FindBidsUseCase{
		repo: bidRepository,
	}
}

func GetFindBidsUseCase() *FindBidsUseCase {
	return findBidsUseCase
}

func (uc *FindBidsUseCase) FindBidById(id string) (*BidDTO, error) {
	bid, err := uc.repo.FindBidById(id)
	if err != nil {
		return nil, err
	}
	return bid, nil
}
