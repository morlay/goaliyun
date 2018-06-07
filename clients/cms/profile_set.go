package cms

import (
	"github.com/morlay/goaliyun"
)

type ProfileSetRequest struct {
	EnableInstallAgentNewECS string `position:"Query" name:"EnableInstallAgentNewECS"`
	EnableActiveAlert        string `position:"Query" name:"EnableActiveAlert"`
	AutoInstall              string `position:"Query" name:"AutoInstall"`
	UserId                   int64  `position:"Query" name:"UserId"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *ProfileSetRequest) Invoke(client goaliyun.Client) (*ProfileSetResponse, error) {
	resp := &ProfileSetResponse{}
	err := client.Request("cms", "ProfileSet", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ProfileSetResponse struct {
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Success      bool
	RequestId    goaliyun.String
}
