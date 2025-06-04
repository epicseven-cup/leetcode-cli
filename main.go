package main

import (
	"fmt"
	"github.com/epicseven-cup/leetcode-cli/internal/graphql"
	"github.com/epicseven-cup/leetcode-cli/internal/utils"
	"os"
	"os/exec"
	"strings"
)

var templates = map[string]string{
	"python": "https://github.com/epicseven-cup/Leetcode-template-python.git",
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: leetcode_cli <command> <args>")
		os.Exit(1)
	}

	if args[0] == "daily" && len(args) == 2 {
		lc := graphql.Leetcode{}
		daily, err := lc.GetDaily()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tp, ok := templates[args[1]]
		if !ok {
			fmt.Println("Template not found:", args[1])
			os.Exit(1)
		}

		index := strings.LastIndex(tp, "/")
		folderName := tp[index : len(tp)-4]

		err = exec.Command("git", "clone", tp).Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		problemFolder := fmt.Sprintf("./%s-%s", daily.Data.ActiveDailyCodingChallengeQuestion.Question.QuestionFrontendID, strings.Replace(daily.Data.ActiveDailyCodingChallengeQuestion.Question.Title, " ", "_", -1))

		err = exec.Command("mv", fmt.Sprintf("./%s/", folderName), problemFolder).Run()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = os.Chdir(problemFolder)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ps := ""
		ps += fmt.Sprintf("<p>Date: %s</p>", daily.Data.ActiveDailyCodingChallengeQuestion.Date)
		ps += fmt.Sprintf("<p>Daily Question: %s </p>", daily.Data.ActiveDailyCodingChallengeQuestion.Question.Title)
		ps += fmt.Sprintf("<a href='https://leetcode.com%s'>Link: https://leetcode.com%s </a>", daily.Data.ActiveDailyCodingChallengeQuestion.Link, daily.Data.ActiveDailyCodingChallengeQuestion.Link)
		ps += fmt.Sprintf("<p>Difficulty: %s</p>", daily.Data.ActiveDailyCodingChallengeQuestion.Question.Difficulty)
		ps += fmt.Sprintf("<p>Content:</p>\n%s", daily.Data.ActiveDailyCodingChallengeQuestion.Question.Content)

		err = os.WriteFile("problem.md", []byte(ps), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch args[1] {

		case "python":
			err := utils.PythonTemplate(daily)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		os.Exit(0)

	}

}
