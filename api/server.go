package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/karan0033/bank/db/sqlc"
	"github.com/karan0033/bank/token"
	"github.com/karan0033/bank/utils"
)

// Server serves http request to out banking serivce
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// Newserver create new http server and creates routes
func Newserver(config utils.Config, store db.Store) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	fmt.Printf("TokenSymmetricKey: %s, Length: %d\n", config.TokenSymmetricKey, len(config.TokenSymmetricKey))
	if err != nil {
		return nil, fmt.Errorf("Cannot create a token maker : %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: token,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users/login", server.loginUser)
	router.POST("/createaccount", server.createAccount)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.GET("/account", server.listAccount)
	authRoutes.POST("/deleteaccount/:id", server.deleteAccount)
	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.POST("/user", server.createUser)

	server.router = router
}

// starts server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
