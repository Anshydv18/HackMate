package dto

type User struct {
	Name          string      `json:"name"`
	College       string      `json:"college"`
	TechStacks    []TechStack `json:"tech_stacks"`
	Phone         string      `json:"phone_number"`
	Email         string      `json:"email"`
	Age           int         `json:"age"`
	GithubLink    string      `json:"github_link"`
	PortfolioLink string      `json:"portfolio_link"`
	ProfilePhoto  string      `json:"profile_photo"`
}
