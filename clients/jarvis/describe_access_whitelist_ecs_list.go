package jarvis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessWhitelistEcsListRequest struct {
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	SourceCode string `position:"Query" name:"SourceCode"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessWhitelistEcsListRequest) Invoke(client goaliyun.Client) (*DescribeAccessWhitelistEcsListResponse, error) {
	resp := &DescribeAccessWhitelistEcsListResponse{}
	err := client.Request("jarvis", "DescribeAccessWhitelistEcsList", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessWhitelistEcsListResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	Module     goaliyun.String
	EcsList    DescribeAccessWhitelistEcsListEcsList
}

type DescribeAccessWhitelistEcsListEcs struct {
	InstanceName goaliyun.String
	InstanceId   goaliyun.String
	IP           goaliyun.String
}

type DescribeAccessWhitelistEcsListEcsList []DescribeAccessWhitelistEcsListEcs

func (list *DescribeAccessWhitelistEcsListEcsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessWhitelistEcsListEcs)
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
