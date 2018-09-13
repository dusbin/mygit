package mymd5
import(
	"crypto/md5"
	"io/ioutil"
	//"os"
)
func Md5SumFile(file_path string)(value [md5.Size]byte,err error){
	data,err := ioutil.ReadFile(file_path)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return
}