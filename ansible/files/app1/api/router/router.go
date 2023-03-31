package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"app1/api/middlewares"
)

type routerImpl struct {
	router *gin.Engine
	port   string
}

type Router interface {
	Setup()
	GetRouter() *gin.Engine
}

func NewRouter(port string) *routerImpl {
	return &routerImpl{
		router: gin.Default(),
		port:   port,
	}
}

// Router Setup
func (r routerImpl) configure() {
	r.router.Use(middlewares.CORSMiddleware())
	r.routes()
	if err := r.router.Run(r.port); err != nil {
		fmt.Errorf("Unable to start router error: %!v(MISSING)", err)
		panic(err)
	}
}

func (r routerImpl) Setup() {
	r.configure()
}

func (r routerImpl) GetRouter() *gin.Engine {
	return r.router
}
