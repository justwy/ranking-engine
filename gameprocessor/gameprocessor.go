package gameprocessor

import "github.com/justwy/ranking-engine/models"

// GameResultReader reads game results from a uri.
// Note this is not stream reader. It loads all bytes behind the uri all into memory
// TODO - Need to change to stream reader if input is large
type GameResultReader interface {
	read(uri string) []models.GameResult
}
