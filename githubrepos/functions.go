package githubrepos

import (
	"encoding/json"
	"net/http"
)

func GetGitHubRepos(name string) *[]Repo {
	resp, err := http.Get(URLrepos + name + "/repos")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	repos := []Repo{}
	json.NewDecoder(resp.Body).Decode(&repos)

	return &repos
}
