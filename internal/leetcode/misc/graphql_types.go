package misc

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

type Question struct {
	ID                 string  `json:"id"`
	TitleSlug          string  `json:"titleSlug"`
	Title              string  `json:"title"`
	TranslatedTitle    *string `json:"translatedTitle"` // Use pointer for nullability
	QuestionFrontendID string  `json:"questionFrontendId"`
	PaidOnly           bool    `json:"paidOnly"`
	Difficulty         string  `json:"difficulty"`
	TopicTags          []struct {
		Name           string  `json:"name"`
		Slug           string  `json:"slug"`
		NameTranslated *string `json:"nameTranslated"` // Use pointer for nullability
	} `json:"topicTags"`
	Status          *string `json:"status"` // Use pointer for nullability
	IsInMyFavorites bool    `json:"isInMyFavorites"`
	AcRate          float64 `json:"acRate"`
	Frequency       *string `json:"frequency"` // Use pointer for nullability
}
