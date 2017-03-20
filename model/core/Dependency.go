package core

import "strings"

type Dependency struct {
	Reference string `json:"reference"`
	Relationship string `json:"relationship"`
	IsSameLevel bool `json:"isSameLevel"`
	Resilience string
	IsBrowserBased bool `json:"isBrowserBased"`

}

// Returns the name of the "component" and "service" this dependecy points to
// service might be empty if the dependency just defined the component
func (Dependency *Dependency) GetComponentAndServiceNames() (string, string) {
	if (strings.Contains(Dependency.Reference,".")) {
		splitted := strings.Split(Dependency.Reference, ".")
		return splitted[0], splitted[1];
	}
	return Dependency.Reference, ""
}

func (Dependency *Dependency) GetComponent(Project *Project) (Component, error) {
	componentName, _ := Dependency.GetComponentAndServiceNames()
	return Project.FindComponent(componentName)
}