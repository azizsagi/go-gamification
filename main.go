package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/hassan-algo/pc-gamification/extras"
	"github.com/hassan-algo/pc-gamification/gamification"
	"github.com/hassan-algo/pc-gamification/profile"
	"github.com/hassan-algo/pc-gamification/seasons"
	"github.com/hassan-algo/pc-gamification/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	API_KEY                   string = "f5f5e61bc9d84f1c9f7c5ffc8f66eb4c"
	DB_CONNECTION             string = "mongodb+srv://personal-corner-db:MaH0dp4YgcGknJp7@personalcorner.v0pkq.mongodb.net/pcc-Prod?retryWrites=true&w=majority"
	DB_NAME                   string = "pcc-Prod"
	DB_COLLECTION_GET         string = "lineups"
	DB_COLLECTION_NFT         string = "allStatsForNFT"
	DB_COLLECTION_PLAYERS     string = "allStatsForPlayer"
	DB_COLLECTION_ALL_PLAYERS string = "collections"
)

type Matches struct {
	Matches []types.Match
	Stats   []types.NFTStats
}

func (m Matches) ExistStats(userID string, playerId int, tier string, season string, seasonName string) int {
	for id, stat := range m.Stats {
		if stat.UserID == userID && stat.PlayerID == playerId && stat.Tier == tier && stat.Season == season && stat.SeasonName == seasonName {
			return id
		}
	}
	return -1
}

func NewMatches(results []types.Result) *Matches {
	m := Matches{}
	m.GetMatches(results)
	m.GetPlayerStats()
	go m.SavePlayerStats()
	return &m
}

func (m Matches) ExistMatch(userID string, Team string, Week int, Tier string, Season string) int {
	for id, match := range m.Matches {
		if match.UserId == userID && match.Team == Team && match.Week == Week && match.Tier == Tier && match.Season == Season {
			return id
		}
	}
	return -1
}

func (m Matches) ExistMatchStats(playerID int, tier string) int {
	for id, match := range m.Stats {
		if match.PlayerID == playerID && match.Tier == tier {
			return id
		}
	}
	return -1
}

func (m Matches) PrintMatches() {
	for _, match := range m.Matches {
		fmt.Println("[Match]", match.Team, match.Week, match.Season, match.Boost, match.Tier, match.Players)
	}
}

func (m Matches) PrintStats() {
	for _, match := range m.Stats {
		fmt.Println("[STATS]", match.PlayerID, match.Tier, len(match.Stats))
	}
}

func (m *Matches) GetMatches(results []types.Result) {
	for _, player := range results {
		for _, nft := range player.NFTs {
			for _, week := range nft.Weeks {
				matchExists := m.ExistMatch(player.UserId, nft.Team, week.Week, nft.TierName, week.SeasonName)
				if matchExists == -1 {
					m.Matches = append(m.Matches, types.Match{
						UserId:     player.UserId,
						Team:       nft.Team,
						Week:       week.Week,
						Season:     week.Season,
						SeasonName: week.SeasonName,
						Boost:      week.AppliedBoostValue,
						Tier:       nft.TierName,
					})
					matchExists = len(m.Matches) - 1
				}
				m.Matches[matchExists].Players = append(m.Matches[matchExists].Players, nft.PlayerId)
			}
		}
	}
}

func (m *Matches) GetPlayerStats() {
	for _, match := range m.Matches {

		URL := "https://api.sportsdata.io/v3/nfl/stats/json/PlayerGameStatsByTeam/" + match.SeasonName + "/" + strconv.Itoa(match.Week) + "/" + match.Team + "?key=" + API_KEY
		res, err := http.Get(URL)
		if err != nil {
			log.Println(err)
			continue
		}

		var players []types.PlayerStats
		if err = json.NewDecoder(res.Body).Decode(&players); err != nil {
			log.Println(err)
			continue
		}

		for _, player := range players {
			if extras.ContainsInt(match.Players, player.PlayerID) {
				statExists := m.ExistStats(match.UserId, player.PlayerID, match.Tier, match.Season, match.SeasonName)
				if statExists == -1 {
					m.Stats = append(m.Stats, types.NFTStats{
						UserID:     match.UserId,
						PlayerID:   player.PlayerID,
						Tier:       match.Tier,
						Season:     match.Season,
						SeasonName: match.SeasonName,
					})
					statExists = len(m.Stats) - 1
				}
				newPlayer := player
				newPlayer.AppliedBoost = match.Boost
				m.Stats[statExists].Stats = append(m.Stats[statExists].Stats, newPlayer)
			}
		}
	}
}

