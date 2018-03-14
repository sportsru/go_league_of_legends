package go_league_of_legends

type MatchLists struct {
    Matches    []Match `json:"matches"`
    StartIndex int     `json:"startIndex"`
    EndIndex   string  `json:"endIndex"`
    TotalGames int     `json:"totalGames"`
}
