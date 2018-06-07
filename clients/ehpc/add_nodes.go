package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddNodesRequest struct {
	AutoRenewPeriod       int64  `position:"Query" name:"AutoRenewPeriod"`
	Period                int64  `position:"Query" name:"Period"`
	ImageId               string `position:"Query" name:"ImageId"`
	Count                 int64  `position:"Query" name:"Count"`
	ClusterId             string `position:"Query" name:"ClusterId"`
	ComputeSpotStrategy   string `position:"Query" name:"ComputeSpotStrategy"`
	ImageOwnerAlias       string `position:"Query" name:"ImageOwnerAlias"`
	PeriodUnit            string `position:"Query" name:"PeriodUnit"`
	AutoRenew             string `position:"Query" name:"AutoRenew"`
	EcsChargeType         string `position:"Query" name:"EcsChargeType"`
	ComputeSpotPriceLimit string `position:"Query" name:"ComputeSpotPriceLimit"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *AddNodesRequest) Invoke(client goaliyun.Client) (*AddNodesResponse, error) {
	resp := &AddNodesResponse{}
	err := client.Request("ehpc", "AddNodes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddNodesResponse struct {
	RequestId   goaliyun.String
	InstanceIds AddNodesInstanceIdList
}

type AddNodesInstanceIdList []goaliyun.String

func (list *AddNodesInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
