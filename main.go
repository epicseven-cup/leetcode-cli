package main

import (
	"fmt"
	"github.com/epicseven-cup/leetcode-cli/internal/graphql"
	"github.com/epicseven-cup/leetcode-cli/internal/utils"
	"os"
)

var templates = map[string]string{
	"python": "https://github.com/epicseven-cup/Leetcode-template-python.git",
	"go":     "https://github.com/epicseven-cup/leetcode-go-template.git",
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: leetcode_cli <command> <args>")
		os.Exit(1)
	}

	if args[0] == "daily" && len(args) == 2 {
		lc := graphql.LeetcodeClient{}
		daily, err := lc.GetDaily()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		repo, ok := templates[args[1]]

		if !ok {
			fmt.Println("Template not found:", args[1])
			os.Exit(1)
		}

		err = utils.CreateDailyFolder(daily, repo, args[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch args[1] {
		case "python":
			err := utils.TemplateCode(daily, "python", "solution.py", os.O_WRONLY|os.O_CREATE)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "go":
			err := utils.TemplateCode(daily, "golang", "./problem/solution.go", os.O_APPEND|os.O_WRONLY)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		os.Exit(0)

	}

}
