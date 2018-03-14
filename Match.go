package go_league_of_legends

type Match struct {
    PlatformId string `json:"platformId"`
    GameId     int    `json:"gameId"`
    Champion   int    `json:"champion"`
    Queue      int    `json:"queue"`
    Season     int    `json:"season"`
    Timestamp  int    `json:"timestamp"`
    Role       string `json:"role"`
    Lane       string `json:"lane"`
}
