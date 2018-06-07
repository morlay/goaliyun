package nas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMountTargetsRequest struct {
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	PageSize          int64  `position:"Query" name:"PageSize"`
	PageNumber        int64  `position:"Query" name:"PageNumber"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *DescribeMountTargetsRequest) Invoke(client goaliyun.Client) (*DescribeMountTargetsResponse, error) {
	resp := &DescribeMountTargetsResponse{}
	err := client.Request("nas", "DescribeMountTargets", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMountTargetsResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageSize     goaliyun.Integer
	PageNumber   goaliyun.Integer
	MountTargets DescribeMountTargetsMountTargetList
}

type DescribeMountTargetsMountTarget struct {
	MountTargetDomain goaliyun.String
	NetworkType       goaliyun.String
	VpcId             goaliyun.String
	VswId             goaliyun.String
	AccessGroup       goaliyun.String
	Status            goaliyun.String
}

type DescribeMountTargetsMountTargetList []DescribeMountTargetsMountTarget

func (list *DescribeMountTargetsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMountTargetsMountTarget)
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
