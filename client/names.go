package client

type TerraformApplyOpts struct {
	Vars *[]string
	Format string
}

type TerraformPlanOpts struct {
	Vars   *[]string
	Format string
	Output string
}
