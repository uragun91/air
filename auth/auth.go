package auth

import (
	airerrors "air-api/air-errors"
	"air-api/database"
	"air-api/models"
	"air-api/roles"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var input models.RegistrationInput;

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, airerrors.GetErrorResponse(airerrors.UnexpectedError, err));
		return;
	}

	query := `
		SELECT id FROM users
		WHERE email = $1;
	`;
	var user_id string;
	row := database.DB.QueryRow(query, input.Email);
	err := row.Scan(&user_id);

	if (err == nil) {
		context.JSON(http.StatusBadRequest, airerrors.GetErrorResponse(airerrors.UserAlreadyExists, errors.New("user already exists")));
		return;
	}

	if (err != nil && err != sql.ErrNoRows) {
		context.JSON(http.StatusInternalServerError, airerrors.GetErrorResponse(airerrors.UnexpectedError, err));
		return;
	}

	transaction, _ := database.DB.Begin();

	pwd, err := HashPassword(input.Password);
	if err != nil {
		context.JSON(http.StatusInternalServerError, airerrors.GetErrorResponse(airerrors.UnexpectedError, err));
		log.Print(err);
		return;
	}

	userQuery := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`
	var id string;

	err = database.DB.QueryRow(userQuery, input.Email, pwd).Scan(&id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, airerrors.GetErrorResponse(airerrors.UnexpectedError, err));
		log.Print(err);
		transaction.Rollback();
		return;
	}

	userRoleQuery := `
		INSERT INTO roles_users (role_id, user_id)
		(
			SELECT roles.id, $1 as user_id
			FROM roles
			WHERE roles.name = $2
		)
	`
	_, err = database.DB.Exec(userRoleQuery, id, roles.RoleCustomer);
	if err != nil {
		context.JSON(http.StatusInternalServerError, airerrors.GetErrorResponse(airerrors.UnexpectedError, err));
		log.Print(err);
		transaction.Rollback();
		return;
	}

	newJwt, err := GenerateToken(id, roles.RoleCustomer);
	if err != nil {
		context.JSON(http.StatusInternalServerError, airerrors.GetErrorResponse(airerrors.UnexpectedError, err));
		transaction.Rollback();
		log.Print(err);
		return;
	}

	transaction.Commit();

	context.JSON(http.StatusOK, models.LoginResponse {AccessToken: newJwt});
}