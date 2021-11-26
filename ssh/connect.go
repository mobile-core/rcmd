package ssh

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	connect()
}

func connect() {
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	in, _ := session.StdinPipe()

	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('n')
		fmt.Fprint(in, str)
	}

}
