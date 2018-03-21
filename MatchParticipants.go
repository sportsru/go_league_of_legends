package go_league_of_legends

type MatchParticipant struct {
	ParticipantId             int                      `json:"participantId"`
	TeamId                    int                      `json:"teamId"`
	ChampionId                int                      `json:"championId"`
	Spell1Id                  int                      `json:"spell1Id"`
	Spell2Id                  int                      `json:"spell2Id"`
	HighestAchievedSeasonTier string                   `json:"highestAchievedSeasonTier"`
	Stats                     MatchParticipantStats    `json:"stats"`
	Timeline                  MatchParticipantTimeline `json:"timeline"`
}
