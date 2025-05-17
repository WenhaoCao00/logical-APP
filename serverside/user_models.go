package main

type User struct {
	UUID            string         `json:"uuid"`
	Name            string         `json:"name"`
	Username        string         `json:"username"`
	Email           string         `json:"email"`
	Password        string         `json:"password,omitempty"`
	AvatarUrl       string         `json:"avatarUrl,omitempty"`
	WebsiteUrls     []string       `json:"websiteUrls,omitempty"`
	Bio             string         `json:"bio,omitempty"`
	Description     string         `json:"description,omitempty"`
	TotalQuizNumber int            `json:"totalQuizNumber,omitempty"`
	QuizScores      map[string]int `json:"quizScores,omitempty"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
