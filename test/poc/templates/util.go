package templates

import "fmt"

// Helper function: Converts CamelCase to snake_case
func snake_case(input string) string {
	var result []rune
	for i, char := range input {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, char)
	}
	return string(result)
}

// Helper function: Determines the SQL data type
func db_type(field Field) string {
	if field.DbType != "" {
		return field.DbType
	}
	switch field.Type {
	case "int", "bigint":
		return "BIGINT"
	case "string":
		if field.Name == "Id" || field.Name[len(field.Name)-2:] == "Id" {
			return "UUID"
		}
		return "TEXT"
	case "bool":
		return "BOOLEAN"
	case "float", "float64":
		return "DOUBLE PRECISION"
	default:
		return "TEXT"
	}
}

// Helper function: Generates field constraints
func field_constraints(field Field) string {
	constraints := ""
	if field.IsRequired {
		constraints += "NOT NULL "
	}
	if field.IsUnique {
		constraints += "UNIQUE "
	}
	if field.DefaultValue != "" {
		constraints += "DEFAULT '" + field.DefaultValue + "' "
	}
	return constraints
}

// Helper function: Generates foreign key constraints
func foreign_keys(fields []Field) string {
	foreignKeys := ""
	for _, field := range fields {
		if field.ForeignKey {
			foreignKeys += fmt.Sprintf(`
    CONSTRAINT fk_%s FOREIGN KEY (%s) REFERENCES %s(id),
`, snake_case(field.Name), snake_case(field.Name), foreign_table(field.Name))
		}
	}
	return foreignKeys
}

// Helper function: Infers foreign table name from field name
func foreign_table(fieldName string) string {
	// Assuming foreign table name is the plural form of fieldName without "Id"
	if len(fieldName) > 2 && fieldName[len(fieldName)-2:] == "Id" {
		return snake_case(fieldName[:len(fieldName)-2]) + "s"
	}
	return snake_case(fieldName) + "s"
}
