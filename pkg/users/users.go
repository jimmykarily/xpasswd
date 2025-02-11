// Package users provides primitives for user information on a system.
package users

import (
	"strconv"
)

// User is an interface that represents a user on a system
type User interface {
	// UID returns the user's unique ID
	UID() (int, error)
	// GID returns the user's group ID
	GID() (int, error)
	// Username returns the user's username
	Username() string
	// Password returns the user's password (this is usually not used but we include it for completeness)
	Password() string
	// HomeDir returns the user's home directory
	HomeDir() string
	// Shell returns the user's shell
	Shell() string
	// RealName returns the user's real name
	RealName() string
}

type CommonUser struct {
	uid      string
	gid      string
	username string
	password string
	homeDir  string
	shell    string
	realName string
}

func (u CommonUser) UID() (int, error) {
	return strconv.Atoi(u.uid)
}

func (u CommonUser) GID() (int, error) {
	return strconv.Atoi(u.gid)
}

func (u CommonUser) Username() string {
	return u.username
}

func (u CommonUser) Password() string {
	return u.password
}

func (u CommonUser) HomeDir() string {
	return u.homeDir
}

func (u CommonUser) Shell() string {
	return u.shell
}

func (u CommonUser) RealName() string {
	return u.realName
}

// UserList is an interface that represents a list of users
type UserList interface {
	// Get returns a user from the list by username
	Get(username string) User
	// GetAll returns all users in the list
	GetAll() ([]User, error)
	GenerateUID() int
	LastUID() int
	SetPath(path string)
	Load() error
}

// CommonUserList is a common implementation of UserList
type CommonUserList struct {
	users   []User
	lastUID int
}

// Get checks if a user with the given username exists
func (list CommonUserList) Get(username string) User {
	for _, user := range list.users {
		if user.Username() == username {
			return user
		}
	}
	return nil
}

func (list CommonUserList) LastUID() int {
	return list.lastUID
}

func (list CommonUserList) GenerateUID() int {
	if len(list.users) == 0 {
		return 0
	}
	return list.lastUID + 1
}
