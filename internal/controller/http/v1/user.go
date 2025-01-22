package v1

import (
	"net/http"
	"payment-service/internal/service"

	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	userService service.User
}

func newAccountRoutes(g *echo.Group, userService service.User) {
	r := &userRoutes{
		userService: userService,
	}

	g.POST("/deposit", r.deposit)
}

type userDepositInput struct {
	UUID   string `json:"uuid" validate:"required"`
	Amount uint64 `json:"amount" validate:"required"`
}

func (r *userRoutes) deposit(c echo.Context) error {
	var input userDepositInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := r.userService.Deposit(c.Request().Context(), service.DepositInput{
		UUID:   input.UUID,
		Amount: input.Amount,
	})

	if err != nil {
		if err == service.ErrUserNotFound {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type userWithdrawInput struct {
	UUID   string `json:"uuid" validate:"required"`
	Amount uint64 `json:"amount" validate:"required"`
}

func (r *userRoutes) withdraw(c echo.Context) error {
	var input userWithdrawInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := r.userService.Withdraw(c.Request().Context(), service.WithdrawInput{
		UUID:   input.UUID,
		Amount: input.Amount,
	})
	if err != nil {
		if err == service.ErrUserNotFound {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type UserTransferInput struct {
	FromUUID string `json:"fromuuid" validate:"required"`
	ToUUID   string `json:"touuid" validate:"required"`
	Amount   uint64 `json:"amount" validate:"required"`
}

func (r *userRoutes) transfer(c echo.Context) error {
	var input UserTransferInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := r.userService.Transfer(c.Request().Context(), service.TransferInput{
		FromUUID: input.FromUUID,
		ToUUID:   input.ToUUID,
		Amount:   input.Amount,
	})
	if err != nil {
		if err == service.ErrUserNotFound {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

type getBalanceInput struct {
	UUID string `json:"id" validate:"required"`
}

func (r *userRoutes) getBalance(c echo.Context) error {
	var input getBalanceInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	balance, err := r.userService.GetBalanceByUUID(c.Request().Context(), input.UUID)
	if err != nil {
		if err == service.ErrUserNotFound {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	type response struct {
		UUID    string `json:"uuid"`
		Balance uint64 `json:"balance"`
	}

	return c.JSON(http.StatusOK, response{
		Balance: balance,
	})
}
