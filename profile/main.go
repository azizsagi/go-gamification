package profile

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	API_KEY string = "f5f5e61bc9d84f1c9f7c5ffc8f66eb4c"
	// DB_CONNECTION         string = "mongodb+srv://personal-corner-db:MaH0dp4YgcGknJp7@personalcorner.v0pkq.mongodb.net/pcc-Dev?retryWrites=true&w=majority"
	DB_CONNECTION         string = "mongodb+srv://personal-corner-db:MaH0dp4YgcGknJp7@personalcorner.v0pkq.mongodb.net/pcc-Prod?retryWrites=true&w=majority"
	DB_NAME               string = "pcc-Prod"
	DB_COLLECTION_NFT     string = "allStatsForNFT"
	DB_COLLECTION_USER    string = "users"
	DB_COLLECTION_WEEK    string = "weeks"
	DB_COLLECTION_HISTORY string = "pcchistories"
)

type Stats struct {
	Week             int     `bson:"Week"`
	PositionCategory string  `bson:"PositionCategory"`
	Position         string  `bson:"Position"`
	Season           string  `bson:"Season"`
	UserID           string  `bson:"UserID"`
	Points           float64 `bson:"Points"`
	Reward           int     `bson:"Reward"`
	SeasonName       string  `bson:"SeasonName"`
}

type SaveUser struct {
	Id             string    `bson:"_id"`
	Pcc            int       `bson:"Pcc"`
	Xp             int       `bson:"Xp"`
	Level          int       `bson:"level"`
	TotalEarnedPcc int       `bson:"totalEarnedPcc"`
	CardUnlocked   int       `bson:"cardUnlocked"`
	Week           int       `bson:"Week"`
	UpdatedAt      time.Time `bson:"updatedAt"`
}

type User struct {
	Id             string    `bson:"_id"`
	Type           string    `bson:"type"`
	Status         string    `bson:"status"`
	Pcc            int       `bson:"Pcc"`
	Xp             int       `bson:"Xp"`
	Level          int       `bson:"level"`
	TotalEarnedPcc int       `bson:"totalEarnedPcc"`
	CardUnlocked   int       `bson:"cardUnlocked"`
	Week           int       `bson:"Week"`
	UpdatedAt      time.Time `bson:"updatedAt"`
	AllStats       []Stats
}

type AllUsers struct {
	Users          []User
	CurrentWeek    int
	Season         string
	SeasonName     string
	firstMatchDate time.Time
	lastMatchDate  time.Time
	UsersToStore   []SaveUser
}

func NewAllUsers() *AllUsers {
	a := AllUsers{}

	a.GetCurrentWeek()
	// FOR TEST
	// isoDateStr := "2023-09-12T04:10:00.000+00:00"
	// targetDate, _ := time.Parse(time.RFC3339Nano, isoDateStr)
	//
	timeDifference := a.lastMatchDate.Sub(time.Now())
	// timeDifference := a.lastMatchDate.Sub(targetDate)
	if timeDifference < 8*time.Minute && a.CurrentWeek != 0 {
		fmt.Println("less than 8 mins")
		a.GetUsers()
		a.GetAllXP()
		a.Calculate()
		// a.Print()
		a.PrintAllUsersToSave()
		a.SaveAllUsers()
		a.SaveHistories()
	}
	return &a
}

func (a AllUsers) SaveHistories() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error Disconnecting the db")
		}
	}()
	for _, user := range a.UsersToStore {
		collection := client.Database(DB_NAME).Collection(DB_COLLECTION_HISTORY)
		objectID, err := primitive.ObjectIDFromHex(user.Id)
		if err != nil {
			fmt.Println("OBJECT ID unknown")
		}
		insert := bson.M{
			"userId":          objectID,
			"Pcc":             user.Pcc,
			"Season":          a.Season,
			"SeasonName":      a.SeasonName,
			"Week":            user.Week,
			"creditType":      "LineUp",
			"transactionType": "Credit",
			"createdAt":       user.UpdatedAt,
		}
		_, err = collection.InsertOne(ctx, insert)
		if err != nil {
			fmt.Println("[PROFILE] Unable to insert History", err)
		}
	}
}

func (a AllUsers) SaveAllUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error Disconnecting the db")
		}
	}()
	for _, user := range a.UsersToStore {
		collection := client.Database(DB_NAME).Collection(DB_COLLECTION_USER)
		objectID, err := primitive.ObjectIDFromHex(user.Id)
		if err != nil {
			fmt.Println("OBJECT ID unknown")
		}
		filter := bson.M{"_id": objectID}
		update := bson.M{"$set": bson.M{
			"Pcc":            user.Pcc,
			"Xp":             user.Xp,
			"level":          user.Level,
			"totalEarnedPcc": user.TotalEarnedPcc,
			"cardUnlocked":   user.CardUnlocked,
			"Week":           user.Week,
			"updatedAt":      user.UpdatedAt,
		}}
		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println("[PROFILE] Unable to update", err)
		}
	}
}

func (a AllUsers) PrintAllUsersToSave() {
	for _, user := range a.UsersToStore {
		fmt.Println("[USER} ", user)
	}
}

