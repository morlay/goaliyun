package cms

import (
	"github.com/morlay/goaliyun"
)

type ProfileGetRequest struct {
	UserId   int64  `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ProfileGetRequest) Invoke(client goaliyun.Client) (*ProfileGetResponse, error) {
	resp := &ProfileGetResponse{}
	err := client.Request("cms", "ProfileGet", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ProfileGetResponse struct {
	ErrorCode                goaliyun.Integer
	ErrorMessage             goaliyun.String
	Success                  bool
	RequestId                goaliyun.String
	UserId                   goaliyun.Integer
	AutoInstall              bool
	EnableInstallAgentNewECS bool
	EnableActiveAlert        goaliyun.String
}
