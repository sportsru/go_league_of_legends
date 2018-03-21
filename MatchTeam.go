package go_league_of_legends

type MatchTeam struct {
	TeamId               int         `json:"teamId"`
	Win                  string      `json:"win"`
	FirstBlood           bool        `json:"firstBlood"`
	FirstTower           bool        `json:"firstTower"`
	FirstInhibitor       bool        `json:"firstInhibitor"`
	FirstBaron           bool        `json:"firstBaron"`
	FirstDragon          bool        `json:"firstDragon"`
	FirstRiftHerald      bool        `json:"firstRiftHerald"`
	TowerKills           int         `json:"towerKills"`
	InhibitorKills       int         `json:"inhibitorKills"`
	BaronKills           int         `json:"baronKills"`
	DragonKills          int         `json:"dragonKills"`
	VilemawKills         int         `json:"vilemawKills"`
	RiftHeraldKills      int         `json:"riftHeraldKills"`
	DominionVictoryScore int         `json:"dominionVictoryScore"`
	Bans                 []MatchBans `json:"bans"`
}
