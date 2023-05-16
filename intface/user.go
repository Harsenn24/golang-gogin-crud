package intface

type I_user struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Birthday string `binding:"required"`
}

type I_userMongo struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Birthday int    `binding:"required,number"`
}
