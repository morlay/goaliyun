package ehpc

import (
	"github.com/morlay/goaliyun"
)

type GetAutoScaleConfigRequest struct {
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetAutoScaleConfigRequest) Invoke(client goaliyun.Client) (*GetAutoScaleConfigResponse, error) {
	resp := &GetAutoScaleConfigResponse{}
	err := client.Request("ehpc", "GetAutoScaleConfig", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAutoScaleConfigResponse struct {
	RequestId               goaliyun.String
	Uid                     goaliyun.String
	ClusterId               goaliyun.String
	ClusterType             goaliyun.String
	EnableAutoGrow          bool
	EnableAutoShrink        bool
	GrowIntervalInMinutes   goaliyun.Integer
	ShrinkIntervalInMinutes goaliyun.Integer
	ShrinkIdleTimes         goaliyun.Integer
	GrowTimeoutInMinutes    goaliyun.Integer
	ExtraNodesGrowRatio     goaliyun.Integer
	GrowRatio               goaliyun.Integer
	MaxNodesInCluster       goaliyun.Integer
	ExcludeNodes            goaliyun.String
}
