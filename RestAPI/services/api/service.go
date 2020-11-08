package api

type Service interface {
	User() UserService
	Station() StationService
	Comment() CommentService
}
