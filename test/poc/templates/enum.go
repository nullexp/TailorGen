package templates

type Enum struct {
	Name        string      `json:"name" yaml:"name"`               // Name of the enum type
	Description string      `json:"description" yaml:"description"` // Description of the enum
	Fields      []EnumField `json:"fields" yaml:"fields"`           // List of fields (values) in the enum
}

type EnumField struct {
	Name  string `json:"name" yaml:"name"`   // Name of the enum value (Go constant name)
	Value string `json:"value" yaml:"value"` // Actual string value of the enum
}
