package dtos

type CreateUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
