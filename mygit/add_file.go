package main
import(
	"fmt"
	mymd5 "comm/mymd5"
)
/*
 * 添加文件到仓库主入口
 */
func add_file(file_path,RepoRoot string){
	//fmt.Printf("file_path : [%s]\n",file_path)
	md5Value,err := mymd5.Md5SumFile(file_path)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%x 0644 %s %s  \n",md5Value,RepoRoot,file_path)
	return
}
