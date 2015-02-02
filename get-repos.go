package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	githubApiUserUrl = "https://api.github.com/users/marthjod"
)

type UserResponse struct {
	ReposUrl string `json:"repos_url"`
}

type RepoInfo struct {
	Name     string `json:"name"`
	CloneUrl string `json:"clone_url"`
}

func cloneRepo(repo RepoInfo) error {
	var (
		err error
	)

	fmt.Printf("Cloning %s... (%s)\n", repo.Name, repo.CloneUrl)

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
		go func(repo RepoInfo) {
			cloneRepo(repo)
			w.Done()
		}(repo)
	}

	w.Wait()
}
