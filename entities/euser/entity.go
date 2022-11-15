package euser

type Entity struct {
	username string
	password string
}

func (e Entity) Username() string { return e.username }

type DTO struct {
	Username string
	Password string
}

func Parse(dto DTO) *Entity {
	return &Entity{
		username: dto.Username,
		password: dto.Password,
	}
}

func (e Entity) PasswordIs(password string) bool { return e.password == password }
