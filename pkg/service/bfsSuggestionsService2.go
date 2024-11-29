package service

import (
	"fmt"
	"social-network-algorithm/internal/model"
	"sort"
)

// BFS busca o menor caminho no grafo e calcula sugestões.
func BFS(graph *model.Graph, startUserID int) []model.UserSuggestion {
	// Mapas para armazenar distâncias e conexões comuns.
	distance := make(map[int]int)
	commonConnections := make(map[int]int)

	// Inicializa todos os usuários como não visitados.
	for _, user := range graph.Users {
		distance[user.ID] = -1
	}

	// Fila para o BFS.
	queue := []int{startUserID}
	distance[startUserID] = 0

	// Realiza a BFS.
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Percorre os vizinhos do nó atual.
		for neighbor := graph.AdjList[current]; neighbor != nil; neighbor = neighbor.Next {
			neighborID := neighbor.User.ID

			// Ignora o usuário inicial ou conexões diretas.
			if distance[neighborID] == -1 || distance[neighborID] > distance[current]+1 {
				distance[neighborID] = distance[current] + 1
				queue = append(queue, neighborID)
			}

			// Atualiza conexões comuns para sugestões.
			if distance[neighborID] > 1 {
				commonConnections[neighborID]++
			}
		}
	}

	// Monta as sugestões.
	suggestions := []model.UserSuggestion{}
	for userID, dist := range distance {
		// Ignora o próprio usuário e conexões diretas.
		if userID == startUserID || dist <= 1 {
			continue
		}
		suggestions = append(suggestions, model.UserSuggestion{
			UserID:            userID,
			CommonConnections: commonConnections[userID],
			ShortestPath:      dist,
		})
	}

	// Ordena as sugestões por caminho mais curto e depois por conexões comuns.
	sort.Slice(suggestions, func(i, j int) bool {
		if suggestions[i].ShortestPath == suggestions[j].ShortestPath {
			return suggestions[i].CommonConnections > suggestions[j].CommonConnections
		}
		return suggestions[i].ShortestPath < suggestions[j].ShortestPath
	})

	return suggestions
}

// GetSuggestedConnections usa o grafo para sugerir conexões.
func GetSuggestedConnections2(graph *model.Graph, userID int) ([]model.UserSuggestion, error) {
	// Verifica se o usuário existe no grafo.
	if _, exists := graph.VisitedUsers[userID]; !exists {
		return nil, fmt.Errorf("usuário com ID %d não existe no grafo", userID)
	}

	// Executa o BFS para calcular as sugestões.
	suggestions := BFS(graph, userID)
	return suggestions, nil
}
