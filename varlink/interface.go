// Package varlink implements the varlink protocol.
// See http://varlink.org for more information.
package varlink

// Interface defines a varlink interface.
type Interface interface {
	GetName() string
	GetDescription() string
	IsMethod(methodname string) bool
}

// InterfaceDefinition represents an active interface derived from
// a varlink interface description.
type InterfaceDefinition struct {
	Interface
	Name        string
	Description string
	Methods     map[string]struct{}
}

// GetName returns the reverse-domain varlink interface name.
func (d *InterfaceDefinition) GetName() string {
	return d.Name
}

// GetDescription returns the interface description. The interface description can be retrieved from
// the running service by calling the org.varlink.service.GetInterfaceDescription() method
func (d *InterfaceDefinition) GetDescription() string {
	return d.Description
}

// IsMethod indicates if the given method name is defined in the varlink interface description.
func (d *InterfaceDefinition) IsMethod(methodname string) bool {
	_, ok := d.Methods[methodname]
	return ok
}
