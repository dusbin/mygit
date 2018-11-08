package main
import (
	"fmt"
	//log "Sirupsen/logrus"
	cli "urfave/cli"
	myfile "comm/myfile"
	myenv "comm/myenv"
	"os"
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
		if len(context.Args()) < 1 {
			return fmt.Errorf("${file_path} is error.\ne.g.:\nmygit add .\nmygit add ${file_path}\n")
		}
		/**
			可以使用以下命令进行添加
			mygit add . #添加当前目录
			mygit /bin/c* #添加*匹配文件
			mygit /opt/ #添加目录
			也可以传入多个目录作为多参数。
		 */
		RepoRoot,_ := myenv.GetRepoRoot()
		if RepoRoot == ""{
			fmt.Println("No Git Repo in this path[",myfile.GetCurrentPath(),"]")
			fmt.Println("Please Create A New Repo or Choose Another.")
			return nil
		}
		count :=0
		for _,file_path:= range os.Args[2:]{//mygit add filename1 filename2 ....
			f,_:=os.Stat(file_path)
			if f.IsDir(){
				files,_:=myfile.GetAllFileFromPath(file_path)//自己写的针对文件处理的包，该方法为获取指定目录下所有的文件
				for _,file:=range files{
					add_file(file,RepoRoot)
					count++
				}
				continue
			}
			add_file(file_path,RepoRoot)
			count++
		}
		fmt.Println("you add files(",count,")")
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