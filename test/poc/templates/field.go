package templates

type Field struct {
	Name          string `json:"name" yaml:"name"`
	Type          string `json:"type" yaml:"type"`
	JsonTag       string `json:"jsonTag" yaml:"json_tag"`
	Description   string `json:"description" yaml:"description"`
	Example       string `json:"example" yaml:"example"`
	DbType        string `json:"dbType" yaml:"db_type"`               // e.g., "VARCHAR(255)", "INT"
	IsRequired    bool   `json:"isRequired" yaml:"is_required"`       // If the field is required
	IsUnique      bool   `json:"isUnique" yaml:"is_unique"`           // If the field is unique
	IsExported    bool   `json:"isExported" yaml:"is_exported"`       // If the field is exported
	ForeignKey    bool   `json:"foreignKey" yaml:"foreign_key"`       // If the field is a foreign key
	DefaultValue  string `json:"defaultValue" yaml:"default_value"`   // Default value for the field
	ValidationTag string `json:"validationTag" yaml:"validation_tag"` // If the field is a foreign key
}
