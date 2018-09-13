package main
import(
	myfile "comm/myfile"
)
/*
 * 初始化创建仓库目录主入口
 */
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
	myfile.WriteStringtoFile(exclude_file,"# git ls-files --others --exclude-from=.git/info/exclude\n# Lines that start with '#' are comments.\n# For a project mostly in C, the following would be a good set of\n# exclude patterns (uncomment them if you want to use them):\n# *.[oa]\n# *~\n")
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
	myfile.WriteStringtoFile(config_file,"[core]\n	repositoryformatversion = 0\n	filemode = false\n	bare = false\n	logallrefupdates = true\n	symlinks = false\n	ignorecase = true")
	descruption_file := repo_dir + "/description"
	myfile.CreateFile(descruption_file)
	myfile.WriteStringtoFile(descruption_file,"Unnamed repository; edit this file 'description' to name the repository.")
	HEAD_file := repo_dir + "/HEAD"
	myfile.CreateFile(HEAD_file)
	myfile.WriteStringtoFile(HEAD_file,"ref: refs/heads/master")
	return
}