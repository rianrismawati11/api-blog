package app

import (
	"api-blog/controller"
	"api-blog/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories/", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	
	router.GET("/api/posts", categoryController.FindAll)
	router.GET("/api/posts/:categoryId", categoryController.FindById)
	router.POST("/api/posts/", categoryController.Create)
	router.PUT("/api/posts/:categoryId", categoryController.Update)
	router.DELETE("/api/posts/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
