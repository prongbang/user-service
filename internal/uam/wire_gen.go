// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package uam

import (
	"github.com/prongbang/uam-service/internal/pkg/casbinx"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/prongbang/uam-service/internal/uam/interceptor"
	"github.com/prongbang/uam-service/internal/uam/middleware"
	"github.com/prongbang/uam-service/internal/uam/service/auth"
	"github.com/prongbang/uam-service/internal/uam/service/forgot"
	"github.com/prongbang/uam-service/internal/uam/service/permissions"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"github.com/prongbang/uam-service/internal/uam/service/user_creator"
	"github.com/prongbang/uam-service/internal/uam/service/user_role"
)

// Injectors from wire.go:

func New(dbDriver database.Drivers, casbinXs casbinx.CasbinXs) Services {
	dataSource := auth.NewDataSource(dbDriver)
	repository := auth.NewRepository(dataSource)
	roleDataSource := role.NewDataSource(dbDriver)
	roleRepository := role.NewRepository(roleDataSource)
	useCase := role.NewUseCase(roleRepository)
	userDataSource := user.NewDataSource(dbDriver)
	user_creatorDataSource := user_creator.NewDataSource(dbDriver)
	userRepository := user.NewRepository(userDataSource, user_creatorDataSource)
	permissionsUseCase := permissions.NewUseCase(casbinXs)
	userUseCase := user.NewUseCase(userRepository, permissionsUseCase)
	authUseCase := auth.NewUseCase(repository, useCase, userUseCase)
	handler := auth.NewHandler(authUseCase)
	validate := auth.NewValidate()
	router := auth.NewRouter(handler, validate)
	forgotDataSource := forgot.NewDataSource(dbDriver)
	forgotRepository := forgot.NewRepository(forgotDataSource)
	forgotUseCase := forgot.NewUseCase(forgotRepository)
	forgotHandler := forgot.NewHandler(forgotUseCase)
	forgotValidate := forgot.NewValidate()
	forgotRouter := forgot.NewRouter(forgotHandler, forgotValidate)
	roleHandler := role.NewHandler(useCase)
	roleValidate := role.NewValidate()
	roleRouter := role.NewRouter(roleHandler, roleValidate)
	userHandler := user.NewHandler(userUseCase)
	userValidate := user.NewValidate()
	userRouter := user.NewRouter(userHandler, userValidate)
	user_roleDataSource := user_role.NewDataSource(dbDriver)
	user_roleRepository := user_role.NewRepository(user_roleDataSource)
	user_roleUseCase := user_role.NewUseCase(user_roleRepository)
	user_roleHandler := user_role.NewHandler(user_roleUseCase)
	user_roleValidate := user_role.NewValidate()
	user_roleRouter := user_role.NewRouter(user_roleHandler, user_roleValidate)
	uamRouters := NewRouters(router, forgotRouter, roleRouter, userRouter, user_roleRouter)
	authMiddleware := middleware.NewAuthMiddleware(casbinXs)
	middlewares := middleware.New(authMiddleware)
	uamAPI := NewAPI(uamRouters, middlewares)
	authInterceptor := interceptor.NewJWEInterceptor(casbinXs)
	interceptors := interceptor.New(authInterceptor)
	authServer := auth.NewServer(userUseCase, authUseCase)
	roleServer := role.NewServer(useCase)
	userServer := user.NewServer(userUseCase, useCase)
	uamListeners := NewListeners(interceptors, authServer, roleServer, userServer)
	grpc := NewGRPC(uamListeners)
	uamServices := NewService(uamAPI, grpc)
	return uamServices
}
