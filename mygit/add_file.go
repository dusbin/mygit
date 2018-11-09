package main
import(
	"fmt"
	mymd5 "comm/mymd5"
	myfile "comm/myfile"
)
/*
 * 添加文件到仓库主入口
 */
func add_file(file_path,RepoRoot string){
	I:=myfile.StrStr(file_path,"/.git/")
	if I > 0 {
		return
	}
	md5Value,err := mymd5.Md5SumFile(file_path)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	add_info := fmt.Sprintf("100644 %x 0    %s\n",md5Value,file_path)
	Index := RepoRoot + "/index"
	//fmt.Println("add_info:",add_info)
	myfile.AppendStringToFile(Index,add_info)
	return
}
