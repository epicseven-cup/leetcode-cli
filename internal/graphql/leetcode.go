package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Leetcode struct {
	//TODO: internal information that will need to be storage
}

type questionOfTodayV2 struct {
	ActiveDailyCodingChallengeQuestion ActiveDailyCodingChallengeQuestion `json:"activeDailyCodingChallengeQuestion"`
}

type Data struct {
	Query string `json:"query"`
}

func (lc *Leetcode) GetDaily() (*DailyCodingChallenge, error) {
	query := NewQuery("questionOfTodayV2")

	activeDailyCodingChallengeQuestionQuery := NewQuery("activeDailyCodingChallengeQuestion")
	activeDailyCodingChallengeQuestionQuery.AddField("date")
	activeDailyCodingChallengeQuestionQuery.AddField("userStatus")
	activeDailyCodingChallengeQuestionQuery.AddField("link")

	question := NewQuery("question")
	question.AddField("id: questionId")
	question.AddField("title")
	question.AddField("content")
	question.AddField("questionFrontendId")
	question.AddField("difficulty")

	codeSnippets := NewQuery("codeSnippets")
	codeSnippets.AddField("lang")
	codeSnippets.AddField("langSlug")
	codeSnippets.AddField("code")

	// starts to nest
	query.AddQuery(activeDailyCodingChallengeQuestionQuery)
	activeDailyCodingChallengeQuestionQuery.AddQuery(question)
	question.AddQuery(codeSnippets)

	q := query.Build()
	data := &Data{
		Query: "query " + q,
	}
	fmt.Println(q)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, "https://leetcode.com/graphql/", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:139.0) Gecko/20100101 Firefox/139.0")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	dailyQuestion := DailyCodingChallenge{}
	err = json.Unmarshal(responseData, &dailyQuestion)
	if err != nil {
		return nil, err
	}

	return &dailyQuestion, nil
}
