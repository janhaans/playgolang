package main

import (
	"fmt"

	"github.com/janhaans/playgolang/githubrepos"
)

func main() {
	var u string
	fmt.Print("What user ? ")
	fmt.Scanf("%s", &u)

	var repos *[]githubrepos.Repo = githubrepos.GetGitHubRepos(u)

	fmt.Printf("User %s has the following repositories:\n", u)
	for _, v := range *repos {
		fmt.Printf("repository = %s; owner = %s\n", v.Name, v.Owner.Login)
	}
}
