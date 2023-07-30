## Terraform cli client For GO
This is a lightweight wrapper around the terraform cli and makes it easy to run terraform commands from within a go application. It is intended to be used in a CI/CD pipeline to run terraform commands and parse the output.
### Usage
```go
package main
import (
    "fmt"
    "github.com/eunanhardy/terraform-client/client"
)

func main() {
    options := client.TerraformPlanOpts{
		Format: "json",
	}
    client.Init()
    client.Plan(options) //runs a terraform plan and returns the output as json. Format value is optional
}
```