package main
import (
	"fmt"
	//log "Sirupsen/logrus"
	cli "urfave/cli"
)
//初始化命令定义，定义了init的具体操作，此操作为内部方法，禁止外部调用
var dus = cli.Command{
	Name: 	"dus",
	Usage: 	"my own command",
	Action: func(context *cli.Context) error{
		fmt.Printf("error my own command\n")
		return nil
	},
}
var help = cli.Command{
	Name: 	"help",
	Usage: 	"mygit help",
	Action: func(context *cli.Context) error{
		fmt.Printf("mygit\n")
		fmt.Printf("    help\n")
		fmt.Printf("    initrepo\n")
		fmt.Printf("    dus\n")
		return nil
	},
}
var initrepo = cli.Command{
	Name: 	"init",
	Usage: 	"mygit init\n	mygit init .\n	mygit init repo_name",
	Action: func(context *cli.Context) error{
		fmt.Printf("this is init\n")
		return nil
	},
}