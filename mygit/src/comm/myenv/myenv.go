package myenv
import (
	"comm/myfile"
	"strings"
)
func GetRepoRoot()(files string,flag bool){
	gitrepo:=".git"
	i := 0
	currentPath:=myfile.GetCurrentPath()
	BEGIN:	
		i = strings.LastIndex(currentPath,"/") //截取最后一个/作为路径
		if i == 0 {
			return "",false
		}	
		currentPath = string(currentPath[0:i])
		files=currentPath+"/"+gitrepo
		if myfile.IsDirExist(files){
			return files,true
		}else{
			if currentPath == "/" {
				return "",false
			}
		}
		currentPath = myfile.GetParentDir(currentPath)
	goto BEGIN
	return "",false
}
