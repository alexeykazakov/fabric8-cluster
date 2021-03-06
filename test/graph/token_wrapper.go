package graph

import (
	tokenRepo "github.com/fabric8-services/fabric8-cluster/authorization/token/repository"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

// tokenWrapper represents an RPT Token domain object
type tokenWrapper struct {
	baseWrapper
	token *tokenRepo.Token
}

func newTokenWrapper(g *TestGraph, params []interface{}) interface{} {
	w := tokenWrapper{baseWrapper: baseWrapper{g}}

	w.token = &tokenRepo.Token{}

	var identityID *uuid.UUID

	for i := range params {
		switch t := params[i].(type) {
		case *userWrapper:
			identityID = &t.Identity().ID
		case userWrapper:
			identityID = &t.Identity().ID
		}
	}

	if identityID != nil {
		w.token.IdentityID = *identityID
	} else {
		w.token.IdentityID = w.graph.CreateUser().Identity().ID
	}

	err := g.app.TokenRepository().Create(g.ctx, w.token)
	require.NoError(g.t, err)

	return &w
}

func (w *tokenWrapper) Token() *tokenRepo.Token {
	return w.token
}

func (w *tokenWrapper) TokenID() uuid.UUID {
	return w.token.TokenID
}
