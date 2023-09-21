package device

type VariedOptions struct {
	Path  string        `json:"path"`
	Opts  []interface{} `json:"opts"`
	Probs []string      `json:"probs"`
}

type VariedRange struct {
	Path  string `json:"path"`
	Range []int  `json:"range"`
	Unit  string `json:"unit"`
}

type Device struct {
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Static map[string]interface{} `json:"static"`
	Varies interface{}            `json:"varies"`
}