func CalculateStats() {
	min := time.Minute
	for {
		iterationStartTime := time.Now()

		data := GetDataFromDB()
		for _, player := range data {
			for _, nft := range player.NFTs {
				sort.Slice(nft.Weeks, func(i, j int) bool {
					return nft.Weeks[i].Week < nft.Weeks[j].Week
				})
			}
		}
		match := NewMatches(data)
		gamification := gamification.NewGamification()
		for _, player := range match.Stats {
			completeStats := gamification.CompleteCalculation(player.Stats, player.Tier, player.UserID)
			SaveGamification(completeStats, player.Season, player.SeasonName)
		}

		elapsedTime := time.Since(iterationStartTime)
		if elapsedTime < min {
			remainingTime := min - elapsedTime
			fmt.Printf("[Player] Sleeping for %v in this iteration\n", remainingTime)

			time.Sleep(remainingTime)
		} else {
			fmt.Println("[Player] Iteration took more than a minute")
		}
	}
}

func CalculateSeason() {
	fiveMin := 5 * time.Minute
	for {
		iterationStartTime := time.Now()
		allPlayers := seasons.NewSeaonStats()
		allPlayers.GetPlayerStats()
		allPlayers.PrintPlayerStats()
		allPlayers.SaveSeasonRecords()
		elapsedTime := time.Since(iterationStartTime)
		if elapsedTime < fiveMin {
			remainingTime := fiveMin - elapsedTime
			fmt.Printf("[SEASON] Sleeping for %v in this iteration\n", remainingTime)

			time.Sleep(remainingTime)
		} else {
			fmt.Println("[SEASON] Iteration took more than 5 minutes")
		}
	}
}

func CalculateProfile() {
	fiveMin := 5 * time.Minute
	for {
		iterationStartTime := time.Now()
		profile.NewAllUsers()
		elapsedTime := time.Since(iterationStartTime)
		if elapsedTime < fiveMin {
			remainingTime := fiveMin - elapsedTime
			fmt.Printf("[PROFILE] Sleeping for %v in this iteration\n", remainingTime)
			time.Sleep(remainingTime)
		} else {
			fmt.Println("[PROFILE] Iteration took more than 5 minutes")
		}
	}
}

func main() {
	go CalculateSeason()
	go CalculateProfile()
	CalculateStats()
}

func SaveGamification(playerCalculation map[string]interface{}, season string, seasonName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error checking MongoDB connection:", err)
	}

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_NFT)
	playerCalculation["Season"] = season
	playerCalculation["SeasonName"] = seasonName
	duplicate := playerCalculation
	if playerCalculation["PreviousWeek"] != playerCalculation["Week"] {
		// find prev week and subtract
		var result bson.M
		if err := collection.FindOne(ctx, bson.M{"SeasonName": seasonName, "Season": season, "UserID": playerCalculation["UserID"], "PlayerID": playerCalculation["PlayerID"], "Tier": playerCalculation["Tier"], "Week": playerCalculation["PreviousWeek"]}).Decode(&result); err == nil && len(result) > 0 {
			dupPoint, _ := duplicate["Points"].(float64)
			resPoint, _ := result["Points"].(float64)
			duplicate["Points"] = dupPoint - resPoint
			dupPoint, _ = duplicate["Reward"].(float64)
			resPoint, _ = result["Reward"].(float64)
			duplicate["Reward"] = dupPoint - resPoint
		}
	}
	// check if already exists
	var result bson.M
	if err := collection.FindOne(ctx, bson.M{"SeasonName": seasonName, "Season": season, "UserID": playerCalculation["UserID"], "PlayerID": playerCalculation["PlayerID"], "Tier": playerCalculation["Tier"], "Week": playerCalculation["Week"]}).Decode(&result); err == nil && len(result) > 0 {
		// update prev
		duplicate["updatedAt"] = time.Now().Format(time.RFC3339)
		collection.UpdateOne(ctx, bson.M{"SeasonName": seasonName, "Season": season, "UserID": playerCalculation["UserID"], "PlayerID": playerCalculation["PlayerID"], "Tier": playerCalculation["Tier"], "Week": playerCalculation["Week"]}, bson.M{"$set": duplicate})
	} else {
		// create New
		duplicate["updatedAt"] = time.Now().Format(time.RFC3339)
		collection.InsertOne(ctx, duplicate)
	}

	client.Disconnect(ctx)
}

