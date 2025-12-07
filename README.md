## 介绍
Go语言实现的一个轻量的表达式执行库，适合简单场景，比如判断用户是否有某个角色或权限、是否在白名单等。

## 使用示例

### DEMO 1
判断用户是否有`admin`角色，或者有`normal_user`角色且有`add`权限

```golang
func main() {
	varTable := map[string]interface{}{
		"roles": []interface{}{"normal_user"},
		"perms": []interface{}{"add", "delete"},
	}
	ret := expression.Execute("'admin' in roles or 'normal_user' in roles and 'add' in perms", varTable)
  // 输出 true
	fmt.Println(ret)
}
```

### DEMO 2
判断 用户是否在白名单

```golang
func main() {
	varTable := map[string]interface{}{
		"userId": "uid123456",
	}
	ret := expression.Execute("userId in ['123', '456', 'uid123456']", varTable)
	// 输出 true
	fmt.Println(ret)
}
```

