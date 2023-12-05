package types

type ScoresByWeek struct {
	GameKey    string
	SeasonType int
	Season     int
	Week       int
	Date       string
	AwayTeam   string
	HomeTeam   string
}

type GenerateAblePlayerStats struct {
	PlayerID                          int     `json:"PlayerID" bson:"PlayerID"`
	Week                              int     `json:"Week" bson:"Week"`
	PassingAttempts                   float64 `json:"PassingAttempts" bson:"PassingAttempts"`
	PassingCompletions                float64 `json:"PassingCompletions" bson:"PassingCompletions"`
	PassingYards                      float64 `json:"PassingYards" bson:"PassingYards"`
	PassingTouchdowns                 float64 `json:"PassingTouchdowns" bson:"PassingTouchdowns"`
	PassingCompletionPercentage       float64 `json:"PassingCompletionPercentage" bson:"PassingCompletionPercentage"`
	PassingYardsPerAttempt            float64 `json:"PassingYardsPerAttempt" bson:"PassingYardsPerAttempt"`
	PassingYardsPerCompletion         float64 `json:"PassingYardsPerCompletion" bson:"PassingYardsPerCompletion"`
	PassingInterceptions              float64 `json:"PassingInterceptions" bson:"PassingInterceptions"`
	PassingRating                     float64 `json:"PassingRating" bson:"PassingRating"`
	PassingLong                       float64 `json:"PassingLong" bson:"PassingLong"`
	PassingSacks                      float64 `json:"PassingSacks" bson:"PassingSacks"`
	PassingSackYards                  float64 `json:"PassingSackYards" bson:"PassingSackYards"`
	RushingAttempts                   float64 `json:"RushingAttempts" bson:"RushingAttempts"`
	RushingYards                      float64 `json:"RushingYards" bson:"RushingYards"`
	RushingYardsPerAttempt            float64 `json:"RushingYardsPerAttempt" bson:"RushingYardsPerAttempt"`
	RushingTouchdowns                 float64 `json:"RushingTouchdowns" bson:"RushingTouchdowns"`
	RushingLong                       float64 `json:"RushingLong" bson:"RushingLong"`
	ReceivingTargets                  float64 `json:"ReceivingTargets" bson:"ReceivingTargets"`
	Receptions                        float64 `json:"Receptions" bson:"Receptions"`
	ReceivingYards                    float64 `json:"ReceivingYards" bson:"ReceivingYards"`
	ReceivingYardsPerReception        float64 `json:"ReceivingYardsPerReception" bson:"ReceivingYardsPerReception"`
	ReceivingTouchdowns               float64 `json:"ReceivingTouchdowns" bson:"ReceivingTouchdowns"`
	ReceivingLong                     float64 `json:"ReceivingLong" bson:"ReceivingLong"`
	Fumbles                           float64 `json:"Fumbles" bson:"Fumbles"`
	FumblesLost                       float64 `json:"FumblesLost" bson:"FumblesLost"`
	PuntReturns                       float64 `json:"PuntReturns" bson:"PuntReturns"`
	PuntReturnYards                   float64 `json:"PuntReturnYards" bson:"PuntReturnYards"`
	PuntReturnYardsPerAttempt         float64 `json:"PuntReturnYardsPerAttempt" bson:"PuntReturnYardsPerAttempt"`
	PuntReturnTouchdowns              float64 `json:"PuntReturnTouchdowns" bson:"PuntReturnTouchdowns"`
	PuntReturnLong                    float64 `json:"PuntReturnLong" bson:"PuntReturnLong"`
	KickReturns                       float64 `json:"KickReturns" bson:"KickReturns"`
	KickReturnYards                   float64 `json:"KickReturnYards" bson:"KickReturnYards"`
	KickReturnYardsPerAttempt         float64 `json:"KickReturnYardsPerAttempt" bson:"KickReturnYardsPerAttempt"`
	KickReturnTouchdowns              float64 `json:"KickReturnTouchdowns" bson:"KickReturnTouchdowns"`
	KickReturnLong                    float64 `json:"KickReturnLong" bson:"KickReturnLong"`
	SoloTackles                       float64 `json:"SoloTackles" bson:"SoloTackles"`
	AssistedTackles                   float64 `json:"AssistedTackles" bson:"AssistedTackles"`
	TacklesForLoss                    float64 `json:"TacklesForLoss" bson:"TacklesForLoss"`
	Sacks                             float64 `json:"Sacks" bson:"Sacks"`
	SackYards                         float64 `json:"SackYards" bson:"SackYards"`
	QuarterbackHits                   float64 `json:"QuarterbackHits" bson:"QuarterbackHits"`
	PassesDefended                    float64 `json:"PassesDefended" bson:"PassesDefended"`
	FumblesForced                     float64 `json:"FumblesForced" bson:"FumblesForced"`
	FumblesRecovered                  float64 `json:"FumblesRecovered" bson:"FumblesRecovered"`
	FumbleReturnYards                 float64 `json:"FumbleReturnYards" bson:"FumbleReturnYards"`
	FumbleReturnTouchdowns            float64 `json:"FumbleReturnTouchdowns" bson:"FumbleReturnTouchdowns"`
	Interceptions                     float64 `json:"Interceptions" bson:"Interceptions"`
	InterceptionReturnYards           float64 `json:"InterceptionReturnYards" bson:"InterceptionReturnYards"`
	InterceptionReturnTouchdowns      float64 `json:"InterceptionReturnTouchdowns" bson:"InterceptionReturnTouchdowns"`
	BlockedKicks                      float64 `json:"BlockedKicks" bson:"BlockedKicks"`
	SpecialTeamsSoloTackles           float64 `json:"SpecialTeamsSoloTackles" bson:"SpecialTeamsSoloTackles"`
	SpecialTeamsAssistedTackles       float64 `json:"SpecialTeamsAssistedTackles" bson:"SpecialTeamsAssistedTackles"`
	MiscSoloTackles                   float64 `json:"MiscSoloTackles" bson:"MiscSoloTackles"`
	MiscAssistedTackles               float64 `json:"MiscAssistedTackles" bson:"MiscAssistedTackles"`
	Punts                             float64 `json:"Punts" bson:"Punts"`
	PuntYards                         float64 `json:"PuntYards" bson:"PuntYards"`
	PuntAverage                       float64 `json:"PuntAverage" bson:"PuntAverage"`
	FieldGoalsAttempted               float64 `json:"FieldGoalsAttempted" bson:"FieldGoalsAttempted"`
	FieldGoalsMade                    float64 `json:"FieldGoalsMade" bson:"FieldGoalsMade"`
	FieldGoalsLongestMade             float64 `json:"FieldGoalsLongestMade" bson:"FieldGoalsLongestMade"`
	ExtraPointsMade                   float64 `json:"ExtraPointsMade" bson:"ExtraPointsMade"`
	TwoPointConversionPasses          float64 `json:"TwoPointConversionPasses" bson:"TwoPointConversionPasses"`
	TwoPointConversionRuns            float64 `json:"TwoPointConversionRuns" bson:"TwoPointConversionRuns"`
	TwoPointConversionReceptions      float64 `json:"TwoPointConversionReceptions" bson:"TwoPointConversionReceptions"`
	FantasyPoints                     float64 `json:"FantasyPoints" bson:"FantasyPoints"`
	FantasyPointsPPR                  float64 `json:"FantasyPointsPPR" bson:"FantasyPointsPPR"`
	ReceptionPercentage               float64 `json:"ReceptionPercentage" bson:"ReceptionPercentage"`
	ReceivingYardsPerTarget           float64 `json:"ReceivingYardsPerTarget" bson:"ReceivingYardsPerTarget"`
	Tackles                           float64 `json:"Tackles" bson:"Tackles"`
	OffensiveTouchdowns               float64 `json:"OffensiveTouchdowns" bson:"OffensiveTouchdowns"`
	DefensiveTouchdowns               float64 `json:"DefensiveTouchdowns" bson:"DefensiveTouchdowns"`
	SpecialTeamsTouchdowns            float64 `json:"SpecialTeamsTouchdowns" bson:"SpecialTeamsTouchdowns"`
	Touchdowns                        float64 `json:"Touchdowns" bson:"Touchdowns"`
	FieldGoalPercentage               float64 `json:"FieldGoalPercentage" bson:"FieldGoalPercentage"`
	FumblesOwnRecoveries              float64 `json:"FumblesOwnRecoveries" bson:"FumblesOwnRecoveries"`
	FumblesOutOfBounds                float64 `json:"FumblesOutOfBounds" bson:"FumblesOutOfBounds"`
	KickReturnFairCatches             float64 `json:"KickReturnFairCatches" bson:"KickReturnFairCatches"`
	PuntReturnFairCatches             float64 `json:"PuntReturnFairCatches" bson:"PuntReturnFairCatches"`
	PuntTouchbacks                    float64 `json:"PuntTouchbacks" bson:"PuntTouchbacks"`
	PuntInside20                      float64 `json:"PuntInside20" bson:"PuntInside20"`
	PuntNetAverage                    float64 `json:"PuntNetAverage" bson:"PuntNetAverage"`
	ExtraPointsAttempted              float64 `json:"ExtraPointsAttempted" bson:"ExtraPointsAttempted"`
	BlockedKickReturnTouchdowns       float64 `json:"BlockedKickReturnTouchdowns" bson:"BlockedKickReturnTouchdowns"`
	FieldGoalReturnTouchdowns         float64 `json:"FieldGoalReturnTouchdowns" bson:"FieldGoalReturnTouchdowns"`
	Safeties                          float64 `json:"Safeties" bson:"Safeties"`
	FieldGoalsHadBlocked              float64 `json:"FieldGoalsHadBlocked" bson:"FieldGoalsHadBlocked"`
	PuntsHadBlocked                   float64 `json:"PuntsHadBlocked" bson:"PuntsHadBlocked"`
	ExtraPointsHadBlocked             float64 `json:"ExtraPointsHadBlocked" bson:"ExtraPointsHadBlocked"`
	PuntLong                          float64 `json:"PuntLong" bson:"PuntLong"`
	BlockedKickReturnYards            float64 `json:"BlockedKickReturnYards" bson:"BlockedKickReturnYards"`
	FieldGoalReturnYards              float64 `json:"FieldGoalReturnYards" bson:"FieldGoalReturnYards"`
	PuntNetYards                      float64 `json:"PuntNetYards" bson:"PuntNetYards"`
	SpecialTeamsFumblesForced         float64 `json:"SpecialTeamsFumblesForced" bson:"SpecialTeamsFumblesForced"`
	SpecialTeamsFumblesRecovered      float64 `json:"SpecialTeamsFumblesRecovered" bson:"SpecialTeamsFumblesRecovered"`
	MiscFumblesForced                 float64 `json:"MiscFumblesForced" bson:"MiscFumblesForced"`
	MiscFumblesRecovered              float64 `json:"MiscFumblesRecovered" bson:"MiscFumblesRecovered"`
	TwoPointConversionReturns         float64 `json:"TwoPointConversionReturns" bson:"TwoPointConversionReturns"`
	FieldGoalsMade0to19               float64 `json:"FieldGoalsMade0to19" bson:"FieldGoalsMade0to19"`
	FieldGoalsMade20to29              float64 `json:"FieldGoalsMade20to29" bson:"FieldGoalsMade20to29"`
	FieldGoalsMade30to39              float64 `json:"FieldGoalsMade30to39" bson:"FieldGoalsMade30to39"`
	FieldGoalsMade40to49              float64 `json:"FieldGoalsMade40to49" bson:"FieldGoalsMade40to49"`
	FieldGoalsMade50Plus              float64 `json:"FieldGoalsMade50Plus" bson:"FieldGoalsMade50Plus"`
	FantasyPointsDraftKings           float64 `json:"FantasyPointsDraftKings" bson:"FantasyPointsDraftKings"`
	FantasyPointsYahoo                float64 `json:"FantasyPointsYahoo" bson:"FantasyPointsYahoo"`
	OffensiveFumbleRecoveryTouchdowns float64 `json:"OffensiveFumbleRecoveryTouchdowns" bson:"OffensiveFumbleRecoveryTouchdowns"`
}

