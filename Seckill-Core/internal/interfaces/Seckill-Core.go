package interfaces

import (
	"context"
	"fmt"

	pb "api/Seckill-Core/v1"

	"Seckill-Core/internal/response"

	"github.com/BitofferHub/pkg/constant"
	"github.com/BitofferHub/pkg/middlewares/log"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserInfo(c *gin.Context) {
	traceID := c.Request.Header.Get(constant.TraceID)

	var req pb.GetUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("CreateShortUrlV1 err: %v", err)
		response.Fail(c, response.ShouldBindError, nil)
		return
	}

	ctx := context.WithValue(context.Background(), constant.TraceID, traceID)
	resp, err := h.Seckill - CoreService.GetUser(ctx, &req)
	if err != nil {
		fmt.Println("get user err", err)
		response.Fail(c, response.GetUserError, nil)
		return
	}

	response.Success(c, resp)
}

func (h *Handler) CreateUserInfo(ctx *gin.Context) {
	traceID := ctx.Request.Header.Get(constant.TraceID)
	var req pb.CreateUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		log.Errorf("CreateShortUrlV1 err: %v", err)
		response.Fail(ctx, response.ShouldBindError, nil)
		return
	}

	c := context.WithValue(context.Background(), constant.TraceID, traceID)
	_, err := h.Seckill - CoreService.CreateUser(c, &req)
	if err != nil {
		log.Errorf("CreateShortUrlV1 err: %v", err)
		response.Fail(ctx, response.CreateUserErr, nil)
		return

	}
	response.Success(ctx, nil)
	return
}
