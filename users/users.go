package users

type user struct {
	Id        int
	Name      string
	Age       int
	WorkHours int
	Post      string
	Update    bool
}

var user_1 = user{
	Id:        1,
	Name:      "woeg",
	Age:       13,
	WorkHours: 5,
	Post:      "работяга",
}

var user_2 = user{
	Id:        2,
	Name:      "bebra",
	Age:       18,
	WorkHours: 1000,
	Post:      "Admin",
}

var user_3 = user{
	Id:        3,
	Name:      "антон",
	Age:       16,
	WorkHours: 45,
	Post:      "нищета",
}

var Users = []user{
	user_1,
	user_2,
	user_3,
}
