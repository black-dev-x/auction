package user

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserEntity struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (u *User) ToEntity() UserEntity {
	return UserEntity{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u *UserEntity) ToModel() *User {
	return &User{
		ID:   u.ID,
		Name: u.Name,
	}
}
