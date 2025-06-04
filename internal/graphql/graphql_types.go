package graphql

type Date string
type Link string
type UserStatus string

// DailyCodingChallenge represents the overall structure of the API response
type DailyCodingChallenge struct {
	Data struct {
		ActiveDailyCodingChallengeQuestion ActiveDailyCodingChallengeQuestion `json:"activeDailyCodingChallengeQuestion"`
	} `json:"data"`
}

type ActiveDailyCodingChallengeQuestion struct {
	Date       Date       `json:"date"`
	UserStatus UserStatus `json:"userStatus"`
	Link       Link       `json:"link"`
	Question   Question   `json:"question"`
}

type TopicTag struct {
	Name           string  `json:"name"`
	Slug           string  `json:"slug"`
	NameTranslated *string `json:"nameTranslated"` // Use pointer for nullability
}

type CodeSnippets struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}
type Question struct {
	ID                 string         `json:"id"`
	TitleSlug          string         `json:"titleSlug"`
	Title              string         `json:"title"`
	TranslatedTitle    *string        `json:"translatedTitle"` // Use pointer for nullability
	QuestionFrontendID string         `json:"questionFrontendId"`
	PaidOnly           bool           `json:"paidOnly"`
	Content            string         `json:"content"`
	Difficulty         string         `json:"difficulty"`
	TopicTags          []TopicTag     `json:"topicTags"`
	CodeSnippets       []CodeSnippets `json:"codeSnippets"`
	Status             *string        `json:"status"` // Use pointer for nullability
	IsInMyFavorites    bool           `json:"isInMyFavorites"`
	AcRate             float64        `json:"acRate"`
	Frequency          *string        `json:"frequency"` // Use pointer for nullability
}
