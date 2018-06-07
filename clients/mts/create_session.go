package mts

import (
	"github.com/morlay/goaliyun"
)

type CreateSessionRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	SessionTime          int64  `position:"Query" name:"SessionTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndUserId            string `position:"Query" name:"EndUserId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateSessionRequest) Invoke(client goaliyun.Client) (*CreateSessionResponse, error) {
	resp := &CreateSessionResponse{}
	err := client.Request("mts", "CreateSession", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSessionResponse struct {
	RequestId goaliyun.String
	Session   goaliyun.String
}
