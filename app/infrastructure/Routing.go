package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/attumkwg/mplatmm/app/controller"
)

// Routing はRouting用の定義です。
type Routing struct {
	DB *DB
	Gin *gin.Engine
	Port string
}

// NewRouting はRoutingを新たに作成します。
func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB: db,
		Gin: gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controller.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func (c *gin.Context) { usersController.Get(c) })
}

// Run はRoutingを実行します。
func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
