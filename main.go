package main

import (
	"backend-tes-qoin/structs"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("========== PERMAIAN DADU ==========")
	playerCount := InputInteger("Masukkan jumlah pemain     : ", "Masukan harus bilangan bulat lebih dari 0")
	diceCount := InputInteger("Masukkan jumlah dadu (1-6) : ", "Masukan harus bilangan bulat antara 1-6")
	fmt.Println(fmt.Sprintf("Pemain = %d, Dadu = %d", playerCount, diceCount))
	
	if playerCount == 1 {
		fmt.Println("===================================")
		fmt.Println("Tidak ada pemain yang menang ataupun kalah karena hanya ada 1 pemain.")
	} else {
		players := GameInit(playerCount, diceCount)
		winnerIndex, looserIndex := GameStart(players, diceCount)
	
		fmt.Println("===================================")
		fmt.Println(fmt.Sprintf("Game berakhir karena hanya %s yang memiliki dadu.", players[looserIndex].Name))
		fmt.Println(fmt.Sprintf("Game dimenangkan oleh %s karena memiliki poin lebih banyak dari pemain lainnya.", players[winnerIndex].Name))
	}
}

func InputInteger(placeholder string, alert string) (input int) {
	for {
		fmt.Print(placeholder)
		_, err := fmt.Scanln(&input)
		if err == nil && input > 0{
			return input
		} else {
			fmt.Println(alert)
			fmt.Scanf("%s")
		}
	} 
}

func DiceInit(diceCount int) (dices []int) {
	for i := 0; i < diceCount; i++ {
		dices = append(dices, rand.Intn(6) + 1)
	}
	return dices
}

func RandomPlayerDice(player structs.Player) []int {
	for index, _ := range player.Dice {
		player.Dice[index] = rand.Intn(6) + 1
	}
	return player.Dice
}

func PrintDice(dices []int) (result string) {
	for index, dice := range dices {
		if index == len(dices) - 1 {
			result += fmt.Sprintf("%d", dice)
		} else {
			result += fmt.Sprintf("%d,", dice)
		}
	}

	if result == "" {
		result = "_ (Berhenti bermain karena tidak memiliki dadu)"
	}

	return result
}

func GameInit(playerCount int, diceCount int) (players []structs.Player) {
	for i := 0; i < playerCount; i++ {
		players = append(players, structs.Player{
			Name: fmt.Sprintf("Pemain #%d", i+1),
			Point: 0,
			Dice: DiceInit(diceCount),
			IsActive: true,
		})
	}
	return players
}

func CheckIsGameEnd(players []structs.Player) (check bool, playerActiveIndex int, playerTheMostPointIndex int) {
	playersActiveCount := 0
	maxPoint := 0
	
	for playerIndex, player := range players {
		if player.IsActive {
			playerActiveIndex = playerIndex
			playersActiveCount++
		} else {
			if maxPoint < player.Point {
				playerTheMostPointIndex = playerIndex
			}
		}
	}

	check = playersActiveCount == 1
	return check, playerActiveIndex, playerTheMostPointIndex
}

func NextPlayerIndex(players []structs.Player, playerIndex int) (nextPlayerIndex int) {
	nextPlayerIndex = (playerIndex + 1) % len(players)
	for {
		if players[nextPlayerIndex].IsActive && nextPlayerIndex != playerIndex {
			break
		} else {
			nextPlayerIndex = (nextPlayerIndex + 1) % len(players)
		}
	}
	return nextPlayerIndex
}

func GameStart(players []structs.Player, diceCount int) (looserIndex int, winnerIndex int) {
	gameIndex := 1
	var isGameEnd bool
	for {
		fmt.Println("===================================")
		fmt.Println(fmt.Sprintf("Giliran %d lempar dadu:", gameIndex))
		for playerIndex, player := range players {
			players[playerIndex].Dice = RandomPlayerDice(player)
			fmt.Println(fmt.Sprintf("        %s (%d): %s", player.Name, player.Point, PrintDice(player.Dice)))
		}

		for playerIndex, player := range players {
			if player.IsActive {
				var tempDice []int
				var tempDiceNextPlayer []int
				nextPlayerIndex := NextPlayerIndex(players, playerIndex)
				for j := 0; j < len(player.Dice); j++ {
					switch player.Dice[j] {
					case 1:
						tempDiceNextPlayer = append(tempDiceNextPlayer, 1)
						break
					case 6:
						players[playerIndex].Point += 1
						break
					default:
						tempDice = append(tempDice, player.Dice[j])
						break
					}
					players[playerIndex].Dice = tempDice
				}
				players[nextPlayerIndex].DiceFromPlayer = tempDiceNextPlayer
			}
		}
		
		fmt.Println("Setelah evaluasi:")
		for playerIndex, player := range players {
			players[playerIndex].Dice = append(player.Dice, player.DiceFromPlayer...)
			if len(players[playerIndex].Dice) == 0 {
				players[playerIndex].IsActive = false
				players[playerIndex].DiceFromPlayer = []int{}
			}
			fmt.Println(fmt.Sprintf("        %s (%d): %s", player.Name, player.Point, PrintDice(players[playerIndex].Dice)))
		}

		isGameEnd, looserIndex, winnerIndex = CheckIsGameEnd(players)
		if isGameEnd {
			break
		}
		gameIndex++
	}
	return looserIndex, winnerIndex
}