type PlayerStats struct {
	UserID                            string  `json:"UserID" bson:"UserID"`
	Points                            float64 `json:"Points"`
	AppliedBoost                      bool    `json:"appliedBoostValue" bson:"appliedBoostValue"`
	GameKey                           string  `json:"GameKey" bson:"GameKey"`
	SeasonType                        int     `json:"SeasonType" bson:"SeasonType"`
	Season                            int     `json:"Season" bson:"Season"`
	Week                              int     `json:"Week" bson:"Week"`
	PlayerID                          int     `json:"PlayerID" bson:"PlayerID"`
	Played                            int     `json:"Played" bson:"Played"`
	Started                           int     `json:"Started" bson:"Started"`
	Name                              string  `json:"Name" bson:"Name"`
	Number                            int     `json:"Number" bson:"Number"`
	Position                          string  `json:"Position" bson:"Position"`
	PositionCategory                  string  `json:"PositionCategory" bson:"PositionCategory"`
	Avtivated                         int     `json:"Avtivated" bson:"Avtivated"`
	PassingAttempts                   float64 `json:"PassingAttempts" bson:"PassingAttempts"`
	PassingCompletions                float64 `json:"PassingCompletions" bson:"PassingCompletions"`
	PassingYards                      float64 `json:"PassingYards" bson:"PassingYards"`
	PassingTouchdowns                 float64 `json:"PassingTouchdowns" bson:"PassingTouchdowns"`
	PassingCompletionPercentage       float64 `json:"PassingCompletionPercentage" bson:"PassingCompletionPercentage"`
	PassingYardsPerAttempt            float64 `json:"PassingYardsPerAttempt" bson:"PassingYardsPerAttempt"`
	PassingYardsPerCompletion         float64 `json:"PassingYardsPerCompletion" bson:"PassingYardsPerCompletion"`
	PassingInterceptions              float64 `json:"PassingInterceptions" bson:"PassingInterceptions"`
	PassingRating                     float64 `json:"PassingRating" bson:"PassingRating"`
	PassingLong                       float64 `json:"PassingLong" bson:"PassingLong"`
	PassingSacks                      float64 `json:"PassingSacks" bson:"PassingSacks"`
	PassingSackYards                  float64 `json:"PassingSackYards" bson:"PassingSackYards"`
	RushingAttempts                   float64 `json:"RushingAttempts" bson:"RushingAttempts"`
	RushingYards                      float64 `json:"RushingYards" bson:"RushingYards"`
	RushingYardsPerAttempt            float64 `json:"RushingYardsPerAttempt" bson:"RushingYardsPerAttempt"`
	RushingTouchdowns                 float64 `json:"RushingTouchdowns" bson:"RushingTouchdowns"`
	RushingLong                       float64 `json:"RushingLong" bson:"RushingLong"`
	ReceivingTargets                  float64 `json:"ReceivingTargets" bson:"ReceivingTargets"`
	Receptions                        float64 `json:"Receptions" bson:"Receptions"`
	ReceivingYards                    float64 `json:"ReceivingYards" bson:"ReceivingYards"`
	ReceivingYardsPerReception        float64 `json:"ReceivingYardsPerReception" bson:"ReceivingYardsPerReception"`
	ReceivingTouchdowns               float64 `json:"ReceivingTouchdowns" bson:"ReceivingTouchdowns"`
	ReceivingLong                     float64 `json:"ReceivingLong" bson:"ReceivingLong"`
	Fumbles                           float64 `json:"Fumbles" bson:"Fumbles"`
	FumblesLost                       float64 `json:"FumblesLost" bson:"FumblesLost"`
	PuntReturns                       float64 `json:"PuntReturns" bson:"PuntReturns"`
	PuntReturnYards                   float64 `json:"PuntReturnYards" bson:"PuntReturnYards"`
	PuntReturnYardsPerAttempt         float64 `json:"PuntReturnYardsPerAttempt" bson:"PuntReturnYardsPerAttempt"`
	PuntReturnTouchdowns              float64 `json:"PuntReturnTouchdowns" bson:"PuntReturnTouchdowns"`
	PuntReturnLong                    float64 `json:"PuntReturnLong" bson:"PuntReturnLong"`
	KickReturns                       float64 `json:"KickReturns" bson:"KickReturns"`
	KickReturnYards                   float64 `json:"KickReturnYards" bson:"KickReturnYards"`
	KickReturnYardsPerAttempt         float64 `json:"KickReturnYardsPerAttempt" bson:"KickReturnYardsPerAttempt"`
	KickReturnTouchdowns              float64 `json:"KickReturnTouchdowns" bson:"KickReturnTouchdowns"`
	KickReturnLong                    float64 `json:"KickReturnLong" bson:"KickReturnLong"`
	SoloTackles                       float64 `json:"SoloTackles" bson:"SoloTackles"`
	AssistedTackles                   float64 `json:"AssistedTackles" bson:"AssistedTackles"`
	TacklesForLoss                    float64 `json:"TacklesForLoss" bson:"TacklesForLoss"`
	Sacks                             float64 `json:"Sacks" bson:"Sacks"`
	SackYards                         float64 `json:"SackYards" bson:"SackYards"`
	QuarterbackHits                   float64 `json:"QuarterbackHits" bson:"QuarterbackHits"`
	PassesDefended                    float64 `json:"PassesDefended" bson:"PassesDefended"`
	FumblesForced                     float64 `json:"FumblesForced" bson:"FumblesForced"`
	FumblesRecovered                  float64 `json:"FumblesRecovered" bson:"FumblesRecovered"`
	FumbleReturnYards                 float64 `json:"FumbleReturnYards" bson:"FumbleReturnYards"`
	FumbleReturnTouchdowns            float64 `json:"FumbleReturnTouchdowns" bson:"FumbleReturnTouchdowns"`
	Interceptions                     float64 `json:"Interceptions" bson:"Interceptions"`
	InterceptionReturnYards           float64 `json:"InterceptionReturnYards" bson:"InterceptionReturnYards"`
	InterceptionReturnTouchdowns      float64 `json:"InterceptionReturnTouchdowns" bson:"InterceptionReturnTouchdowns"`
	BlockedKicks                      float64 `json:"BlockedKicks" bson:"BlockedKicks"`
	SpecialTeamsSoloTackles           float64 `json:"SpecialTeamsSoloTackles" bson:"SpecialTeamsSoloTackles"`
	SpecialTeamsAssistedTackles       float64 `json:"SpecialTeamsAssistedTackles" bson:"SpecialTeamsAssistedTackles"`
	MiscSoloTackles                   float64 `json:"MiscSoloTackles" bson:"MiscSoloTackles"`
	MiscAssistedTackles               float64 `json:"MiscAssistedTackles" bson:"MiscAssistedTackles"`
	Punts                             float64 `json:"Punts" bson:"Punts"`
	PuntYards                         float64 `json:"PuntYards" bson:"PuntYards"`
	PuntAverage                       float64 `json:"PuntAverage" bson:"PuntAverage"`
	FieldGoalsAttempted               float64 `json:"FieldGoalsAttempted" bson:"FieldGoalsAttempted"`
	FieldGoalsMade                    float64 `json:"FieldGoalsMade" bson:"FieldGoalsMade"`
	FieldGoalsLongestMade             float64 `json:"FieldGoalsLongestMade" bson:"FieldGoalsLongestMade"`
	ExtraPointsMade                   float64 `json:"ExtraPointsMade" bson:"ExtraPointsMade"`
	TwoPointConversionPasses          float64 `json:"TwoPointConversionPasses" bson:"TwoPointConversionPasses"`
	TwoPointConversionRuns            float64 `json:"TwoPointConversionRuns" bson:"TwoPointConversionRuns"`
	TwoPointConversionReceptions      float64 `json:"TwoPointConversionReceptions" bson:"TwoPointConversionReceptions"`
	FantasyPoints                     float64 `json:"FantasyPoints" bson:"FantasyPoints"`
	FantasyPointsPPR                  float64 `json:"FantasyPointsPPR" bson:"FantasyPointsPPR"`
	ReceptionPercentage               float64 `json:"ReceptionPercentage" bson:"ReceptionPercentage"`
	ReceivingYardsPerTarget           float64 `json:"ReceivingYardsPerTarget" bson:"ReceivingYardsPerTarget"`
	Tackles                           float64 `json:"Tackles" bson:"Tackles"`
	OffensiveTouchdowns               float64 `json:"OffensiveTouchdowns" bson:"OffensiveTouchdowns"`
	DefensiveTouchdowns               float64 `json:"DefensiveTouchdowns" bson:"DefensiveTouchdowns"`
	SpecialTeamsTouchdowns            float64 `json:"SpecialTeamsTouchdowns" bson:"SpecialTeamsTouchdowns"`
	Touchdowns                        float64 `json:"Touchdowns" bson:"Touchdowns"`
	FantasyPosition                   string  `json:"FantasyPosition" bson:"FantasyPosition"`
	FieldGoalPercentage               float64 `json:"FieldGoalPercentage" bson:"FieldGoalPercentage"`
	PlayerGameID                      int     `json:"PlayerGameID" bson:"PlayerGameID"`
	FumblesOwnRecoveries              float64 `json:"FumblesOwnRecoveries" bson:"FumblesOwnRecoveries"`
	FumblesOutOfBounds                float64 `json:"FumblesOutOfBounds" bson:"FumblesOutOfBounds"`
	KickReturnFairCatches             float64 `json:"KickReturnFairCatches" bson:"KickReturnFairCatches"`
	PuntReturnFairCatches             float64 `json:"PuntReturnFairCatches" bson:"PuntReturnFairCatches"`
	PuntTouchbacks                    float64 `json:"PuntTouchbacks" bson:"PuntTouchbacks"`
	PuntInside20                      float64 `json:"PuntInside20" bson:"PuntInside20"`
	PuntNetAverage                    float64 `json:"PuntNetAverage" bson:"PuntNetAverage"`
	ExtraPointsAttempted              float64 `json:"ExtraPointsAttempted" bson:"ExtraPointsAttempted"`
	BlockedKickReturnTouchdowns       float64 `json:"BlockedKickReturnTouchdowns" bson:"BlockedKickReturnTouchdowns"`
	FieldGoalReturnTouchdowns         float64 `json:"FieldGoalReturnTouchdowns" bson:"FieldGoalReturnTouchdowns"`
	Safeties                          float64 `json:"Safeties" bson:"Safeties"`
	FieldGoalsHadBlocked              float64 `json:"FieldGoalsHadBlocked" bson:"FieldGoalsHadBlocked"`
	PuntsHadBlocked                   float64 `json:"PuntsHadBlocked" bson:"PuntsHadBlocked"`
	ExtraPointsHadBlocked             float64 `json:"ExtraPointsHadBlocked" bson:"ExtraPointsHadBlocked"`
	PuntLong                          float64 `json:"PuntLong" bson:"PuntLong"`
	BlockedKickReturnYards            float64 `json:"BlockedKickReturnYards" bson:"BlockedKickReturnYards"`
	FieldGoalReturnYards              float64 `json:"FieldGoalReturnYards" bson:"FieldGoalReturnYards"`
	PuntNetYards                      float64 `json:"PuntNetYards" bson:"PuntNetYards"`
	SpecialTeamsFumblesForced         float64 `json:"SpecialTeamsFumblesForced" bson:"SpecialTeamsFumblesForced"`
	SpecialTeamsFumblesRecovered      float64 `json:"SpecialTeamsFumblesRecovered" bson:"SpecialTeamsFumblesRecovered"`
	MiscFumblesForced                 float64 `json:"MiscFumblesForced" bson:"MiscFumblesForced"`
	MiscFumblesRecovered              float64 `json:"MiscFumblesRecovered" bson:"MiscFumblesRecovered"`
	ShortName                         string  `json:"ShortName" bson:"ShortName"`
	PlayingSurface                    string  `json:"PlayingSurface" bson:"PlayingSurface"`
	IsGameOver                        bool    `json:"IsGameOver" bson:"IsGameOver"`
	SafetiesAllowed                   float64 `json:"SafetiesAllowed" bson:"SafetiesAllowed"`
	Stadium                           string  `json:"Stadium" bson:"Stadium"`
	Temperature                       float64 `json:"Temperature" bson:"Temperature"`
	Humidity                          float64 `json:"Humidity" bson:"Humidity"`
	WindSpeed                         float64 `json:"WindSpeed" bson:"WindSpeed"`
	FanDuelSalary                     float64 `json:"FanDuelSalary" bson:"FanDuelSalary"`
	DraftKingsSalary                  float64 `json:"DraftKingsSalary" bson:"DraftKingsSalary"`
	FantasyDataSalary                 float64 `json:"FantasyDataSalary" bson:"FantasyDataSalary"`
	OffensiveSnapsPlayed              int     `json:"OffensiveSnapsPlayed" bson:"OffensiveSnapsPlayed"`
	DefensiveSnapsPlayed              int     `json:"DefensiveSnapsPlayed" bson:"DefensiveSnapsPlayed"`
	SpecialTeamsSnapsPlayed           int     `json:"SpecialTeamsSnapsPlayed" bson:"SpecialTeamsSnapsPlayed"`
	OffensiveTeamSnaps                int     `json:"OffensiveTeamSnaps" bson:"OffensiveTeamSnaps"`
	DefensiveTeamSnaps                int     `json:"DefensiveTeamSnaps" bson:"DefensiveTeamSnaps"`
	SpecialTeamsTeamSnaps             int     `json:"SpecialTeamsTeamSnaps" bson:"SpecialTeamsTeamSnaps"`
	VictivSalary                      float64 `json:"VictivSalary" bson:"VictivSalary"`
	TwoPointConversionReturns         float64 `json:"TwoPointConversionReturns" bson:"TwoPointConversionReturns"`
	FieldGoalsMade0to19               float64 `json:"FieldGoalsMade0to19" bson:"FieldGoalsMade0to19"`
	FieldGoalsMade20to29              float64 `json:"FieldGoalsMade20to29" bson:"FieldGoalsMade20to29"`
	FieldGoalsMade30to39              float64 `json:"FieldGoalsMade30to39" bson:"FieldGoalsMade30to39"`
	FieldGoalsMade40to49              float64 `json:"FieldGoalsMade40to49" bson:"FieldGoalsMade40to49"`
	FieldGoalsMade50Plus              float64 `json:"FieldGoalsMade50Plus" bson:"FieldGoalsMade50Plus"`
	FantasyPointsDraftKings           float64 `json:"FantasyPointsDraftKings" bson:"FantasyPointsDraftKings"`
	FantasyPointsYahoo                float64 `json:"FantasyPointsYahoo" bson:"FantasyPointsYahoo"`
	ScoreID                           int     `json:"ScoreID" bson:"ScoreID"`
	OffensiveFumbleRecoveryTouchdowns float64 `json:"OffensiveFumbleRecoveryTouchdowns" bson:"OffensiveFumbleRecoveryTouchdowns"`
	SnapCountsConfirmed               bool    `json:"SnapCountsConfirmed" bson:"SnapCountsConfirmed"`
}

