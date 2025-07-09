package utils

import (
	"github.com/epicseven-cup/leetcode-cli/internal/graphql"
	"os"
	"fmt"
	"strings"
	"os/exec"
	"time"
)

func CreateDailyFolder(daily *graphql.DailyCodingChallenge, gitHubTemplate string) error {
	problemFolder := fmt.Sprintf("./[%s]-[%s]-[%s]",
		time.Now().Format(time.DateOnly),

		daily.Data.ActiveDailyCodingChallengeQuestion.Question.QuestionFrontendID,

		strings.ReplaceAll(daily.Data.ActiveDailyCodingChallengeQuestion.Question.Title,
		" ", "_"),

	)

	err := exec.Command("git", "clone", gitHubTemplate, problemFolder).Run()

	if err != nil {
		return err
	}

	// Change Dir into created folder
	err = os.Chdir(problemFolder)

	if err != nil {
		return err
	}

	var ps string
	ps += fmt.Sprintf("<p>Date: %s</p>\r\n", daily.Data.ActiveDailyCodingChallengeQuestion.Date)
	ps += fmt.Sprintf("<p>Daily Question: %s </p>\r\n", daily.Data.ActiveDailyCodingChallengeQuestion.Question.Title)
	ps += fmt.Sprintf("<a href='https://leetcode.com%s'>Link: https://leetcode.com%s </a>\n", daily.Data.ActiveDailyCodingChallengeQuestion.Link, daily.Data.ActiveDailyCodingChallengeQuestion.Link)
	ps += fmt.Sprintf("<p>Difficulty: %s</p>\r\n", daily.Data.ActiveDailyCodingChallengeQuestion.Question.Difficulty)
	ps += fmt.Sprintf("<p>Content:%s</p>\r\n", daily.Data.ActiveDailyCodingChallengeQuestion.Question.Content)

	err = os.WriteFile("Problem.md", []byte(ps), 0644)
	if err != nil {
		return err
	}
	return nil
}

func TemplateCode(lc *graphql.DailyCodingChallenge, langSlug string, filePath string) error {
	for _, c := range lc.Data.ActiveDailyCodingChallengeQuestion.Question.CodeSnippets {
		if c.LangSlug == langSlug {
			err := os.WriteFile(filePath, []byte(c.Code), 0677)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}


