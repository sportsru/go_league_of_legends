package go_league_of_legends

type MatchPlayer struct {
	PlatformId        string `json:"platformId"`
	AccountId         int    `json:"accountId"`
	SummonerName      string `json:"summonerName"`
	SummonerId        int    `json:"summonerId"`
	CurrentPlatformId string `json:"currentPlatformId"`
	CurrentAccountId  int    `json:"currentAccountId"`
	MatchHistoryUri   string `json:"matchHistoryUri"`
	ProfileIcon       int    `json:"profileIcon"`
}
