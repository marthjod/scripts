package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

// TODO flag.*
const (
	githubApiUserUrl = "https://api.github.com/users/marthjod"
	cloneDir         = "/tmp/cloned-repos"
)

type UserResponse struct {
	ReposUrl string `json:"repos_url"`
}

type RepoInfo struct {
	Name     string `json:"name"`
	CloneUrl string `json:"clone_url"`
}

func gitClone(repo string) error {
	var (
		cmd *exec.Cmd
	)

	cmd = exec.Command("git", "clone", repo)
	return cmd.Run()
}

func mkCloneDir(cloneDir string) error {
	var (
		err error
	)

	if _, err = os.Stat(cloneDir); os.IsNotExist(err) {
		err = os.MkdirAll(cloneDir, 0755)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Printf("Created %s\n", cloneDir)
		}
	}

	return err
}

func cloneRepo(repo RepoInfo) error {
	var (
		err error
	)

	err = gitClone(repo.CloneUrl)
	if err != nil {
		fmt.Printf("Error cloning %s (%s): %s\n", repo.Name, repo.CloneUrl, err.Error())
	} else {
		fmt.Printf("Cloned %s (%s).\n", repo.Name, repo.CloneUrl)
	}
	return err
}

func main() {
	var (
		err           error
		resp          *http.Response
		body          []byte
		userResponse  UserResponse
		reposResponse []RepoInfo
		w             sync.WaitGroup
		// save os.Stat() calls
		cloneDirExists bool
	)

	resp, err = http.Get(githubApiUserUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &userResponse)

	resp, err = http.Get(userResponse.ReposUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &reposResponse)

	w.Add(len(reposResponse))
	for _, repo := range reposResponse {

		if !cloneDirExists {
			err = mkCloneDir(cloneDir)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			} else {
				cloneDirExists = true
				err = os.Chdir(cloneDir)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}
		}

		go func(repo RepoInfo) {

			cloneRepo(repo)
			w.Done()
		}(repo)
	}

	w.Wait()
}
