package cfg

type node struct {
	Nodes []nodes `yaml:"node"`
}

type nodes struct {
	Name    string `yaml:"name"`
	User    string `yaml:"user"`
	Address string `yaml:"address"`
}
