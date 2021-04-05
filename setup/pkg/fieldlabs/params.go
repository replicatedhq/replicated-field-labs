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
	APIToken         string
	APIOrigin        string
	GraphQLOrigin    string
	KURLSHOrigin     string
	LabsJSON         string
}

