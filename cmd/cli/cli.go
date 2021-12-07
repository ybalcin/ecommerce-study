package cli

import (
	"bufio"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/ports"
	"log"
	"os"
)

func Execute() {
	for i := 1; i <= 5; i++ {
		cli := ports.NewCLI()

		pwd, _ := os.Getwd()

		file, err := os.Open(fmt.Sprintf("%s/cmd/cli/cases/c_%d.txt", pwd, i))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("\n> > -------------- READING c_%d.txt -------------- < < \n", i)

		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanLines)
		var cmds []string

		for scanner.Scan() {
			cmds = append(cmds, scanner.Text())
		}

		file.Close()

		for _, cmd := range cmds {
			cli.Execute(cmd)
		}

		fmt.Printf("> > -------------- END c_%d.txt -------------- < < \n", i)
	}
}
