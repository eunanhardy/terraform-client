package client

import (
	"fmt"
)

func CreateVarArgs(vars []string) []string{
	var_list := []string{}
	for _, v := range vars {
		var_list = append(var_list,fmt.Sprintf("-var=%s",v))
	}

	return var_list
}