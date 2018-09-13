package main
import(
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
	myfile.Mkdir(repo_dir)
	repo_dir = repo_dir + ".git"
	myfile.Mkdir(repo_dir)
	hooks_path := repo_dir + "/hooks"
	myfile.Mkdir(hooks_path)
	info_path := repo_dir + "/info"
	myfile.Mkdir(info_path)
	exclude_file := info_path + "/exclude"
	myfile.CreateFile(exclude_file)
	objects_path := repo_dir + "/objects"
	myfile.Mkdir(objects_path)
	info2_path := objects_path + "/info"
	myfile.Mkdir(info2_path)
	pack_path := objects_path + "/pack"
	myfile.Mkdir(pack_path)
	refs_path := repo_dir + "/refs"
	myfile.Mkdir(refs_path)
	heads_path := refs_path + "/heads"
	myfile.Mkdir(heads_path)
	tags_path := refs_path + "/tags"
	myfile.Mkdir(tags_path)
	config_file := repo_dir + "/config"
	myfile.CreateFile(config_file)
	descruption_file := repo_dir + "/description"
	myfile.CreateFile(descruption_file)
	HEAD_file := repo_dir + "/HEAD"
	myfile.CreateFile(HEAD_file)
	return
}