package utils

import (
	"fmt"
	"github.com/epicseven-cup/leetcode-cli/internal/graphql"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CreateDailyFolder(daily *graphql.DailyCodingChallenge, gitHubTemplate string, lang string) error {
	problemFolder := fmt.Sprintf("./[%s]-[%s]-[%s]-[%s]",
		time.Now().Format(time.DateOnly),

		daily.Data.ActiveDailyCodingChallengeQuestion.Question.QuestionFrontendID,

		strings.ReplaceAll(daily.Data.ActiveDailyCodingChallengeQuestion.Question.Title,
			" ", "_"),
		lang,
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

func TemplateCode(lc *graphql.DailyCodingChallenge, langSlug string, filePath string, options int) error {
	file, err := os.OpenFile(filePath, options, 0666)
	if err != nil {
		return err
	}
	for _, c := range lc.Data.ActiveDailyCodingChallengeQuestion.Question.CodeSnippets {
		if c.LangSlug == langSlug {
			_, err := file.Write([]byte(c.Code))
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
