package repo

import "fmt"

//Port
type Response struct {
	Success bool
	Message string
	Code    int
	Items   string
}

type Repository interface {
	GetInfo() (*Response, error)
}

//Adaptor
type ConnectDB struct {
	Username string
	Password string
}

func NewRepo(c ConnectDB) Repository {
	fmt.Printf("NewRepo: Init ConnectDB instance\n")
	return ConnectDB{
		Username: c.Username,
		Password: c.Password,
	}
}

func (c ConnectDB) GetInfo() (*Response, error) {
	fmt.Printf("call: GetInfo check instance and return respose\n")
	fmt.Println(c.Username)
	fmt.Println(c.Password)

	r := Response{
		Success: true,
		Message: "Initials Success",
		Code:    0,
		Items:   "success",
	}

	return &r, nil
}
