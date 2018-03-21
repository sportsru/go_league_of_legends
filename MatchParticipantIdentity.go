package go_league_of_legends

type MatchParticipantIdentity struct {
	ParticipantId int         `json:"participantId"`
	Player        MatchPlayer `json:"player"`
}
