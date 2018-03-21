package go_league_of_legends

type MatchParticipantTimeline struct {
	ParticipantId               int          `json:"participantId"`
	CreepsPerMinDeltas          PerMinDeltas `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              PerMinDeltas `json:"xpPerMinDeltas"`
	GoldPerMinDeltas            PerMinDeltas `json:"goldPerMinDeltas"`
	CsDiffPerMinDeltas          PerMinDeltas `json:"csDiffPerMinDeltas"`
	XpDiffPerMinDeltas          PerMinDeltas `json:"xpDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     PerMinDeltas `json:"damageTakenPerMinDeltas"`
	DamageTakenDiffPerMinDeltas PerMinDeltas `json:"damageTakenDiffPerMinDeltas"`

	Role string `json:"role"`
	Lane string `json:"lane"`
}
