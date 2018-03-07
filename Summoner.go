package go_league_of_legends



type Summoner struct {
	Id int `json:"id"`
	AccountId int `json:"accountId"`
	Name string `json:"name"`
	ProfileIconId int `json:"profileIconId"`
	RevisionDate int `json:"revisionDate"`
	SummonerLevel int `json:"summonerLevel"`
}