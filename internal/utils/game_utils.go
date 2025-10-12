package utils

import (
	"fmt"
	"math/rand"
)

func generateGameID() string {
	return fmt.Sprintf("coup-%d", rand.Intn(1000))
}
