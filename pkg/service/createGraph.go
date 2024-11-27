package service

import "social-network-algorithm/internal/model"

// Função para criar um novo grafo
func CreateGraph(numUsers int) model.Graph {
	newGraph := &model.Graph{
		Users:        make([]*model.User, numUsers),
		AdjList:      make([]*model.AdjacencyNode, numUsers),
		NumUsers:     numUsers,
		VisitedUsers: make(map[int]bool),
	}

	for i := 0; i < numUsers; i++ {
		newGraph.Users[i] = nil
		newGraph.AdjList[i] = nil
	}

	return *newGraph
}
