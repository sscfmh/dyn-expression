package main

import (
	"fmt"

	expression "github.com/sscfmh/dyn-expression"
)

func main() {
	varTable := map[string]interface{}{
		"userId": "uid123456",
	}
	ret := expression.Execute("userId in ['123', '456', 'uid123456']", varTable)
	// 输出 true
	fmt.Println(ret)
}

func main1() {
	varTable := map[string]interface{}{
		"roles": []interface{}{"normal_user"},
		"perms": []interface{}{"add", "delete"},
	}
	ret := expression.Execute("'admin' in roles or 'normal_user' in roles and 'add' in perms", varTable)
	fmt.Println(ret)
}
