package service

import (
	"social-network-algorithm/internal/model"
	"social-network-algorithm/pkg/repository"

	"sort"
)

type SuggestionService struct {
	userRepo *repository.BFSSuggestionsRepository
}

func NewSuggestionService() *SuggestionService {
	return &SuggestionService{
		userRepo: repository.NewBFSSuggestionsRepository(),
	}
}

// GetSuggestedConnections realiza a BFS e retorna sugestões de conexão.
func (us *SuggestionService) GetSuggestedConnections(userID int) ([]model.UserSuggestion, error) {
	queue := []int{userID}
	visited := make(map[int]bool)
	distance := make(map[int]int)
	suggestions := make(map[int]*model.UserSuggestion)

	visited[userID] = true
	distance[userID] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Obtém conexões diretas do repositório.
		connections, err := us.userRepo.GetConnections(current)
		if err != nil {
			return nil, err
		}

		for _, neighbor := range connections {
			// Se não visitado, adiciona à fila e marca como visitado.
			if !visited[neighbor] {
				visited[neighbor] = true
				distance[neighbor] = distance[current] + 1
				queue = append(queue, neighbor)
			}

			// Evita sugerir conexões diretas ou o próprio usuário.
			if neighbor == userID || distance[neighbor] == 1 {
				continue
			}

			// Atualiza ou inicializa a sugestão.
			if _, exists := suggestions[neighbor]; !exists {
				suggestions[neighbor] = &model.UserSuggestion{
					UserID:            neighbor,
					CommonConnections: 0,
					ShortestPath:      distance[neighbor],
				}
			}
			suggestions[neighbor].CommonConnections++
		}
	}

	// Converte para slice e ordena os resultados.
	result := []model.UserSuggestion{}
	for _, suggestion := range suggestions {
		result = append(result, *suggestion)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].ShortestPath == result[j].ShortestPath {
			return result[i].CommonConnections > result[j].CommonConnections
		}
		return result[i].ShortestPath < result[j].ShortestPath
	})

	return result, nil
}
