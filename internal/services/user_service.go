package services

import (
	"db-sharding/internal/router"
	"db-sharding/internal/user"
)

type UserService struct{
	shardRouter *router.ShardRouter
}

func NewUserService(shardRouter *router.ShardRouter) *UserService{
	return &UserService{
		shardRouter: shardRouter,
	}
}

func (s *UserService) CreateUser(newUser user.User) error {
	selectedShard, err := s.shardRouter.Route(newUser.ID)

	if err != nil {
		return err
	}

	return selectedShard.Store.Create(newUser)
}

func (s *UserService) GetUser(userID int) (user.User, error){
	selectedShard, err := s.shardRouter.Route(userID)

	if err != nil {
		return user.User{}, err
	}

	return selectedShard.Store.Get(userID)
}

func (s *UserService) UpdateUser(userID int, newData user.User) (user.User, error){
	selectedShard, err := s.shardRouter.Route(userID)

	if err != nil {
		return user.User{}, err
	}

	return selectedShard.Store.Update(userID, newData)
}

func (s *UserService) DeleteUser(userID int) (user.User, error){
	selectedShard, err := s.shardRouter.Route(userID)

	if err != nil {
		return user.User{}, err
	}

	return selectedShard.Store.Delete(userID)
}