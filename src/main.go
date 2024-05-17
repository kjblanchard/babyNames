package main

import (
	"github.com/kjblanchard/baby_names/cli"
	"github.com/kjblanchard/baby_names/utils"
)


func main() {
	utils.ChangeToWorkDir()
	var mode = cli.CheckForType()
	switch mode {
	case 1:
		cli.RateNames()
	case 2:
		cli.ShowResults()
	case 3:
		cli.CompareResults()
	}
}
