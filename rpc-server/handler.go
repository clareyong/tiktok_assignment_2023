package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"strings"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	resp := rpc.NewSendResponse()
	log.Info(fmt.Sprintf("req.Message.Chat is '%s'", req.Message.Chat))
	if strings.Count(req.Message.Chat, ":") != 1 {
		resp.Code, resp.Msg = 1, "wrong format for chat"
		return resp, nil
	}

	chatParticipants := strings.Split(req.Message.Chat, ":")
	if len(chatParticipants) != 2 {
		panic("this shouldn't happen")
	}

	var sender, receiver string
	if chatParticipants[0] == req.Message.Sender {
		sender, receiver = chatParticipants[0], chatParticipants[1]
	} else if chatParticipants[1] == req.Message.Sender {
		receiver, sender = chatParticipants[0], chatParticipants[1]
	} else {
		resp.Code, resp.Msg = 2, "sender name must be in the chat"
		return resp, nil
	}

	if chatParticipants[0] == chatParticipants[1] {
		resp.Code, resp.Msg = 3, "chat name must contain different participants"
		return resp, nil
	}

	newMessage := Message{Sender: sender, Receiver: receiver, Text: req.Message.Text}

	result := defaultDB.Create(newMessage)
	if result.Error != nil {
		log.Error(result.Error)
		resp.Code, resp.Msg = 4, "error saving message"
		return resp, nil
	}

	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	resp := rpc.NewPullResponse()
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}
