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
		fmt.Printf("    init\n")
		fmt.Printf("	add")
		fmt.Printf("    dus\n")
		fmt.Printf("    usedb\n")
		return nil
	},
}
var initrepo = cli.Command{
	Name: 	"init",
	Usage: 	"mygit init .\n	mygit init ${repo_name}",
	Action: func(context *cli.Context) error{
		if len(context.Args()) != 1 {
			return fmt.Errorf("${repo_name} is error.\ne.g.:\nmygit init .\nmygit init ${repo_name}\n")
		}
		repo_name := context.Args().Get(0)
		init_repo(repo_name)
		return nil
	},
}
var addfile = cli.Command{
	Name: 	"add",
	Usage: 	"mygit add .\n	mygit add ${file_path}",
	Action: func(context *cli.Context) error{
		if len(context.Args()) != 1 {
			return fmt.Errorf("${file_path} is error.\ne.g.:\nmygit add .\nmygit add ${file_path}\n")
		}
		file_path := context.Args().Get(0)
		add_file(file_path)
		return nil
	},
}
var usedb = cli.Command{
	Name: 	"usedb",
	Usage: 	"mygit usedb .\n	mygit usedb ${file_path}",
	Action: func(context *cli.Context) error{
		//if len(context.Args()) != 1 {
		//	return fmt.Errorf("${file_path} is error.\ne.g.:\nmygit usedb .\nmygit usedb ${file_path}\n")
		//}
		use_db()
		return nil
	},
}