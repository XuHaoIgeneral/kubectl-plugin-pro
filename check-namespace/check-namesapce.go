package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mritd/promptx"
)

func main() {

	// 先拿到当前的 context
	cmd := exec.Command("kubectl", "config", "current-context")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()
	checkAndExit(err)
	currentContext := strings.TrimSpace(string(b))

	// 如果提供了 NameSpace 字符串，直接进行
	if len(os.Args) > 1 {
		cmd = exec.Command("kubectl", "config", "set-context", currentContext, "--namespace="+os.Args[1])
		cmd.Stdout = os.Stdout
		checkAndExit(cmd.Run())
		fmt.Printf("Kubernetes namespace switch to %s.\n", os.Args[1])
	} else {
		// 获取所有的namespace
		cmd = exec.Command("kubectl", "get", "ns", "-o", "template", "--template={{range .items }}\n{{.metadata.name}}\n{{end}}")
		b, err = cmd.Output()
		checkAndExit(err)
		allNameSpace := strings.Fields(string(b))
		// 下拉列表
		cfg := &promptx.SelectConfig{
			ActiveTpl:    "» {{.| cyan}}",
			InactiveTpl:  " {{.|white}}",
			SelectPrompt: "NameSpace",
			SelectedTpl:  "{{ \"» \" |green}}{{\"NameSpace:\"|cyan}}{{.}}",
			DisPlaySize:  9,
			DetailsTpl:   ` `,
		}
		s := &promptx.Select{
			Items:  allNameSpace,
			Config: cfg,
		}

		// 用户选中一个 NameSpace 后我就拿到了想要设置的 NameSpace 字符串
		selectNameSpace := allNameSpace[s.Run()]

		// 跟上面套路一样，写进去就行了
		cmd = exec.Command("kubectl", "config", "set-context", currentContext, "--namespace="+selectNameSpace)
		cmd.Stdout = os.Stdout
		checkAndExit(cmd.Run())
		fmt.Printf("Kubernetes namespace switch to %s.\n", selectNameSpace)
	}
}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func checkAndExit(err error) {
	if !checkErr(err) {
		os.Exit(1)
	}
}
