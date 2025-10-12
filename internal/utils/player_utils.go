package utils

import "fmt"

func generateUserIDs(gameId string, playerCount int) []string {
	ids := make([]string, playerCount)
	for i := 0; i < playerCount; i++ {
		ids[i] = fmt.Sprintf("%s-player-%d", gameId, i+1)
	}
	return ids
}
