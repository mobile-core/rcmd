package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type strslice []string

func (s *strslice) String() string {
	return fmt.Sprintf("%v", multiflag)
}

func (s *strslice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var multiflag strslice

func main() {
	out, _ := exec.Command("kubectl", "get", "nodes", "-n", "f5gc", "-o", "wide").CombinedOutput()
	fmt.Println(string(out))

	data, _ := ioutil.ReadFile(".rcmd.yml")
	fmt.Println(string(data))

	//connect()

	//execute()

	//selectedNode()

	allNode()
}

//ssh login
func connect() {
	f := flag.String("node", "", "node name")
	flag.Parse()

	out, err := exec.Command("ssh", *f, "-i", "${HOME}/.ssh/id_rsa").Output()
	if err != nil {
		fmt.Println("connection failed")
	} else {
		fmt.Println(string(out))
		os.Exit(0)
	}
}

//exec command to a node
func execute() {
	f := flag.String("node", "", "node name")
	g := flag.String("command", "", "execute command")
	flag.Parse()

	out, err := exec.Command("ssh", *f, "-i", "${HOME}/.ssh/id_rsa", *g).Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}

}

//exec command to selected nodes
func selectedNode() {
	flag.Var(&multiflag, "node", "node name")
	g := flag.String("command", "", "execute command")
	flag.Parse()

	for i := 0; i < 10; i++ {
		out, err := exec.Command("ssh", multiflag[i], "-i", "${HOME}/.ssh/id_rsa", *g).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	}
}

//exec command to all nodes
func allNode() {
	nodes := []string{"master", "node1", "node2"}

	g := flag.String("command", "", "execute command")
	flag.Parse()

	for i := 0; i < len(nodes); i++ {
		out, err := exec.Command("ssh", nodes[i], "-i", "${HOME}/.ssh/id_rsa", *g).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	}
}
