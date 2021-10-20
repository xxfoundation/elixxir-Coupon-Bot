////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Handles the high level storage API.
// This layer merges the business logic layer and the database layer

package storage

// Params for creating a storage object
type Params struct {
	Username string
	Password string
	DBName   string
	Address  string
	Port     string
}

// Storage struct interfaces with the API for the storage layer
type Storage struct {
	// Stored Database interface
	database
}

// NewStorage creates a new Storage object wrapping a database interface
// Returns a Storage object, and error
func NewStorage(params Params) (*Storage, error) {
	db, err := newDatabase(params.Username, params.Password, params.DBName, params.Address, params.Port)
	storage := &Storage{db}
	return storage, err
}
