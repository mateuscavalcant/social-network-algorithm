package repository

import (
	"database/sql"
	"log"
	"social-network-algorithm/config/database"
)

type BFSSuggestionsRepository struct {
	db *sql.DB
}

func NewBFSSuggestionsRepository() *BFSSuggestionsRepository {
	return &BFSSuggestionsRepository{
		db: database.GetDB(),
	}
}

// GetConnections busca conexões diretas de um usuário no banco de dados.
func (ur *BFSSuggestionsRepository) GetConnections(userID int) ([]int, error) {
	rows, err := ur.db.Query("SELECT followTo FROM user_follow WHERE followBy = ?", userID)
	if err != nil {
		log.Println("Erro ao buscar conexões diretas:", err)
		return nil, err
	}
	defer rows.Close()

	var connections []int
	for rows.Next() {
		var followTo int
		if err := rows.Scan(&followTo); err != nil {
			log.Println("Erro ao escanear conexão:", err)
			continue
		}
		connections = append(connections, followTo)
	}
	return connections, nil
}
