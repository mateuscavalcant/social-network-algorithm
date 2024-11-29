package controller

import (
	"net/http"
	"social-network-algorithm/pkg/service"

	"github.com/gin-gonic/gin"
)

func SuggestedConnectionsHandler(c *gin.Context) {
	// Obtemos o ID do usuário do contexto.
	userID, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário não encontrado na sessão"})
		return
	}

	suggestionService := service.NewSuggestionService()

	// Calcula sugestões.
	suggestions, err := suggestionService.GetSuggestedConnections(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"suggested_connections": suggestions})
}
