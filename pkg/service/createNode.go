package service

import (
	"social-network-algorithm/internal/model"
)

// Função para criar um novo nó
func CreateNode(user *model.User) *model.AdjacencyNode {
	newNode := &model.AdjacencyNode{
		User: user,
		Next: nil,
	}

	return newNode

}
