package storage

// Storage ...
type Storage interface {
	User() UserRepository
	Station() StationRepository
	Comment() CommentRepository
}
