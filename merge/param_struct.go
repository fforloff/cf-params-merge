package merge

// Param type
type Param struct {
	ParameterKey     string `json:"ParameterKey"`
	ParameterValue   string `json:"ParameterValue,omitempty"`
	UsePreviousValue bool   `json:"UsePreviousValue,omitempty"`
	ResolvedValue    string `json:"ResolvedValue,omitempty"`
}
