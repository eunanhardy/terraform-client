package client

type TerraformApplyOpts struct {
	Vars *[]string
}

type TerraformPlanOpts struct {
	Vars   *[]string
	output string
}