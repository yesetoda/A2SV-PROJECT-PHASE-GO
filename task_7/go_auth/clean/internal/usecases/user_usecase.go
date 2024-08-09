package usecases
// make the data in the desired format between the entities and the 
import (
"example/clean/internal/entities"
"example/clean/internal/interfaces/db"
)


type UserUseCase struct{
	userRepo  db.UserRepository
}

func (uc *UserUseCase) SignUp(username,password string)(*entities.User,error){
	user := &entities.User{
		Username: username,
		Password: password,
	}
	return uc.userRepo.Save(user)
}
func (uc *UserUseCase) Signin(username,password string)(*entities.User,error){
	user := &entities.User{
		Username: username,
		Password: password,
	}
	return uc.userRepo.Save(user)
}
func (uc *UserUseCase) Promote(username,password string)(*entities.User,error){
	user := &entities.User{
		Username: username,
		Password: password,
	}
	return uc.userRepo.Save(user)
}
func (uc *UserUseCase) ViewAll(username,password string)(*entities.User,error){
	user := &entities.User{
		Username: username,
		Password: password,
	}
	return uc.userRepo.Save(user)
}
func (uc *UserUseCase) ViewByUserName(username,password string)(*entities.User,error){
	user := &entities.User{
		Username: username,
		Password: password,
	}
	return uc.userRepo.Save(user)
}