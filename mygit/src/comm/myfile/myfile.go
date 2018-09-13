package myfile
import(
	"os"
	"strings"
	"path/filepath"
	"fmt"
	//log "../../Sirupsen/logrus"
)
func WriteStringtoFile(path string,data string){
	exist,err := PathExists(path)
	if err != nil{
		fmt.Printf("path[%s] is not Exists\n",path)
	}
	if exist != true{
		CreateFile(path)
	}
	file1,error := os.Create(path)
	if error == nil{
		//file1.Write([]byte(data))
		file1.WriteString(data)
		file1.Close()
	}else{
		fmt.Printf("Write error\n")
	}
}
/*
 * 创建目录
 */
func Mkdir(path string){
	exist,err := PathExists(path)
	if err != nil{
		fmt.Printf("dir[%s] failed\n",path);
	}
	if exist != true{
		if err := os.MkdirAll(path,0600);err != nil{ //递归创建目录
			fmt.Printf("Mkdir %s error.%v\n",path,err)
		}
	}
}
/*
 * 创建文件
 */
func CreateFile(path string){
	_,error := os.Open(path)
	if error != nil{ //文件不存在
		_,error = os.Create(path)
		if error != nil{
			fmt.Printf("create file %s error.%v\n",path,error)
		}
	}

}
/*
 * 判断路径是否存在
 */
func PathExists(path string) (bool,error) {
	_, err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}
/*
 *获取当前程序运行路径
 */
func GetCurrentPath() string{
	s,err := filepath.Abs(os.Args[0])//s 为运行的程序的完整路径
	checkErr(err)
	i := strings.LastIndex(s,"/") //截取最后一个/作为路径
	path := string(s[0:i+1])
	return path
}
func checkErr(err error){
	if err != nil {
		panic(err)
	}
}