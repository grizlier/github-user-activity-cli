package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
	"encoding/json"
)

type action struct {
	Type			string	`json:"type"`
	Repo			repo		`json:"repo"`
	Payload		payl 		`json:"payload"`
}

type repo struct {
	Name			string	`json:"name"`
}

type payl struct {
	Action		string	`json:"action"`
	Ref				string	`json:"ref"`
	RefT			string	`json:"ref_type"`	
	Release 	release	`json:"release"	`
}

type release struct {
	Name			string	`json:"name"`
}

var acts [10]action

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Insert nickname")
		return
	}

	nick := os.Args[1]
	req := "https://api.github.com/users/<username>/events"
	reqN := strings.Replace(req, "<username>", nick, 1)

	resp, err := http.Get(reqN)
	if err != nil {
		fmt.Println("Request error: ", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), &acts)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	} 

	fmt.Println("Last 10 activities: ")
	for _, act := range acts {
		switch form := act.Type; form {
			case "PushEvent":
				fmt.Println("- Pushed to", act.Repo.Name)
			
			case "PullRequestEvent":
				if act.Payload.Action == "closed" {
					fmt.Println("- Closed pull request in", act.Repo.Name)
				}	else if act.Payload.Action == "opened" { 
					fmt.Println("- Opened pull request in", act.Repo.Name)
				}	else {
					fmt.Println("Unentered pull request event: ", act.Payload.Action)
				}
			
			case "WatchEvent":
				fmt.Printf("- Watched %s repo\n", act.Repo.Name)

			case "IssuesEvent":
				if act.Payload.Action == "closed" {
					fmt.Println("- Closed issue in", act.Repo.Name)
				} else if act.Payload.Action == "created" {
					fmt.Println("- Created issue in", act.Repo.Name)
				} else {
					fmt.Println("Unentered issue event: ", act.Payload.Action)
				}

			case "IssueCommentEvent":
				if act.Payload.Action == "created" {
					fmt.Println("- Commented issue in", act.Repo.Name)
				} else {
					fmt.Println("Unentered issue comment event: ", act.Payload.Action)
				}
			
			case "DeleteEvent":
				fmt.Printf("- Deleted %s: \"%s\" in %s \n", act.Payload.RefT, act.Payload.Ref, act.Repo.Name)

			case "ReleaseEvent":
				if act.Payload.Action == "published" {
					fmt.Printf("- Released: %s in %s\n", act.Payload.Release.Name, act.Repo.Name)
				}	else {
					fmt.Println("Unentered release event: ", act.Payload.Action)
				}
			case "CreateEvent":
				fmt.Printf("- Created %s: \"%s\" in %s \n", act.Payload.RefT, act.Payload.Ref, act.Repo.Name)

			default:
				fmt.Println("- Unentered action: ", form)
		}
	}
}
