package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/Bengkel/dto"
	"github.com/IrvanWijayaSardam/Bengkel/entity"
	"github.com/IrvanWijayaSardam/Bengkel/helper"
	"github.com/IrvanWijayaSardam/Bengkel/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionContoller interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Delete(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

// All implements TransactionContoller
func (c *transactionController) All(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	trx := c.transactionService.All(userID)
	res := helper.BuildResponse(true, "OK!", trx)
	context.JSON(http.StatusOK, res)
}

// Delete implements TransactionContoller
func (c *transactionController) Delete(context *gin.Context) {
	var transaction entity.Transaction
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	transaction.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.transactionService.IsAllowedToEdit(userID, transaction.ID) {
		c.transactionService.Delete(transaction)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

// Insert implements TransactionContoller
func (c *transactionController) Insert(context *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			transactionCreateDTO.UserID = convertedUserID
		}
		result := c.transactionService.InsertTransaction(transactionCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func NewTransactionController(trxServ service.TransactionService, jwtServ service.JWTService) TransactionContoller {
	return &transactionController{
		transactionService: trxServ,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userid"])
	return id
}
