package ehpc

import (
	"github.com/morlay/goaliyun"
)

type SetAutoScaleConfigRequest struct {
	ShrinkIdleTimes         int64  `position:"Query" name:"ShrinkIdleTimes"`
	GrowTimeoutInMinutes    int64  `position:"Query" name:"GrowTimeoutInMinutes"`
	ClusterId               string `position:"Query" name:"ClusterId"`
	EnableAutoGrow          string `position:"Query" name:"EnableAutoGrow"`
	EnableAutoShrink        string `position:"Query" name:"EnableAutoShrink"`
	MaxNodesInCluster       int64  `position:"Query" name:"MaxNodesInCluster"`
	ExcludeNodes            string `position:"Query" name:"ExcludeNodes"`
	ShrinkIntervalInMinutes int64  `position:"Query" name:"ShrinkIntervalInMinutes"`
	ExtraNodesGrowRatio     int64  `position:"Query" name:"ExtraNodesGrowRatio"`
	GrowIntervalInMinutes   int64  `position:"Query" name:"GrowIntervalInMinutes"`
	GrowRatio               int64  `position:"Query" name:"GrowRatio"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *SetAutoScaleConfigRequest) Invoke(client goaliyun.Client) (*SetAutoScaleConfigResponse, error) {
	resp := &SetAutoScaleConfigResponse{}
	err := client.Request("ehpc", "SetAutoScaleConfig", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetAutoScaleConfigResponse struct {
	RequestId goaliyun.String
}
