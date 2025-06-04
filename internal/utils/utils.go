package utils

import (
	"github.com/epicseven-cup/leetcode-cli/internal/graphql"
	"os"
)

func PythonTemplate(lc *graphql.DailyCodingChallenge) error {
	for _, c := range lc.Data.ActiveDailyCodingChallengeQuestion.Question.CodeSnippets {
		if c.LangSlug == "python" {
			err := os.WriteFile("solution.py", []byte(c.Code), 0677)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
