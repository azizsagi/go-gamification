package gamification

import (
	"log"
	"math"
	"path"
	"sort"

	"github.com/hassan-algo/pc-gamification/extras"
	"github.com/hassan-algo/pc-gamification/types"
	"github.com/tealeg/xlsx"
)

type Gamification struct {
	MAX_XP_VALUE   float64
	GOAL_FILE_NAME string
}

func NewGamification() *Gamification {
	return &Gamification{
		MAX_XP_VALUE:   80000,
		GOAL_FILE_NAME: path.Join("NFL_Player_Goals_Updated2.xlsx"),
	}
}

func (g *Gamification) CompleteCalculation(df []types.PlayerStats, playerTier string, userId string) map[string]interface{} {
	num := make([]float64, 0)
	a := 0.0
	for i := 0; i < 101; i++ {
		num = append(num, float64(i)+a)
		a += float64(i)
	}
	levels := make([]int, 0)
	pointsRequired := make([]float64, 0)
	difference := make([]float64, 0)
	maxNum := extras.Max(num)
	for i := 0; i < 100; i++ {
		levels = append(levels, i)
		pointsRequired = append(pointsRequired, math.Round(g.MinMaxNormalization(num[i], maxNum, g.MAX_XP_VALUE)*100)/100)
		if i == 0 {
			difference = append(difference, 0)
		} else {
			difference = append(difference, pointsRequired[i]-pointsRequired[i-1])
		}
	}
	array := make([]float64, 0)
	for _, value := range types.MeanValue {
		array = append(array, value)
	}
	positionRate := make(map[string]float64)
	totalPositionXP := make(map[string]float64)
	for key, _ := range types.MeanValue {
		positionRate[key] = g.CalculateMedian(array) / types.MeanValue[key]
		totalPositionXP[key] = g.MAX_XP_VALUE / positionRate[key]
	}
	positionWisePoints := make(map[string][]float64)
	for position, value := range totalPositionXP {
		levels = make([]int, 0)
		pointsRequired = make([]float64, 0)
		difference = make([]float64, 0)
		maxNum = extras.Max(num)
		for i := 0; i < 99; i++ {
			levels = append(levels, i)
			pointsRequired = append(pointsRequired, math.Round(g.MinMaxNormalization(num[i], maxNum, value)))
			if i == 0 {
				difference = append(difference, 0)
			} else {
				difference = append(difference, pointsRequired[i]-pointsRequired[i-1])
			}
		}
		pointsRequired[98] = pointsRequired[98] - (math.Mod(pointsRequired[98], 100))
		difference[98] = difference[98] - (math.Mod(difference[98], 100))
		positionWisePoints[position] = append([]float64{}, pointsRequired...)
	}
	workbook, err := xlsx.OpenFile(g.GOAL_FILE_NAME)
	if err != nil {
		log.Fatal("FILE NOT FOUND")
	}

	rewards := make(map[string][]map[string]interface{})
	for _, sheet := range workbook.Sheets {
		data := extras.SheetToJson(sheet)
		rewards[sheet.Name] = data
	}

	players := 0.0
	allPositions := ""
	allLevels := 0
	allRewards := 0.0
	allMatches := 0
	allPlayerName := ""
	allPlayerCategories := ""
	allPlayerID := 0
	if len(df) < 1 {
		prevWeek := df[len(df)-1].Week
		if len(df) > 1 {
			prevWeek = df[len(df)-2].Week
		}
		return map[string]interface{}{
			"UserID":           userId,
			"Points":           players,
			"Tier":             playerTier,
			"Position":         allPositions,
			"Week":             df[len(df)-1].Week,
			"PreviousWeek":     prevWeek,
			"Level":            allLevels,
			"Reward":           allRewards,
			"MatchesPlayed":    allMatches,
			"PositionCategory": allPlayerCategories,
			"PlayerName":       allPlayerName,
			"PlayerID":         allPlayerID,
		}

	}
	for id, row := range df {
		deepdf := make([]types.PlayerStats, len(df))
		copy(deepdf, df)
		if id == 0 {
			players = 0
			allPositions = row.Position
			allLevels = 0
			allRewards = 0
			allMatches = 0
			allPlayerCategories = ""
			allPlayerID = row.PlayerID
			allPlayerName = row.Name
		}
		active := false
		for _, z := range types.ColumnData {
			if extras.GetField(row, z).(float64) > 0 && extras.GetField(row, z) != nil && extras.GetField(row, z) != "" {
				active = true
				break
			}
		}
		if active == false {
			continue
		}
		var mainkey string
		for key := range types.PositionsX {
			if g.sliceContains(types.PositionsX[key], row.Position) {
				mainkey = key
				allPlayerCategories = key
				break
			}
		}
		if mainkey != "" {
			points := row.Points
			for _, val := range types.Positions[mainkey] {
				points += extras.GetField(row, val.Key).(float64) * val.Value
				// players = points
			}
			// if extras.GetField(row, "Week").(int) < 6 {
			if row.AppliedBoost {
				// players += points * types.PlayerTiers[playerTier]
				players += (points * 2)
			} else {
				players += points
			}
			// } else {
			// 	players += points
			// }
			allLevels = g.CalculateLevel(players, allPositions, types.PositionsX, positionWisePoints)
			allRewards += g.CalculateReward(rewards[mainkey], row, allLevels, mainkey, types.PositionsX)
			allMatches++
		}
		row.Points = players
	}
	prevWeek := df[len(df)-1].Week
	if len(df) > 1 {
		prevWeek = df[len(df)-2].Week
	}

	return map[string]interface{}{
		"UserID":           userId,
		"Points":           players,
		"Tier":             playerTier,
		"Position":         allPositions,
		"Week":             df[len(df)-1].Week,
		"PreviousWeek":     prevWeek,
		"Level":            allLevels,
		"Reward":           allRewards,
		"MatchesPlayed":    allMatches,
		"PositionCategory": allPlayerCategories,
		"PlayerName":       allPlayerName,
		"PlayerID":         allPlayerID,
	}
}

