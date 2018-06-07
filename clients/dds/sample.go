package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SampleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SampleRequest) Invoke(client goaliyun.Client) (*SampleResponse, error) {
	resp := &SampleResponse{}
	err := client.Request("dds", "Sample", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SampleResponse struct {
	RequestId        goaliyun.String
	SecurityIps      goaliyun.String
	SecurityIpGroups SampleSecurityIpGroupList
}

type SampleSecurityIpGroup struct {
	SecurityIpGroupName      goaliyun.String
	SecurityIpGroupAttribute goaliyun.String
	SecurityIpList           goaliyun.String
}

type SampleSecurityIpGroupList []SampleSecurityIpGroup

func (list *SampleSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SampleSecurityIpGroup)
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
