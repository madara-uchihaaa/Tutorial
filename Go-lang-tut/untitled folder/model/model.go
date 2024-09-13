package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Chat struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Image    string `json:"image"`
	UserID   string `json:"userid"`
}

type SignUpInput struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

func (user *SignUpInput) Validate() bool {
	return user.Email != "" && user.Password != "" && user.ConfirmPassword != "" && user.Password == user.ConfirmPassword
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type QueryInput struct {
	Query string `json:"query"`
}

func (user *QueryInput) Validate() bool {
	return user.Query != ""
}

func (user *SignInInput) Validate() bool {
	return user.Email != "" && user.Password != ""
}
