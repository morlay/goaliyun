package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type RemovePunishRequest struct {
	BizPunishRequest string `position:"Query" name:"BizPunishRequest"`
	Sign             string `position:"Query" name:"Sign"`
	PunishResult     string `position:"Query" name:"PunishResult"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *RemovePunishRequest) Invoke(client goaliyun.Client) (*RemovePunishResponse, error) {
	resp := &RemovePunishResponse{}
	err := client.Request("phoenixsp-inner", "RemovePunish", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemovePunishResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.String
}
