package models

type ErrorResponse struct {
	ErrorCode	int `json:"errorCode"`
	Error 		any `json:"error"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type RegistrationInput struct {
	Email 		string `json:"email" binding:"required,isvalidemail"`
	Password 	string `json:"password" binding:"required"`
};
