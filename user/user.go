package user

type UserDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserEntity struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (u *UserDTO) ToEntity() UserEntity {
	return UserEntity{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u *UserEntity) ToDTO() *UserDTO {
	return &UserDTO{
		ID:   u.ID,
		Name: u.Name,
	}
}
