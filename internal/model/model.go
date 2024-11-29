package model

type User struct {
	ID       int
	Username string
}

type AdjacencyNode struct {
	User *User
	Next *AdjacencyNode
}

type Graph struct {
	Users        []*User
	AdjList      []*AdjacencyNode
	NumUsers     int
	VisitedUsers map[int]bool
}

type Queue struct {
	Items     []int
	FirstItem int
	LastItem  int
}

type UserSuggestion struct {
	UserID            int
	CommonConnections int
	ShortestPath      int
}
