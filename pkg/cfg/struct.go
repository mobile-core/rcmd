package cfg

type node struct {
	Nodes []nodes `yaml:"node"`
}

type nodes struct {
	HostName string `yaml:"hostname"`
	UserName string `yaml:"username"`
	Address  string `yaml:"address"`
}
