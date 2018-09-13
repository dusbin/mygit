package main
import(
	"os"
	log "Sirupsen/logrus"
	myfile "comm/myfile"
)
func init_repo(repo_name string){
	root_path := myfile.GetCurrentPath()
	repo_dir :=  root_path
	if repo_name != "." {
		repo_dir = root_path+ repo_name +"/"
	}else{
		repo_dir = root_path
	}
	exist,err := myfile.PathExists(repo_dir)
	if err != nil{
		log.Infof("dir[%s] failed",repo_dir);
	}
	if exist != true{
		if err := os.Mkdir(repo_dir,0600);err != nil{
			log.Errorf("Mkdir %s error.%v",repo_dir,err)
		}
	}
	repo_dir = repo_dir + ".git"
	exist,err = myfile.PathExists(repo_dir)
	if err != nil{
		log.Infof("dir[%s] failed",repo_dir);
	}
	if exist != true{
		if err := os.Mkdir(repo_dir,0600);err != nil{
			log.Errorf("Mkdir %s error.%v",repo_dir,err)
		}
	}
	return
}