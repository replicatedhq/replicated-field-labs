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
	// alternative to JSON, handy for exports from sheets / google forms
	EnvironmentsCSV string
	LabsJSON        string

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
