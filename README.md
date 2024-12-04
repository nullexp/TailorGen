# **MetaGen**

**MetaGen** is a flexible, meta-driven static code generator designed to transform metadata (e.g., JSON, YAML) into tailored, production-ready code. By leveraging customizable templates, MetaGen empowers developers to scaffold projects, generate domain-specific components, and integrate seamlessly with their workflow.

---

## **Key Features**
1. **Multi-Input Support**  
   - Accepts metadata in multiple formats: JSON, YAML, TOML, and XML.  
   - Automatically detects the input format to streamline the process.  

2. **Template-Based Code Generation**  
   - Built-in templates for common use cases (e.g., entities, repositories, services).  
   - Supports exporting embedded templates for developer customization.  

3. **Customizable Output**  
   - Generate code in Go, Python, Java, or other languages.  
   - Configure naming conventions (e.g., PascalCase, snake_case).  

4. **Domain-Driven Design (DDD) Support**  
   - Scaffolding for entities, value objects, aggregates, and repositories.  
   - Built-in templates for bounded contexts and event-driven architectures.  

5. **Interactive CLI**  
   - Command-line interface for generating, previewing, and exporting templates.  

6. **Live Preview**  
   - Visualize generated code before writing it to disk.  

7. **Plugin Architecture**  
   - Extend MetaGen with custom parsers or additional language support.  

---

## **Installation**
### Prerequisites
- Go (1.22+)

### Install MetaGen
Clone the repository and build the binary:
```bash
git clone https://github.com/nullexp/metagen.git
cd metagen
go build -o metagen
```

Alternatively, use `go install`:
```bash
go install github.com/nullexp/metagen@latest
```

---

## **Usage**

### **Basic Code Generation**
Generate code from a metadata file using built-in templates:
```bash
metagen --input models.yaml --output ./generated
```

### **Export Embedded Templates**
Export built-in templates for customization:
```bash
metagen export-templates --output ./templates
```

### **Use Custom Templates**
Specify a directory of customized templates:
```bash
metagen --input models.yaml --template-dir ./templates --output ./generated
```

---

## **Example Input**
Here’s an example of a metadata file in YAML format:

```yaml
PackageName: entity
EntityName: Server
Fields:
  - Name: Id
    Type: string
    JsonTag: id
    Required: true
  - Name: Name
    Type: string
    JsonTag: name
  - Name: CreatedAt
    Type: time.Time
    JsonTag: created_at
  - Name: UpdatedAt
    Type: time.Time
    JsonTag: updated_at
```

---

## **Generated Output**
Based on the example input, MetaGen generates Go code like this:

```go
package entity

import "time"

type Server struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s Server) IsIdEmpty() bool {
	return s.Id == ""
}

func (s *Server) SetId(id string) {
	s.Id = id
}
```

---

## **Roadmap**
MetaGen is evolving! Upcoming features include:
- Support for additional languages (e.g., Python, Java).
- Template validation and versioning.
- Integration with CI/CD pipelines.
- IDE plugins for VSCode and JetBrains.

---

## **Contributing**
We welcome contributions! To get started:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request.

---

## **License**
This project is licensed under the MIT License.

---

## **Contact**
Have questions or need help? Reach out:  
- **GitHub Issues:** [MetaGen Issues](https://github.com/nullexp/metagen/issues)  
- **Email:** hopegolestany@gmail.com

---

Feel free to modify this as needed. Let me know if you’d like any sections expanded or tailored further!