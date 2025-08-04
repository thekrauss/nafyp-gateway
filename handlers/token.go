package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thekrauss/nafyp-gateway/internal"
	pbauth "github.com/thekrauss/nafyp-protos/gen/go/auth"
)

type AuthHandlers struct {
	clients *internal.GRPCClients
}

func NewAuthHandler(clients *internal.GRPCClients) *AuthHandlers {
	return &AuthHandlers{clients: clients}
}

func (h *AuthHandlers) ValidateToken(c *gin.Context) {
	// récupére Authorization: Bearer <token>
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// appeler AuthService.ValidateToken
	resp, err := h.clients.Auth.ValidateToken(context.Background(), &pbauth.TokenRequest{
		Token: token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur AuthService: " + err.Error()})
		return
	}

	//  si valide
	if !resp.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"valid": false})
		return
	}

	// retourne OK
	c.JSON(http.StatusOK, gin.H{"valid": true})
}
