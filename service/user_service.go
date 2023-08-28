package service

import "github.com/br4tech/go-to-mysql/repository"

type IUserService interface {
  GetUserByID(id int) (*repository.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByID(id int) (*repository.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, err // Não use panic, apenas retorne o erro
	}

	return user, nil
}