package intface

import "github.com/dgrijalva/jwt-go"

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
	Active   bool
}

type I_Login struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type I_LoginResult struct {
	Email    string
	Password string
	Id       string
	Active   bool
}

type JwtClaim struct {
	Email    string
	Password string
	Id       string
	jwt.StandardClaims
}

type CheckAccount struct {
	Email string
	Id    string
}

type BodyUserUpdate struct {
	Id string `binding:"required"`
}
