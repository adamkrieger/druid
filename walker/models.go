package walker

//EntityTrees - Parent input structure
type EntityTrees struct {
	Entities map[string][]*Entity `yaml:"entities"`
}

//Entity - An Entity
type Entity struct {
	Name      string            `yaml:"name"`
	Atom      string            `yaml:"atom"`
	Details   map[string]string `yaml:"details"`
	Relations []*Relation       `yaml:"relations"`
}

//Relations - Each relation connecting two entities
type Relation struct {
	Type string `yaml:"type"`
	Atom string `yaml:"atom"`
}

//Node - A Node
type Node struct {
	Atom      string            `json:"atom"`
	Type      string            `json:"type"`
	Name      string            `json:"name"`
	Relations []string          `json:"rels"`
	Details   map[string]string `json:"details"`
}

//Edge - An Edge
type Edge struct {
	Atom      string `json:"atom"`
	Type      string `json:"type"`
	PeerLeft  string `json:"peerleft"`
	PeerRight string `json:"peerright"`
}