func GetDataFromDB() []types.Result {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error checking MongoDB connection:", err)
	}

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_GET)

	pipeline := bson.A{
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$nfts"},
					{"preserveNullAndEmptyArrays", false},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"as", "tierInfo"},
					{"from", "tiers"},
					{"let", bson.D{{"tierId", "$nfts.tier"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$and",
													bson.A{
														bson.D{
															{"$eq",
																bson.A{
																	"$_id",
																	"$$tierId",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							bson.D{{"$project", bson.D{{"name", 1}}}},
						},
					},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$tierInfo"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"as", "player"},
					{"from", "players"},
					{"let", bson.D{{"playerId", "$nfts.playerId"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$match",
									bson.D{
										{"$expr",
											bson.D{
												{"$and",
													bson.A{
														bson.D{
															{"$eq",
																bson.A{
																	"$PlayerID",
																	"$$playerId",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							bson.D{
								{"$project",
									bson.D{
										{"Position", 1},
										{"Team", 1},
									},
								},
							},
							bson.D{{"$limit", 1}},
						},
					},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$player"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"nfts.Season", "$Season"},
					{"nfts.SeasonName", "$SeasonName"},
					{"nfts.Week", "$Week"},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$nfts.nft"},
					{"nftId", bson.D{{"$first", "$nfts.nft"}}},
					{"walletAddress", bson.D{{"$first", "$user"}}},
					{"userId", bson.D{{"$first", "$userId"}}},
					{"playerId", bson.D{{"$first", "$nfts.playerId"}}},
					{"team", bson.D{{"$first", "$player.Team"}}},
					{"tierId", bson.D{{"$first", "$nfts.tier"}}},
					{"tierName", bson.D{{"$first", "$tierInfo.name"}}},
					{"weeks",
						bson.D{
							{"$addToSet",
								bson.D{
									{"appliedBoostValue",
										bson.D{
											{"$gte",
												bson.A{
													"$nfts.appliedBoostValue",
													1,
												},
											},
										},
									},
									{"Week", "$nfts.Week"},
									{"SeasonName", "$nfts.SeasonName"},
									{"Season", "$nfts.Season"},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"nfts._id", "$nftId"},
					{"nfts.playerId", "$playerId"},
					{"nfts.tierId", "$tierId"},
					{"nfts.tierName", "$tierName"},
					{"nfts.weeks", "$weeks"},
					{"nfts.team", "$team"},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"walletAddress", 1},
					{"nfts", 1},
					{"userId", 1},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$walletAddress"},
					{"userId", bson.D{{"$first", "$userId"}}},
					{"nfts", bson.D{{"$push", "$nfts"}}},
				},
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	defer cursor.Close(ctx)
	if err != nil {
		fmt.Println("Error performing aggregation:", err)
	}

	var result []types.Result
	if err = cursor.All(ctx, &result); err != nil {
		fmt.Println("Error performing aggregation:", err)
	}

	client.Disconnect(ctx)
	return result
}

func (m Matches) SavePlayerStats() {
	for _, player := range m.Stats {
		for _, stat := range player.Stats {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			clientOptions := options.Client().ApplyURI(DB_CONNECTION)

			client, err := mongo.Connect(ctx, clientOptions)
			if err != nil {
				fmt.Println("Error connecting to MongoDB:", err)
			}

			err = client.Ping(ctx, nil)
			if err != nil {
				fmt.Println("Error checking MongoDB connection:", err)
			}

			collection := client.Database(DB_NAME).Collection(DB_COLLECTION_PLAYERS)

			var result types.PlayerStats
			err = collection.FindOne(ctx, bson.M{"PlayerID": stat.PlayerID, "Week": stat.Week, "Season": stat.Season, "SeasonType": stat.SeasonType}).Decode(&result)
			if err == nil {
				collection.UpdateOne(ctx, bson.M{"PlayerID": stat.PlayerID, "Week": stat.Week, "Season": stat.Season, "SeasonType": stat.SeasonType}, bson.M{"$set": stat})
			} else {
				collection.InsertOne(ctx, stat)
			}

			client.Disconnect(ctx)
		}
	}
}
