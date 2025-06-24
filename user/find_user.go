package user

type FindUserUseCase struct {
	repo *UserRepository
}

var findUserUseCase *FindUserUseCase

func init() {
	userRepository := GetUserRepository()
	findUserUseCase = &FindUserUseCase{
		repo: userRepository,
	}
}

func GetFindUserUseCase() *FindUserUseCase {
	return findUserUseCase
}

func (uc *FindUserUseCase) FindUserById(id string) (*UserDTO, error) {
	user, err := uc.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
