package intface

type I_NewData struct {
	Name       string `binding:"required"`
	Age        int    `binding:"required,number"`
	BirthPlace string `json:"birth_place"`
}