package interaction

import (
	"fmt"
	"os"
	"time"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action.")
	fmt.Println("--------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	// alternatively, make this a struct method
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("--------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("--------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	file, err := os.OpenFile("gamelog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		fmt.Println("Saving a log file failed. Exiting...")
		os.Exit(1)
	}

	header := fmt.Sprintf("New Game Session at time: %v\n", time.Now())
	_, err = file.WriteString(header)

	if err != nil {
		fmt.Println("Writing into log file failed. Exiting...")
		return
	}
	for index, roundData := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                roundData.Action,
			"Player Attack Damage":  fmt.Sprint(roundData.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(roundData.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(roundData.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(roundData.PlayerHealth),
			"Monster Health":        fmt.Sprint(roundData.MonsterHealth),
		}
		// logLine := fmt.Sprintf("New Game Session at time: %v\n%v", time.Now(), logEntry)
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed.")
			continue
		}
	}

	file.Close()
	pwd, _ := os.Getwd()
	fmt.Printf("Successfully wrote data to the log file at: %v/gamelog.txt\n", pwd)
}
