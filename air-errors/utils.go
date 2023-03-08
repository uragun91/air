package airerrors

import "air-api/models"

func GetErrorResponse(errorCode int, error error) models.ErrorResponse {
	return models.ErrorResponse{
		ErrorCode: errorCode,
		Error: error.Error(),
	};
}