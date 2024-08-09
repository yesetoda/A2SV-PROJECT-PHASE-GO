package thehttp

import (
	"example/clean/internal/usecases"

	"github.com/gin-gonic/gin"
)


type UserController struct{
	useCase usecases.UserUseCase
}

func (uc *UserController) CreateUser(ctx *gin.Context){
	// Parse request, call use case, and respond

}