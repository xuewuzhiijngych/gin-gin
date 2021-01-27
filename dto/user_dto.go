package dto

import "ych.hub/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// ToUserDto 转换
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
