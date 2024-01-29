package main
import(
	"fmt"
	"github.com/JonecoBoy/obsidian-sync/config"
	"log"
	"os/exec"
	"strings"
)


//var string obsidianDir;
//  ??: Untracked file. This means a file is present in your working directory, but Git is not tracking it.
//
//	M: Modified. The file has been modified, and the changes haven't been committed yet.
//
//  A: Added. A new file has been added to the staging area (index) and is ready to be committed.
//
//	D: Deleted. A file has been deleted, and the deletion is staged for the next commit.
//
//	R: Renamed. The file has been renamed, and the change is staged for the next commit.
//
//	C: Copied. The file has been copied, and the change is staged for the next commit.
//
//	U: Unmerged. This indicates that a file is unmerged due to a conflict. You need to resolve the conflict before you can commit.
//
//	!!: Missing or deleted. The file is missing or has been deleted in the working directory.
//
//	DD: Unmerged, both deleted. This indicates that both sides of the conflict deleted the same file.
//
//	AU: Unmerged, added by us. This indicates that you added the file, but the other side deleted it.
//
//	UD: Unmerged, deleted by them. This indicates that you deleted the file, but the other side added it.
//
//	UA: Unmerged, added by them. This indicates that they added the file, but you deleted it.
//
//	DU: Unmerged, deleted by us. This indicates that they deleted the file, but you added it.


func main() {
	config,err := config.GetConfig()
	fmt.Println(config)
	if err != nil{
		fmt.Println(err)
	}

	err = CheckGitCli()
	GitPull()
	changes,err := GitStatus()
	if err != nil{
		fmt.Println(err)
	}
	
	
	if(len(changes) > 0 ){
		err =GitAddAndPush(&changes)
		if err != nil{
			fmt.Println(err)
		}
	}else{
		log.Fatal("No changes")
	}
	
	
}

func CheckGitCli() error {
	cmd := exec.Command("git", "-v")
	output, err := cmd.Output()
	if err != nil {
		//fmt.Println("Error:", err)
		return fmt.Errorf("Git is not installed")
	}
	
	 gitExists := strings.Contains(strings.ToLower(string(output)),"git version")
	 if !gitExists {
		 return fmt.Errorf("Git is not installed")
	}

	return nil
}

func GitStatus() ([]string, error){
	cmd := exec.Command("git", "status","-s")
	output, err := cmd.Output()
	if err != nil {
		return []string{},fmt.Errorf("Error in git status")
	}
	if string(output) == "" {
		return []string{},fmt.Errorf("No Changes")
	}
	
	lines := strings.Split(string(output), "\n")
	return lines,err
	
}

func GitPull() error{
	cmd := exec.Command("git", "pull","-q")
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error in git status")
	}
	return nil
}

func GitAddAndPush(changes *[]string) error{
	if(len(*changes) < 1 ){
		log.Fatal("No changes")
		return fmt.Errorf("Error in git status")
	}
	
	cmd := exec.Command("git", "add",".")
	output, err := cmd.Output()
	println(output)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error in git add")
	}
	

	changesMessage := strings.Join(*changes, ",")
	// Limit the string to a maximum of 255 characters
	if len(changesMessage) > 255 {
		changesMessage = changesMessage[:255]
	}
	
	cmd = exec.Command("git", "commit","-m",changesMessage)
	output, err = cmd.Output()
	println(output)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error in git commit")
	}
	
	cmd = exec.Command("git", "push")
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("Error in git add")
	}
	
	return nil
}
