package gateway

import (
	"context"
	"fmt"
)

// App struct
type UserController struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewUserController() *UserController {
	return &UserController{}
}

func (a *UserController) Startup(ctx context.Context) {
	a.ctx = ctx
}
func (a *UserController) Login(name string, password string) string {
	fmt.Println("name: ", name)
	if name == "admin" && password == "123456" {
		return "token123456"
	}
	return ""
}
func (a *UserController) Me(name string, password string) string {
	fmt.Println("name: ", name)
	if name == "admin" && password == "123456" {
		return "token123456"
	}
	return ""
}
