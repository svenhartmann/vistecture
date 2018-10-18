package core

import "strings"

type Dependency struct {
	Reference      string            `json:"reference" yaml:"reference"`
	Description    string            `json:"description" yaml:"description"`
	Relationship   string            `json:"relationship" yaml:"relationship"`
	IsSameLevel    bool              `json:"isSameLevel" yaml:"isSameLevel"`
	IsBrowserBased bool              `json:"isBrowserBased" yaml:"isBrowserBased"`
	Status         string            `json:"status" yaml:"status"`
	Properties     map[string]string `json:"properties" yaml:"properties"`
}

// Returns the name of the "component" and "service" this dependecy points to
// service might be empty if the dependency just defined the component
func (Dependency *Dependency) GetComponentAndServiceNames() (string, string) {
	if strings.Contains(Dependency.Reference, ".") {
		splitted := strings.Split(Dependency.Reference, ".")
		return splitted[0], splitted[1]
	}
	return Dependency.Reference, ""
}

func (Dependency *Dependency) GetComponentName() string {
	if strings.Contains(Dependency.Reference, ".") {
		splitted := strings.Split(Dependency.Reference, ".")
		return splitted[0]
	}
	return Dependency.Reference
}

func (Dependency *Dependency) GetServiceName() string {
	_, s := Dependency.GetComponentAndServiceNames()
	return s
}

func (Dependency *Dependency) GetComponent(Project *Project) (Application, error) {
	componentName, _ := Dependency.GetComponentAndServiceNames()
	return Project.FindApplication(componentName)
}
