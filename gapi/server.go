package gapi

import (
	"fmt"

	db "github.com/emzhofb/simplebank/db/sqlc"
	"github.com/emzhofb/simplebank/pb"
	"github.com/emzhofb/simplebank/token"
	"github.com/emzhofb/simplebank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