type ValidResponse struct {
	Valid bool
	Data  any
}

type Match struct {
	UserId     string
	Team       string
	Week       int
	Tier       string
	Players    []int
	Season     string
	SeasonName string
	Boost      bool
}

type Player struct {
	PlayerID int
	Stats    PlayerStats
}

type NFTStats struct {
	UserID     string
	PlayerID   int
	Tier       string
	Season     string
	SeasonName string
	Stats      []PlayerStats
}

type NFT struct {
	NFTId    string `bson:"_id"`
	PlayerId int    `bson:"playerId"`
	Team     string `bson:"team"`
	TierId   string `bson:"tierId"`
	TierName string `bson:"tierName"`
	Weeks    []Week `bson:"weeks"`
}

type Week struct {
	Week              int    `bson:"Week"`
	SeasonName        string `bson:"SeasonName"`
	Season            string `bson:"Season"`
	AppliedBoostValue bool   `bson:"appliedBoostValue" json:"appliedBoostValue"`
}

type Result struct {
	WalletAddress string `bson:"_id"`
	UserId        string `bson:"userId"`
	NFTs          []NFT  `bson:"nfts"`
}

type PositionEntry struct {
	Key   string
	Value float64
}

type PositionData []PositionEntry

var Positions = map[string]PositionData{
	"QB": {
		{"PassingYards", 5 * 0.75},
		{"PassingTouchdowns", 1500 * 0.75},
		{"RushingYards", 5 * 0.75},
		{"RushingTouchdowns", 1000 * 0.75},
		{"PassingCompletions", 35 * 0.75},
	},
	"WR": {
		{"ReceivingYards", 25 * 0.75},
		{"ReceivingTouchdowns", 2000 * 0.75},
		{"Receptions", 175 * 0.75},
	},
	"RB": {
		{"RushingYards", 35 * 0.75},
		{"RushingTouchdowns", 1500 * 0.75},
		{"ReceivingYards", 35 * 0.75},
		{"ReceivingTouchdowns", 1500 * 0.75},
	},
	"DE": {
		{"Sacks", 3000 * 0.75},
		{"SoloTackles", 500 * 0.75},
		{"TacklesForLoss", 1250 * 0.75},
		{"FumblesForced", 1500 * 0.75},
		{"DefensiveTouchdowns", 7500 * 0.75},
		{"Interceptions", 5000 * 0.75},
	},
	"LB": {
		{"Sacks", 2500 * 0.75},
		{"SoloTackles", 500 * 0.75},
		{"TacklesForLoss", 1000 * 0.75},
		{"FumblesForced", 1000 * 0.75},
		{"DefensiveTouchdowns", 5000 * 0.75},
		{"Interceptions", 2500 * 0.75},
	},
	"DB": {
		{"PassesDefended", 3000 * 0.75},
		{"SoloTackles", 400 * 0.75},
		{"FumblesForced", 1500 * 0.75},
		{"DefensiveTouchdowns", 5000 * 0.75},
		{"Interceptions", 3500 * 0.75},
	},
	"S": {
		{"PassesDefended", 3000 * 0.75},
		{"SoloTackles", 400 * 0.75},
		{"FumblesForced", 1500 * 0.75},
		{"DefensiveTouchdowns", 5000 * 0.75},
		{"Interceptions", 3500 * 0.75},
	},
}