func (g *Gamification) CalculateMedian(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0 // Handle the case of an empty array
	}

	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)

	middleIndex := len(sortedNumbers) / 2

	if len(sortedNumbers)%2 == 0 {
		// If the array length is even, calculate the average of the middle two values
		return float64(sortedNumbers[middleIndex-1]+sortedNumbers[middleIndex]) / 2
	} else {
		// If the array length is odd, return the middle value
		return float64(sortedNumbers[middleIndex])
	}
}

func (g *Gamification) CalculateReward(rewardsData []map[string]interface{}, data types.PlayerStats, currentLevel int, tierPosition string, positions map[string][]string) float64 {
	reward := 0.0

	for _, row := range rewardsData {
		if d1, ok := row["Goal Type"].(string); ok {
			if d1 == "Game" {
				reward += g.CalculatePointsGame(
					extras.GetField(data, row["Target"].(string)).(float64),
					row["Min Value"].(float64),
					row["Max Value"].(float64),
					tierPosition,
					data.Position,
					row["PCC Reward"].(float64),
					data.Started,
					row["Started"].(int),
					currentLevel,
					positions,
				)
			}
		}
	}

	return reward
}

func (g *Gamification) CalculatePointsGame(target float64, minValue float64, maxValue float64, tierPosition string, playerPosition string, reward float64, started int, startedCondition int, currentLevel int, positions map[string][]string) float64 {
	if extras.Contains(positions[tierPosition], playerPosition) {
		if started == startedCondition || startedCondition == 0 {
			if target >= minValue && target <= maxValue {
				return float64(reward) * g.LevelMultiplier(currentLevel)
			}
		}
	}
	return 0
}

func (g *Gamification) LevelMultiplier(level int) float64 {
	if level <= 30 {
		return 1
	} else if level <= 60 {
		return 1.5
	} else if level <= 80 {
		return 2
	} else if level <= 95 {
		return 2.5
	} else if level <= 98 {
		return 3
	} else {
		return 4
	}
}

func (g *Gamification) CalculateLevel(points float64, position string, positions map[string][]string, positionWisePoints map[string][]float64) int {
	level := 0
	pos := ""

	if points == 0 {
		return 1
	}
	for key, value := range positions {
		for _, val := range value {
			if val == position {
				pos = key
				break
			}
		}
		if pos != "" {
			break
		}
	}
	if pos == "" {
		return -1
	}
	for _, pPoints := range positionWisePoints[pos] {
		if float64(points) < pPoints {
			return level
		}
		level++
	}
	return 99
}

func (g *Gamification) MinMaxNormalization(value float64, max_value float64, max_xp_value float64) float64 {
	return ((value - 0) / (max_value - 0)) * (max_xp_value - 0)
}

func (g *Gamification) sliceContains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
