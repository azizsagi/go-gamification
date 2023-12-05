package seasons

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hassan-algo/pc-gamification/extras"
	"github.com/hassan-algo/pc-gamification/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	API_KEY                   string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	DB_CONNECTION             string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	DB_NAME                   string = "pcc-Prod"
	DB_COLLECTION_SEASON      string = "playerstatebyseasonsInWay"
	DB_COLLECTION_ALL_PLAYERS string = "collections"
)

type PlayerList struct {
	PlayerID int    `bson:"playerId" json:"playerId"`
	Team     string `bson:"Team"`
}
type PlayerSeasonStat struct {
	PlayerID []int
	Team     string
}
type SeasonStats struct {
	Teams       []PlayerSeasonStat
	PlayerStats []types.PlayerStats
	Season      string
}

func (s SeasonStats) ExistTeam(team string) int {
	for id, myTeam := range s.Teams {
		if myTeam.Team == team {
			return id
		}
	}
	return -1
}

func NewSeaonStats() *SeasonStats {
	s := SeasonStats{}
	s.GetCurrentSeason()
	players := GetAllPlayers()
	for _, player := range players {
		teamExists := s.ExistTeam(player.Team)
		if teamExists == -1 {
			s.Teams = append(s.Teams, PlayerSeasonStat{
				Team: player.Team,
			})
			teamExists = len(s.Teams) - 1
		}
		s.Teams[teamExists].PlayerID = append(s.Teams[teamExists].PlayerID, player.PlayerID)
	}
	// s.PrintPlayerStats()
	return &s
}

func (m *SeasonStats) ExistPlayer(playerID int) bool {
	for _, player := range m.Teams {
		if extras.ContainsInt(player.PlayerID, playerID) {
			return true
		}
	}
	return false
}

func (m *SeasonStats) SaveSeasonRecords() {
	for _, player := range m.PlayerStats {
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

		collection := client.Database(DB_NAME).Collection(DB_COLLECTION_SEASON)

		var result bson.M
		if err := collection.FindOne(ctx, bson.M{"PlayerID": player.PlayerID, "Season": player.Season, "SeasonType": player.SeasonType}).Decode(&result); err == nil && len(result) > 0 {
			collection.UpdateOne(ctx, bson.M{"PlayerID": player.PlayerID, "Season": player.Season, "SeasonType": player.SeasonType}, bson.M{"$set": player})
		} else {
			collection.InsertOne(ctx, player)
		}

		client.Disconnect(ctx)
	}
}

func (m *SeasonStats) GetCurrentSeason() {
	URL := "https://api.sportsdata.io/v3/nfl/scores/json/CurrentSeason?key=" + API_KEY
	res, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	season := string(body)
	m.Season = season
}

func (m *SeasonStats) GetPlayerStats() {
	for _, team := range m.Teams {
		URL := "https://api.sportsdata.io/v3/nfl/stats/json/PlayerSeasonStatsByTeam/" + m.Season + "/" + team.Team + "?key=" + API_KEY
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
			if m.ExistPlayer(player.PlayerID) {
				m.PlayerStats = append(m.PlayerStats, player)
			}
		}
	}
}

func (m SeasonStats) PrintPlayerStats() {
	for _, player := range m.PlayerStats {
		fmt.Println(player.PlayerID)
	}
}

func GetAllPlayers() []PlayerList {
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

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_ALL_PLAYERS)

	pipeline := bson.A{
		bson.D{{"$group", bson.D{{"_id", "$playerId"}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"as", "player"},
					{"from", "players"},
					{"let", bson.D{{"playerId", "$_id"}}},
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
			{"$project",
				bson.D{
					{"playerId", "$_id"},
					{"Team", "$player.Team"},
				},
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	defer cursor.Close(ctx)
	if err != nil {
		fmt.Println("Error performing aggregation:", err)
	}

	var result []PlayerList
	if err = cursor.All(ctx, &result); err != nil {
		fmt.Println("Error performing aggregation:", err)
	}

	client.Disconnect(ctx)
	return result
}
