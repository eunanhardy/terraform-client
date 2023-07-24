package client

import (
	"fmt"
	"log"
)

func Plan(opts TerraformPlanOpts) error {
	cmd := []string{"terraform", "plan"}
	if opts.Output != "" {cmd = append(cmd, fmt.Sprintf("-out=%s", opts.Output))}
	if opts.Vars != nil {cmd = CreateVarArgs(*opts.Vars)}
	fmt.Println(cmd)
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func Apply(opts TerraformApplyOpts) error {
	cmd := []string{"terraform", "apply", "-auto-approve"}
	if opts.Vars != nil {cmd = append(CreateVarArgs(*opts.Vars))}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func Init() error {
	cmd := []string{"terraform", "init"}
	err := RunCommand(cmd...);if err != nil {
		return err
	}
	return nil
}

func InitApply(opts TerraformApplyOpts) {
	Init()
	Apply(opts)
}

func InitPlan(opts TerraformPlanOpts) {
	err := Init();if err != nil {
		log.Print(err)
	}
	go Plan(opts)
}

func GetWorkspace() (string,error) {
	cmd := []string{"terraform", "workspace","show"}
	stdout ,err := RunCommandWithOutput(cmd...);if err != nil {
		log.Print(err)
		return "",err
	}

	return stdout, nil
}

func SelectWorkspace(name string){
	cmd := []string{"terraform", "workspace","select",name}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
	}
}

func CreateWorkspace(name string){
	cmd := []string{"terraform", "workspace","new",name}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
	}
}

func DeleteWorkspace(name string){
	cmd := []string{"terraform", "workspace","delete",name}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
	}
}

func PullWorkspaceState(name string) error{
	cmd := []string{"terraform", "state","pull",">",name}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func PushWorkspaceState(name string){
	cmd := []string{"terraform", "state","push",name}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
	}

}

func ShowJSON() (string,error) {
	cmd := []string{"terraform", "show","-json"}
	stdout,err := RunCommandWithOutput(cmd...);if err != nil {
		log.Print(err)
		return "",err
	}

	return stdout,nil
}

func Format(recursive bool) error {
	cmd := []string{"terraform", "fmt"}
	if recursive {cmd = append(cmd, "--recursive")}
	err := RunCommand(cmd...);if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
