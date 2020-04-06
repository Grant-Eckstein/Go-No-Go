package gonogo
// Go No Go
// A GoLang Connection Orchestration Framework
//


import (
	"strings"
)

// Steps group each function by role.
// Each function must take a singular input and output as a byte slice.
type Step struct {
	Role string
	RecvFunc func(data []byte) []byte
}

// Used to create a new step. Should not be necessary for users.
func NewStep (role string, recvFunc func(data []byte) []byte) Step{
	return Step{Role:role, RecvFunc:recvFunc}
}

// Mostly a placeholder for a Step slice, also has a name which is currently unused.
type Role struct {
	Name string
	Steps []Step
}

// Add a step to a role, used only in (*Negotiation) AddStep.
func (r *Role) addStep(step Step) {
	r.Steps = append(r.Steps, step)
}


// The negitiation type is the main interface for this project.
// It is comprised of a description (which is currently unused) and a Step slice.
type Negotiation struct {
	Steps []Step
	Description string
}

// Used to create a new Negotiation, given a description which is currently unused.
func NewNegotiation(description string) Negotiation {
	return Negotiation{Steps:[]Step{}, Description:description}
}

// Each function you assign is repersented as a step.
// When adding a step to a negotiation, you assign a role name which functions as a group.
// Each function assigned this way must both accept and return a byte slice.
// This is a delibrate decision (rather than using a interface) to enforce
// reasonable ambiguity.
func (c *Negotiation) AddStep(role string, recvFunc func(data []byte) []byte) {
	s := NewStep(role, recvFunc)
	c.Steps = append(c.Steps, s)
}

// When you are ready to execute a role, you must first get the role from the Negotiation.
// Generally, users should not access this function.
func (c *Negotiation) GetRole(role string) Role {
	r := Role{Name:role}
	for _, step := range c.Steps {
		if strings.EqualFold(role, step.Role) {
			r.addStep(step)
		}
	}
	return r
}

// This function accesses the main functionality of this project.
// Specify a role name and data to input to the first step, in some cases (ex server-side roles)
// The Execute function returns the output of the last step, chaining the output of prevous steps
// to the input of the next.
func (n *Negotiation) Execute(role string, initData []byte) []byte {
	r := n.GetRole(role)
	data := r.Steps[0].RecvFunc(initData)
	for _, s := range r.Steps[1:] {
		data  = s.RecvFunc(data)
	}
	return data
}
