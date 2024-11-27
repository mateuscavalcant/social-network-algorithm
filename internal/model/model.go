package model

type User struct {
	id       int
	username string
}

type AdjacencyNode struct {
	user *User
	next *AdjacencyNode
}

type Graph struct {
	users        **User
	adjList      **AdjacencyNode
	numUsers     int
	visitedUsers *int
}

type Queue struct {
	items     int
	firstItem int
	lastItem  int
}
