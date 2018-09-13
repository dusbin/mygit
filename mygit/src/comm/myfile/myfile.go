package myfile
import(
	"os"
	"strings"
	"path/filepath"
	//log "../../Sirupsen/logrus"
)
/*
 * 创建目录
 */
func Mkdir(path string){
	exist,err := PathExists(path)
	if err != nil{
		//log.Infof("dir[%s] failed",path);
	}
	if exist != true{
		if err := os.MkdirAll(path,0600);err != nil{ //递归创建目录
			//log.Errorf("Mkdir %s error.%v",path,err)
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
			//log.Errorf("create file %s error.%v",path,error)
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