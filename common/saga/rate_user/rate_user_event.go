package rate_user

type UserRating struct {
	RatedUser string
	RatedBy   string
	Rate      int32
}

type RateUserCommandType int8

const (
	ApproveRating RateUserCommandType = iota
	CancelRating
	UnknownCommand
)

type RateUserCommand struct {
	Rating UserRating
	Type   RateUserCommandType
}

type RateUserReplyType int8

const (
	RatingApproved RateUserReplyType = iota
	RatingCanceled
	UnknownReply
)

type RateUserReply struct {
	Rating UserRating
	Type   RateUserReplyType
}
