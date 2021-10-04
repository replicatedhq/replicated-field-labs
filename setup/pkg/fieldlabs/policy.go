package fieldlabs

type Policy struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Definition  string `json:"definition"`
}

// Copied from vendor-api/policy
// PolicyDefinition implements the JSON schema a user can write to define a policy
type PolicyDefinition struct {
	V1 PolicyDefinitionV1 `json:"v1"`
}

// Copied from vendor-api/policy
// PolicyDefinitionV1 implements the V1 JSON schema for a policy definition
type PolicyDefinitionV1 struct {
	Name      string            `json:"name"`
	Resources PolicyResourcesV1 `json:"resources"`
}

// Copied from vendor-api/policy
// PolicyResourcesV1 implements the resources list in a V1 JSON policy definition
type PolicyResourcesV1 struct {
	Allowed []string `json:"allowed"`
	Denied  []string `json:"denied"`
}

type PolicyListItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
