package domain

type Address struct {
	Street  string `json:"street"`
	Number  int32  `json:"number"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Address  Address `json:"address"`
	Role     string  `json:"role"`
}

type Rating struct {
	Rating    int32  `json:"rating"`
	RatedBy   string `json:"ratedBy"`
	RatedHost string `json:"ratedHost"`
}

func MakeRating(rating int32, ratedBy, ratedHost string) Rating {
	return Rating{
		Rating:    rating,
		RatedBy:   ratedBy,
		RatedHost: ratedHost,
	}
}
