package store

import (
	"db-sharding/internal/user"
	"errors"
)

type Store struct {
	users map[int]user.User
}

func NewStore() *Store {
	return &Store{
		users: make(map[int]user.User),
	}
}

func (s *Store) Create(newUser user.User) error {
	if newUser.ID < 0 {
		return ErrInvalidUserID
	}

	_, exists := s.users[newUser.ID]

	if exists {
		return ErrDuplicateUserID
	}

	s.users[newUser.ID] = newUser

	return nil
}

func (s *Store) Get(userID int) (user.User, error) {
	if userID < 0 {
		return user.User{}, ErrInvalidUserID
	}

	selectedUser, exists := s.users[userID]

	if !exists {
		return user.User{}, ErrUserNotFound
	}

	return selectedUser, nil
}

func (s *Store) Update(userID int, newUserData user.User) (user.User, error) {
	if userID < 0 {
		return user.User{}, ErrInvalidUserID
	}

	_, exists := s.users[userID]

	if !exists {
		return user.User{}, ErrUserNotFound
	}

	s.users[userID] = newUserData

	return newUserData, nil
}

func (s *Store) Delete(userID int) (user.User, error) {
	if userID < 0 {
		return user.User{}, ErrInvalidUserID
	}

	selectedUser, exists := s.users[userID]

	if !exists {
		return user.User{}, ErrUserNotFound
	}

	delete(s.users, userID)

	return selectedUser, nil
}

var (
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrDuplicateUserID = errors.New("duplicate user ID")
	ErrUserNotFound    = errors.New("user not found")
)