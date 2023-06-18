package rate_user

type UserRating struct {
	Id        string
	RatedUser string
	RatedBy   string
	Rate      int32
	Status    Status
}

type Status int8

const (
	Pending Status = iota
	Approved
	Canceled
)

type RateUserCommandType int8

const (
	CheckUserVisit RateUserCommandType = iota
	CheckUserRateExists
	ApproveRating
	CancelRating
	UnknownCommand
)

type RateUserCommand struct {
	Rating UserRating
	Type   RateUserCommandType
}

type RateUserReplyType int8

const (
	UserRateDoesntExists RateUserReplyType = iota
	UserRateExists
	UserVisitedHost
	UserNotVisitedHost
	RatingApproved
	RatingCanceled
	UnknownReply
)

type RateUserReply struct {
	Rating UserRating
	Type   RateUserReplyType
}
