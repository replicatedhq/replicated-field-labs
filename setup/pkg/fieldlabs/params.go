package fieldlabs

type Action int

const (
	ActionCreate Action = iota
	ActionDestroy
)

type Params struct {
	Action           Action
	ParticipantEmail string
	Branch           string
	TrackSlug        string

	APIToken      string
	APIOrigin     string
	GraphQLOrigin string
	KURLSHOrigin  string
	IDOrigin      string

	// vendor web user's email for sending invites, can't use api token
	InviterEmail string
	// vendor web user's password for sending invites, can't use api token
	InviterPassword string
	// sessionToken
	SessionToken string
}

type LambdaEvent struct {
	Action           string `json:"action"`
	ParticipantEmail string `json:"participant-email"`
	Branch           string `json:"branch"`
	TrackSlug        string `json:"track-slug"`
	InviterEmail     string `json:"inviter-email"`
	InviterPassword  string `json:"inviter-password"`
	APIToken         string `json:"api-token"`
}
