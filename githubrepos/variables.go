package githubrepos

const URLrepos = "https://api.github.com/users/"

type Repo struct {
	Name        string
	FullName    string `json:"full_name"`
	HtmlURL     string `json:"html_url"`
	Description string
	Owner       *Owner
}

type Owner struct {
	Login string
}
