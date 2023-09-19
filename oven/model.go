package oven

type Model struct {
	Type   string                 `json:"type"`
	Static map[string]interface{} `json:"static"`
	Varies interface{}            `json:"varies"`
}

/*func modelForge(definition map[string]interface{}) error {

}*/
