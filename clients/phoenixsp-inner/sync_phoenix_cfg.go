package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type SyncPhoenixCfgRequest struct {
	Data          string `position:"Query" name:"Data"`
	OperationType string `position:"Query" name:"OperationType"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SyncPhoenixCfgRequest) Invoke(client goaliyun.Client) (*SyncPhoenixCfgResponse, error) {
	resp := &SyncPhoenixCfgResponse{}
	err := client.Request("phoenixsp-inner", "SyncPhoenixCfg", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SyncPhoenixCfgResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.String
}
