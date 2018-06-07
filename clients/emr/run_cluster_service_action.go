package emr

import (
	"github.com/morlay/goaliyun"
)

type RunClusterServiceActionRequest struct {
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	OnlyRestartStaleConfigNodes string `position:"Query" name:"OnlyRestartStaleConfigNodes"`
	NodeCountPerBatch           int64  `position:"Query" name:"NodeCountPerBatch"`
	ClusterId                   string `position:"Query" name:"ClusterId"`
	CustomCommand               string `position:"Query" name:"CustomCommand"`
	ComponentNameList           string `position:"Query" name:"ComponentNameList"`
	ServiceActionName           string `position:"Query" name:"ServiceActionName"`
	IsRolling                   string `position:"Query" name:"IsRolling"`
	TotlerateFailCount          int64  `position:"Query" name:"TotlerateFailCount"`
	ServiceName                 string `position:"Query" name:"ServiceName"`
	Comment                     string `position:"Query" name:"Comment"`
	HostIdList                  string `position:"Query" name:"HostIdList"`
	TurnOnMaintenanceMode       string `position:"Query" name:"TurnOnMaintenanceMode"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *RunClusterServiceActionRequest) Invoke(client goaliyun.Client) (*RunClusterServiceActionResponse, error) {
	resp := &RunClusterServiceActionResponse{}
	err := client.Request("emr", "RunClusterServiceAction", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RunClusterServiceActionResponse struct {
	RequestId goaliyun.String
}
