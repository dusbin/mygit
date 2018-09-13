package myfile
import(
	"os"
	"strings"
	"path/filepath"
)
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