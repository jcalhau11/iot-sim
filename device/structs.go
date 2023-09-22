package device

type Varied interface {
	Force(device *Device) error
}

type VariedOptions struct {
	Path  string        `json:"path"`
	Type  string        `json:"type"`
	Opts  []interface{} `json:"opts"`
	Probs []string      `json:"probs"`
}

type VariedRange struct {
	Path  string `json:"path"`
	Type  string `json:"type"`
	Range []int  `json:"range"`
	Unit  string `json:"unit"`
}

type DefaultVaried struct {
	Type string `json:"type"`
}

type Device struct {
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Static map[string]interface{} `json:"static"`
	Varies []Varied               `json:"varies"`
}
