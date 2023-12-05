package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/hassan-algo/pc-gamification/gamification"
	"github.com/hassan-algo/pc-gamification/types"
)

type Result struct {
	UpToWeek       int
	PlayerComplete map[string]interface{}
}

type Players struct {
	PlayerID    int
	PlayerStats []types.PlayerStats
}

type PlayersByTeam struct {
	AllPlayers []Players
}

func TestMain(t *testing.T) {

	// getting all the teams
	allTeams, err := ReadAllTeams()
	if err != nil {
		t.Error("Error Found", err.Error())
	}

	// getting the data per week
	playersByTeam := PlayersByTeam{}
	playersByTeam.GetTeamStats(allTeams)

	// calculating the records
	results := playersByTeam.CalculateEachWeek()

	fmt.Println("Storing")

	file, err := os.Create("results.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"Level", "MatchesPlayed", "PlayerID", "PlayerName", "Points", "Position", "PositionCategory", "Reward", "UpToWeek"}
	writer.Write(header)

	// Write the results to the CSV file
	id := 0
	for completeResult := range results {
		result := completeResult.PlayerComplete
		record := []string{fmt.Sprintf("%d", result["Level"]), fmt.Sprintf("%d", result["MatchesPlayed"]), fmt.Sprintf("%d", result["PlayerID"]), fmt.Sprintf("%s", result["PlayerName"]), fmt.Sprintf("%f", result["Points"]), fmt.Sprintf("%s", result["Position"]), fmt.Sprintf("%s", result["PositionCategory"]), fmt.Sprintf("%f", result["Reward"]), fmt.Sprintf("%d", completeResult.UpToWeek)}
		writer.Write(record)
		fmt.Println("[", id+1, "/", len(results), "] saved!")
		id++
	}

	fmt.Println("Results written to results.csv")

}

func (p *PlayersByTeam) CalculateEachWeek() chan Result {
	allPlayers := p.AllPlayers

	var wg sync.WaitGroup
	numGoroutines := len(allPlayers)
	results := make(chan Result, numGoroutines)
	fmt.Println("Worker Started")
	weeks := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for pid, player := range allPlayers {
		fmt.Println("Calculting for ", pid+1, "of ", len(allPlayers))
		for _, week := range weeks {
			if len(player.PlayerStats) > week {
				newPlayer := player.PlayerStats[:week]
				if newPlayer != nil {
					ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
					wg.Add(1)
					go Calculations(newPlayer, &wg, results, week, ctx, cancel)
				}
			}
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("All goroutines have finished.")
	return results
}

func Calculations(player []types.PlayerStats, wg *sync.WaitGroup, results chan<- Result, upToWeek int, ctx context.Context, cancel context.CancelFunc) {
	defer func(wg *sync.WaitGroup, cancel context.CancelFunc) {
		cancel()
		wg.Done()
	}(wg, cancel)

	gamification := gamification.NewGamification()
	completeStats := gamification.CompleteCalculation(player, "rookie", "")
	myResult := Result{
		UpToWeek:       upToWeek,
		PlayerComplete: completeStats,
	}
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("WOrk has completed")
		results <- myResult
	case <-ctx.Done():
		fmt.Println("Canceled due to timeout")
		var abc map[string]interface{}
		myResult := Result{
			UpToWeek:       upToWeek,
			PlayerComplete: abc,
		}
		results <- myResult
	}
}

func (p *PlayersByTeam) GetTeamStats(allTeams []string) {
	weeks := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	for tid, team := range allTeams {
		fmt.Printf("[%d/%d] WORKING ON %s\n", tid+1, len(allTeams), team)
		for _, week := range weeks {
			URL := "https://api.sportsdata.io/v3/nfl/stats/json/PlayerGameStatsByTeam/2022/" + strconv.Itoa(week) + "/" + team + "?key=" + API_KEY
			res, err := http.Get(URL)
			if err != nil {
				log.Fatal(err)
			}

			var players []types.PlayerStats
			if err = json.NewDecoder(res.Body).Decode(&players); err != nil {
				log.Fatal(err)
			}

			for _, player := range players {
				PlayerExists := p.ContainsPlayer(player.PlayerID)
				if PlayerExists == -1 {
					p.AllPlayers = append(p.AllPlayers, Players{PlayerID: player.PlayerID})
					PlayerExists = len(p.AllPlayers) - 1
				}
				p.AllPlayers[PlayerExists].PlayerStats = append(p.AllPlayers[PlayerExists].PlayerStats, player)
			}
		}
	}
}

func (p PlayersByTeam) ContainsPlayer(pId int) int {
	for id, players := range p.AllPlayers {
		if players.PlayerID == pId {
			return id
		}
	}
	return -1
}

func ReadAllTeams() ([]string, error) {
	var allTeams []string

	file, err := os.Open("nfl_teams.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err != nil {
			break
		}

		allTeams = append(allTeams, record[2])
	}
	return allTeams, nil
}
