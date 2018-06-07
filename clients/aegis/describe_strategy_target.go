package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeStrategyTargetRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Type            string `position:"Query" name:"Type"`
	Config          string `position:"Query" name:"Config"`
	Target          string `position:"Query" name:"Target"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeStrategyTargetRequest) Invoke(client goaliyun.Client) (*DescribeStrategyTargetResponse, error) {
	resp := &DescribeStrategyTargetResponse{}
	err := client.Request("aegis", "DescribeStrategyTarget", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStrategyTargetResponse struct {
	RequestId       goaliyun.String
	Count           goaliyun.Integer
	StrategyTargets DescribeStrategyTargetStringItemList
}

type DescribeStrategyTargetStringItem struct {
	Flag       goaliyun.String
	Target     goaliyun.String
	TargetType goaliyun.String
}

type DescribeStrategyTargetStringItemList []DescribeStrategyTargetStringItem

func (list *DescribeStrategyTargetStringItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStrategyTargetStringItem)
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
