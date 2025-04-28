package controllers

import (
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/middleware"
	"pleasurelove/internal/usecase"
	"pleasurelove/internal/utils"
	"pleasurelove/internal/utils/errorutils"
	"pleasurelove/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	AuthUsecase usecase.AuthUseCase
}

func NewAuthController(
	authUC usecase.AuthUseCase,
) *AuthController {
	return &AuthController{AuthUsecase: authUC}
}

func (ctrl *AuthController) Logout(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return response.SetResponseBadRequest(c, "Missing Authorization header", nil)
	}

	tokenString := utils.ExtractBearerToken(authHeader)

	// Hapus token dari Redis
	err := middleware.DeleteTokenFromRedis(ctx, tokenString)
	if err != nil {
		logger.Error(ctx, "Failed to delete token from Redis", err)
		return response.SetResponseInternalServerError(c, "Failed to logout", err)
	}

	return response.SetResponseOK(c, "Successfully logged out", nil)
}

func (ctrl *AuthController) ValidateCredentials(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqLogin request.ReqLogin
	if err := c.BodyParser(&reqLogin); err != nil {
		logger.Error(ctx, "Failed to parse login request", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	err := reqLogin.ValidateRequest(ctx)
	if err != nil {
		return response.SetResponseBadRequest(c, "Invalid username or password", err)
	}

	// Validasi kredensial
	user, err := ctrl.AuthUsecase.Login(ctx, &reqLogin)
	if err != nil {
		logger.Error(ctx, "error LoginDashboard", err)
		return response.SetResponseBadRequest(c, "Login Failed, Invalid username or password", err)
	}

	// Generate temporary token
	temporaryToken, err := middleware.GenerateGuestTemporaryToken(user)
	if err != nil {
		logger.Error(ctx, "Failed to generate temporary token", err)
		return response.SetResponseInternalServerError(c, "Failed generate token", err)
	}

	return response.SetResponseOK(c, "Temporary token generated", response.ResAuth{Token: temporaryToken})
}

func (ctrl *AuthController) GenerateAccessToken(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var (
		reqToken request.ReqToken
		err      error
	)
	if err = c.BodyParser(&reqToken); err != nil {
		logger.Error(ctx, "Failed to parse token request", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	err = reqToken.ValidateRequest(ctx)
	if err != nil {
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	// Validasi temporary token
	user, err := middleware.ValidateGuestTemporaryToken(reqToken.TemporaryToken)
	if err != nil {
		logger.Error(ctx, "Invalid temporary token", err)
		return response.SetResponseUnauthorized(c, "Invalid or expired temporary token", err.Error())
	}

	// Generate access token
	user, err = ctrl.AuthUsecase.LoginByUserId(ctx, user.ID)
	if err != nil {
		logger.Error(ctx, "Error GetUserByID", err)
		return response.SetResponseUnauthorized(c, errorutils.ErrMessageDataNotFound, err.Error())
	}

	accessToken, err := middleware.GenerateGuestTokenUserDashboard(user)
	if err != nil {
		logger.Error(ctx, "Failed to generate access token", err)
		return response.SetResponseInternalServerError(c, "Failed to generate access token", err)
	}

	// Simpan access token di Redis
	claims := jwt.MapClaims{}
	_, _, _ = new(jwt.Parser).ParseUnverified(accessToken, claims)
	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	err = middleware.SaveTokenToRedis(ctx, accessToken, exp)
	if err != nil {
		logger.Error(ctx, "Failed to save access token to Redis", err)
		return response.SetResponseInternalServerError(c, "Failed to save access token", err)
	}

	return response.SetResponseOK(c, "Access token generated", response.ResAuth{Token: accessToken})
}
