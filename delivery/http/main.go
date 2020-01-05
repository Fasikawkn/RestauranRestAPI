package main

import (
	"log"
	"net/http"

	"github.com/fasikawkn/sample-restaurant-rest-api-master/comment/repository"
	"github.com/fasikawkn/sample-restaurant-rest-api-master/comment/service"
	"github.com/fasikawkn/sample-restaurant-rest-api-master/delivery/http/handler"
	userRep "github.com/fasikawkn/sample-restaurant-rest-api-master/user/repository"
	userSrv "github.com/fasikawkn/sample-restaurant-rest-api-master/user/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:user1@localhost/restsample?sslmode=disable")
	if err != nil {
		log.Println("Failed to connect to the database")
		panic(err)
	}
	log.Println("Database Connection estatblished")
	defer dbconn.Close()
	// roleRepo := repository.NewCommentGormRepo(dbconn)
	// roleSrv := service.NewRoleService(roleRepo)

	//errs := dbconn.CreateTable(entity.Fine{}).GetErrors()

	//if len(errs) > 0 {
	//	panic(errs)
	//}
	userRepo := userRep.NewUserGormRepo(dbconn)
	userSrv := userSrv.NewUserService(userRepo)

	commentRepo := repository.NewCommentGormRepo(dbconn)
	commentSrv := service.NewCommentService(commentRepo)

	adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)
	adminUserHandler := handler.NewAdminUserHandler(userSrv)

	router := httprouter.New()

	router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	router.POST("/v1/admin/comments", adminCommentHandler.PostComment)
	router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)

	router.GET("/v1/admin/users/:id", adminUserHandler.GetUser)
	router.GET("/v1/admin/users/", adminUserHandler.GetUsers)
	router.PUT("/v1/admin/users/:id", adminUserHandler.PutUser)
	router.POST("/v1/admin/users/", adminUserHandler.PostUser)
	router.DELETE("/v1/admin/users/:id", adminUserHandler.DeleteUser)

	http.ListenAndServe(":8181", router)
}
