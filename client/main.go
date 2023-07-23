package client

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Plan(opts TerraformPlanOpts) error {
	cmd := exec.Command("terraform", "plan")
	if opts.output != "" {cmd.Args = append(cmd.Args, fmt.Sprintf("-out=%s", opts.output))}
	if opts.Vars != nil {cmd.Args = CreateVarArgs(*opts.Vars)}
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func Apply(opts TerraformApplyOpts) error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	if opts.Vars != nil {cmd.Args = CreateVarArgs(*opts.Vars)}
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func Init() error {
	cmd := exec.Command("terraform", "init")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func InitApply(opts TerraformApplyOpts) {
	Init()
	Apply(opts)
}

func InitPlan(opts TerraformPlanOpts) {
	Init()
	Plan(opts)
}

func GetWorkspace() error {
	cmd := exec.Command("terraform", "workspace", "show")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func SelectWorkspace(name string){
	cmd := exec.Command("terraform", "workspace", "select", name)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
	}
}

func CreateWorkspace(name string){
	cmd := exec.Command("terraform", "workspace", "new", name)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
	}
}

func DeleteWorkspace(name string){
	cmd := exec.Command("terraform", "workspace", "delete", name)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
	}
}

func PullWorkspaceState(name string) error{
	cmd := exec.Command("terraform","state","pull",">",name)
	cmd.Start()

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	fmt.Println(cmd.Stdout)
	return nil
}

func PushWorkspaceState(name string){
	cmd := exec.Command("terraform","state","push",name)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
	}
}

func ShowJSON() error {
	cmd := exec.Command("terraform", "show", "-json")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func Format(recursive bool) error {
	cmd := exec.Command("terraform", "fmt")
	if recursive {cmd.Args = append(cmd.Args, "--recursive")}
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := cmd.Wait();if err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}
