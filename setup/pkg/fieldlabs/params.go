package fieldlabs

type Action int

const (
	ActionCreate Action = iota
	ActionDestroy
)

type Params struct {
	Action           Action
	ParticipantEmail string
	LabsJSON         string
	LabSlug          string

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
