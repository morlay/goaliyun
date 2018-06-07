package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type GlobalSyncPhoenixCfgRequest struct {
	Data          string `position:"Query" name:"Data"`
	OperationType string `position:"Query" name:"OperationType"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *GlobalSyncPhoenixCfgRequest) Invoke(client goaliyun.Client) (*GlobalSyncPhoenixCfgResponse, error) {
	resp := &GlobalSyncPhoenixCfgResponse{}
	err := client.Request("phoenixsp-inner", "GlobalSyncPhoenixCfg", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GlobalSyncPhoenixCfgResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.String
}
