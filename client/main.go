package client

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Plan(out bool) error {
	cmd := exec.Command("terraform", "plan")
	if out {cmd.Args = append(cmd.Args, "-out=tfplan")}
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

func Apply() error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
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

func InitApply(){
	Init()
	Apply()
}

func InitPlan(out bool){
	Init()
	Plan(out)
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

func PullWorkspaceState(){
	cmd := exec.Command("terraform","state","pull")
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
