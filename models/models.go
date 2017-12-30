package models

import "time"

// Team represents the entity the ranking engine ranks on.
// It could be a single player or a team with multiple players.
type Team struct {
	Id string
	Score float64
}

type Tournament struct {
	Id string
	Description string
}

// GameResult stores the result of the game.
type GameResult struct {
	Time time.Time
	Winner Team
	Loser Team
	Tournament Tournament
}
