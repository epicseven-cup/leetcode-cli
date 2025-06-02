package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/epicseven-cup/leetcode-cli/internal/leetcode/misc"
	"io"
	"log"
	"net/http"
)

type Leetcode struct {
	//TODO: internal information that will need to be storage
}

type questionOfTodayV2 struct {
	ActiveDailyCodingChallengeQuestion ActiveDailyCodingChallengeQuestion `json:"activeDailyCodingChallengeQuestion"`
}

type ActiveDailyCodingChallengeQuestion struct {
}

//type Query struct {
//	questionOfTodayV2 string `json:"questionOfTodayV2"`
//}

type Data struct {
	Query string `json:"query"`
}

func (lc *Leetcode) GetDaily() {
	data := Data{Query: "\nquery questionOfTodayV2 {\n  activeDailyCodingChallengeQuestion {\n    date\n    userStatus\n    link\n    question {\n      id: questionId\n      titleSlug\n      title\n      translatedTitle\n      questionFrontendId\n      paidOnly: isPaidOnly\n      difficulty\n      topicTags {\n        name\n        slug\n        nameTranslated: translatedName\n      }\n      status\n      isInMyFavorites: isFavor\n      acRate\n      frequency: freqBar\n    }\n  }\n}\n    "}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	request, err := http.NewRequest(http.MethodPost, "https://leetcode.com/graphql/", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:139.0) Gecko/20100101 Firefox/139.0")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(responseData))

	dailyQuestion := misc.DailyCodingChallenge{}
	err = json.Unmarshal(responseData, &dailyQuestion)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(dailyQuestion.Data.ActiveDailyCodingChallengeQuestion.Link)

}
