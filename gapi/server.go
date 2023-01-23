package gapi

import (
	"fmt"
	db "github.com/4johndoe/simplebank/db/sqlc"
	"github.com/4johndoe/simplebank/pb"
	"github.com/4johndoe/simplebank/token"
	"github.com/4johndoe/simplebank/util"
)

// Server serves gRPC requests for bonking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
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
