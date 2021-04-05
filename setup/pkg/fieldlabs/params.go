package fieldlabs

type Action int

const (
	ActionCreate Action = iota
	ActionDestroy
)

type Params struct {
	NamePrefix       string
	Action           Action
	EnvironmentsJSON string
	LabsJSON         string

	InstanceJSONOutput string

	APIToken      string
	APIOrigin     string
	GraphQLOrigin string
	KURLSHOrigin  string

	// invite members based on Environment.Email
	InviteUsers bool
	// usually "Admin", but unique per team. required if InviteUsers is set
	RBACPolicyID string
}
