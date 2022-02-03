package cfg

type node struct {
	Nodes []nodes `yaml:"node"`
}

type nodes struct {
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Address      string `yaml:"address"`
	IdentityFile string `yaml:"identity_file"`
	Port         string `yaml:"port"`
}
