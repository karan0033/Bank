package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/karan0033/bank/db/sqlc"
	"github.com/karan0033/bank/token"
)

type transferRequest struct {
	FromAccountID int64  `json:"form_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,gt=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusNotFound, errResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != fromAccount.Owner {
		err := errors.New("From Account does not belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}
	_, valid = server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, AccountId int64, Currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, AccountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return account, false
	}

	if account.Currency != Currency {
		err := fmt.Errorf("Account %d currency mismatch: %s vs %s", AccountId, account.Currency, Currency)
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return account, false
	}
	return account, true
}
