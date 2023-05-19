package intface

type I_Brand struct {
	Detail
	UserAccess []User
}

type Detail struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type I_NewData struct {
	Name       string `binding:"required"`
	Age        int    `binding:"required,number"`
	BirthPlace string `json:"birth_place"`
}
