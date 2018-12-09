package day9

import (
	"container/ring"
)

//SolvePart1 returns the answer for part 1 of day 9
func SolvePart1() int {
	return WinnerResult(438, 71626)
}

//SolvePart2 returns the answer for part 2 of day 9
func SolvePart2() int {
	return WinnerResult(438, 71626*100)
}

//WinnerResult returns the result of the elf winning the game
func WinnerResult(playerCount, lastStone int) int {
	playerResults := simulateGame(playerCount, lastStone)

	max := 0
	for _, v := range playerResults {
		if v > max {
			max = v
		}
	}

	return max
}

func simulateGame(playerCount, lastStone int) map[int]int {
	gameState := ring.New(1)
	gameState.Value = 0
	playerResults := make(map[int]int)

	for stone := 1; stone <= lastStone; stone++ {
		if stone%23 == 0 {
			playerID := stone % playerCount
			gameState = gameState.Move(-6)
			stoneAtMinusSeven := gameState.Move(-2).Link(gameState)
			playerResults[playerID] += (stone + stoneAtMinusSeven.Value.(int))
		} else {
			gameState = gameState.Next()
			gameState.Link(intToRing(stone))
			gameState = gameState.Next()
		}
	}

	return playerResults
}

func intToRing(i int) *ring.Ring {
	result := ring.New(1)
	result.Value = i
	return result
}