func (a *AllUsers) Calculate() {
	XPs := [...]int64{0, 2000, 2600, 3380, 4394, 5712, 7426, 9654, 12550, 16315, 21209, 27572, 35843, 46596, 60575, 78748, 102372, 133083, 173008, 224911, 292384, 380099, 494129, 642368, 835078, 1085602, 1411282, 1834667, 2385067, 3100587, 4030763, 5239991, 6811989, 8855585, 11512261, 14965939, 19455721, 25292437, 32880168, 42744219, 55567484, 72237730, 93909049, 122081763, 158706292, 206318180, 268213633, 348677723, 453281040, 589265353}
	Levels := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	PositionUnLock := [...]int{10, 30, 40, 50}

	// get the sum of all the xps
	for _, user := range a.Users {
		if user.Week == a.CurrentWeek {
			continue
		}
		totalXP := 0.0
		for _, nft := range user.AllStats {
			totalXP += nft.Points
		}
		if totalXP <= 0.0 {
			continue
		}
		// do the further calculations
		TotalPCCFromNFTs := 0
		for _, nft := range user.AllStats {
			TotalPCCFromNFTs += nft.Reward
		}
		myLevel := 0
		for index, xp := range XPs {
			if xp < int64(totalXP+float64(user.Xp)) {
				// fmt.Println("TOTAL", totalXP+float64(user.Xp))
				myLevel = Levels[index]
			}
		}
		myRewards := 0
		unLock := 0
		if myLevel > user.Level {
			for _, posUnlock := range PositionUnLock {
				if myLevel > posUnlock && user.Level < posUnlock {
					unLock++
				}
			}
			lvl := 1
			if user.Level != 0 {
				lvl = user.Level
			}
			for i := lvl; i <= myLevel; i++ {
				if i < 10 && i > 1 {
					myRewards += 50
				} else if i < 20 && i > 10 {
					myRewards += 100
				} else if i < 30 && i >= 20 {
					myRewards += 200
				} else if i < 40 && i > 30 {
					myRewards += 400
				} else if i < 50 && i > 40 {
					myRewards += 800
				}
				if i == 50 {
					myRewards += 1000
				}
			}
			// fmt.Println("POINTS", totalXP, "Level", myLevel, "Reward", myRewards+TotalPCCFromNFTs, "NewPositions", unLock)
			var myUser SaveUser
			myUser.Id = user.Id
			myUser.Level = myLevel
			myUser.Xp = int(totalXP)
			myUser.Pcc = myRewards + TotalPCCFromNFTs + user.Pcc
			myUser.TotalEarnedPcc = myRewards + TotalPCCFromNFTs + user.TotalEarnedPcc
			myUser.Week = a.CurrentWeek
			myUser.CardUnlocked = user.CardUnlocked + unLock
			myUser.UpdatedAt = time.Now().UTC()
			a.UsersToStore = append(a.UsersToStore, myUser)
		}
	}
}

func (a *AllUsers) GetCurrentWeek() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_WEEK)

	currentDate := time.Now().UTC()
	isoDateStr := currentDate.Format("2006-01-02T15:04:05.999Z")
	// for test
	// isoDateStr = "2023-09-12T04:10:00.000+00:00"
	// for test
	targetDate, err := time.Parse(time.RFC3339Nano, isoDateStr)
	if err != nil {
		fmt.Println("Error converting date:", err)
	}
	pipeline := bson.M{
		"firstMatchDate": bson.M{
			"$lt": targetDate,
		},
		"lastMatchDate": bson.M{
			"$gt": targetDate,
		},
	}

	var results bson.M
	if err = collection.FindOne(ctx, pipeline).Decode(&results); err == nil && len(results) > 0 {
		CurrentWeek, _ := results["Week"].(int32)
		a.CurrentWeek = int(CurrentWeek)

		a.Season, _ = results["Season"].(string)
		a.SeasonName, _ = results["SeasonName"].(string)

		dt, _ := results["firstMatchDate"].(primitive.DateTime)
		a.firstMatchDate = time.Unix(int64(dt)/1000, int64(dt)%1000*int64(time.Millisecond))

		dt, _ = results["lastMatchDate"].(primitive.DateTime)
		a.lastMatchDate = time.Unix(int64(dt)/1000, int64(dt)%1000*int64(time.Millisecond))
		fmt.Println("[PROFILE WEEK] successfully found", a.CurrentWeek)
	} else {
		fmt.Println("EROR", err)
	}
	client.Disconnect(ctx)
}

func (a *AllUsers) GetAllXP() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error Disconnecting DB", err)
		}
	}()

	allUsers := a.Users
	for id, user := range allUsers {
		collection := client.Database(DB_NAME).Collection(DB_COLLECTION_NFT)

		cursor, err := collection.Find(ctx, bson.M{"UserID": user.Id, "Week": a.CurrentWeek})
		if err != nil {
			continue
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var elem Stats
			err := cursor.Decode(&elem)
			if err != nil {
				fmt.Println(err)
			}
			a.Users[id].AllStats = append(a.Users[id].AllStats, elem)
		}
	}
}

func (a *AllUsers) Print() {
	for _, user := range a.Users {
		fmt.Println("[PROFILE] ", user)
	}
}

func (a *AllUsers) GetUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(DB_CONNECTION)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	collection := client.Database(DB_NAME).Collection(DB_COLLECTION_USER)

	cursor, err := collection.Find(context.TODO(), bson.D{})

	defer cursor.Close(ctx)

	if err != nil {
		fmt.Println("Error Finding MongoDB connection:", err)
	}

	for cursor.Next(ctx) {
		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		a.Users = append(a.Users, elem)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println(err)
	}
	client.Disconnect(ctx)

}