var PositionsX = map[string][]string{
	"QB": {"QB"},
	"WR": {"WR", "TE"},
	"RB": {"RB", "FB"},
	"DE": {"DL", "DT", "NT", "DE"},
	"LB": {"LB", "ILB", "OLB", "MLB"},
	"DB": {"DB", "CB"},
	"S":  {"SS", "FS", "S"},
}

var MeanValue = map[string]float64{
	"QB": 3077.029411764706,
	"RB": 2371.8676470588234,
	"WR": 2732.830882352941,
	"DE": 2594.485294117647,
	"DB": 2026.3235294117646,
	"LB": 2819.8529411764707,
	"S":  2114.1176470588234,
}

var ColumnData = []string{
	"Interceptions",
	"DefensiveTouchdowns",
	"PassesDefended",
	"Sacks",
	"Sacks",
	"FumblesForced",
	"TacklesForLoss",
	"PassingCompletions",
	"PassingYards",
	"PassingTouchdowns",
	"PassingInterceptions",
	"RushingYards",
	"RushingTouchdowns",
	"Receptions", "ReceivingYards", "ReceivingTouchdowns",
	"SoloTackles",
}
var PlayerTiers = map[string]float64{
	"rookie": 1,
	"vet":    1.15,
	"hof":    1.25,
	"goat":   1.5,
}

type PlayerCalculation struct {
	Points           int
	Team             string
	Position         string
	Level            int
	Reward           int
	MatchesPlayed    int
	PositionCategory string
	PlayerName       string
	PlayerID         string
}
