package userReq

type Login struct {
	Account  string `form:"account" binding:"required,gte=4,lte=18"`
	Password string `form:"password" binding:"required,gte=6,lte=18"`
}

type Add struct {
	Name     string `form:"name" binding:"required,gte=2,lte=18"`
	Email    string `form:"email" binding:"required,email,gte=6,lte=50"`
	Account  string `form:"account" binding:"required,gte=4,lte=18"`
	Password string `form:"password" binding:"required,gte=6,lte=18"`
}

type Search struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

type Detail struct {
	ID string `form:"id" binding:"required"`
}

type Update struct {
	ID    string `form:"id" binding:"required"`
	Name  string `form:"name" binding:"required,gte=2,lte=18"`
	Email string `form:"email" binding:"required,email,gte=6,lte=50"`
}

type Delete struct {
	ID string `form:"id" binding:"required"`
}
