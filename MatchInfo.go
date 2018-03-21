package go_league_of_legends

type MatchInfo struct {
	GameId                int                        `json:"gameId"`
	PlatformId            string                     `json:"platformId"`
	GameCreation          int                        `json:"gameCreation"`
	GameDuration          int                        `json:"gameDuration"`
	QueueId               int                        `json:"queueId"`
	MapId                 int                        `json:"mapId"`
	SeasonId              int                        `json:"seasonId"`
	GameVersion           string                     `json:"gameVersion"`
	GameMode              string                     `json:"gameMode"`
	GameType              string                     `json:"gameType"`
	Teams                 []MatchTeam                `json:"teams"`
	Participants          []MatchParticipant         `json:"participants"`
	ParticipantIdentities []MatchParticipantIdentity `json:"participantIdentities"`
}
