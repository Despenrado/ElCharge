package storage

// Storage ...
type Storage interface {
	User() UserRepository
}
