package cfg

type nodes struct {
	Node []node `yaml:"node"`
}

type node struct {
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Address      string `yaml:"address"`
	IdentityFile string `yaml:"identity_file"`
	Port         string `yaml:"port"`
}
