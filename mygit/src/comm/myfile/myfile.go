package myfile
import(
	"os"
	"strings"
	"path/filepath"
	"fmt"
	"io/ioutil"
	//log "../../Sirupsen/logrus"
)
//获取指定目录下所有文件
func GetAllFileFromPath(dirpath string)(files []string,err error){
	files,dirs,_:=getFilesAndDirs(dirpath)
	for _,table:=range dirs{
		temp,_,_:=getFilesAndDirs(table)
		for _,temp1 := range temp{
			files = append(files,temp1)
		}
		xfiles,_:=getAllFiles(dirpath)
		for _,file:= range xfiles{
			files = append(files,file)
			//fmt.Println("file:",file)
		}
	}
	return files,nil
}
func getAllFiles(dirpath string) (files []string,err error){
	var dirs []string
	dir,err:=ioutil.ReadDir(dirpath)
	if err != nil{
		return nil,err
	}
	
	for _,fi:=range dir{
		if fi.IsDir(){
			dirs = append(dirs,dirpath+"/"+fi.Name())
			getAllFiles(dirpath+"/"+fi.Name())
		}else{
			files = append(files,dirpath+"/"+fi.Name())
		}
	}
	for _,table := range dirs{
		temp,_:=getAllFiles(table)
		for _,temp1:=range temp{
			files = append(files,temp1)
		}
	}
	return files,nil
}
func getFilesAndDirs(dirpath string)(files []string,dirs []string,err error){
	dir,err := ioutil.ReadDir(dirpath)
	if err != nil{
		return nil,nil,err
	}
	for _,fi:=range dir{
		if fi.IsDir(){//目录的处理，递归
			dirs = append(dirs,dirpath+"/"+fi.Name())
			getFilesAndDirs(dirpath+"/"+fi.Name())
		}else{//处理文件
			files = append(files,dirpath+"/"+fi.Name())
		}
	}
	return files,dirs,nil
}
/*
 * 向文件中写入字符串，覆盖写入。
 */
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