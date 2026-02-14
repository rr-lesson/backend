package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
} // @name LoginReq

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
} // @name RegisterReq
