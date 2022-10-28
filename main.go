package main

import (
	"github.com/joelpatel/monsterslayer/actions"
	"github.com/joelpatel/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	// general game logic

	// start the game
	startGame()
	// start a new round
	//	   loop that keeps on looping as long as we do not have a winner
	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	// end game step
	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int // default value = 0
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		// SPECIAL_ATTACK
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()
	// actions.CurrentPlayerHealth
	// actions.CurrentMonsterHealth

	playerHealth, monsterHealth := actions.GetHealthAmount()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
