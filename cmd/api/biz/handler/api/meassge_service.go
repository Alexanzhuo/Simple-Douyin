// Code generated by hertz generator.

package api

import (
	"context"

	"Simple-Douyin/cmd/api/biz/handler/pack"
	api "Simple-Douyin/cmd/api/biz/model/api"
	"Simple-Douyin/cmd/api/rpc"
	"Simple-Douyin/kitex_gen/message"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageActionRequest

	err = c.BindAndValidate(&req)
	hlog.Info("token: ", req.Token)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = rpc.SendMessage(context.Background(), &message.MessageActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := new(api.MessageActionResponse)
	resp.StatusCode = 200
	resp.StatusMsg = "send message success"

	c.JSON(consts.StatusOK, resp)
}

// MessageChat .
// @router /douyin/message/chat/ [POST]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	msgs, err := rpc.GetMessageHistory(context.Background(), &message.MessageChatRequest{
		Token:    req.Token,
		ToUserId: req.ToUserID,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := new(api.MessageChatResponse)
	resp.StatusCode = 200
	resp.StatusMsg = "get message history success"
	resp.MessageList = pack.Messages(msgs)

	c.JSON(consts.StatusOK, resp)
}
