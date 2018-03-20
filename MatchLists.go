package go_league_of_legends

type MatchLists struct {
    Matches    []Match `json:"matches"`
    StartIndex int     `json:"startIndex"`
    EndIndex   int     `json:"endIndex"`
    TotalGames int     `json:"totalGames"`
}
