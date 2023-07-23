package main

import (
	"github.com/eunanhardy/terraform-client/client"
)


func main() {
	err := client.PullWorkspaceState("backup.state.tfstate");if err != nil {
		panic(err)
	}
}