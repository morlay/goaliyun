package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type PunishRequest struct {
	BizPunishRequest string `position:"Query" name:"BizPunishRequest"`
	Sign             string `position:"Query" name:"Sign"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *PunishRequest) Invoke(client goaliyun.Client) (*PunishResponse, error) {
	resp := &PunishResponse{}
	err := client.Request("phoenixsp-inner", "Punish", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PunishResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.String
}
