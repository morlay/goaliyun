package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterHostGroupRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string `position:"Query" name:"HostGroupId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterHostGroupRequest) Invoke(client goaliyun.Client) (*ListClusterHostGroupResponse, error) {
	resp := &ListClusterHostGroupResponse{}
	err := client.Request("emr", "ListClusterHostGroup", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterHostGroupResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	HostList   ListClusterHostGroupHostList
}

type ListClusterHostGroupHost struct {
	HostId         goaliyun.String
	HostName       goaliyun.String
	PublicIp       goaliyun.String
	PrivateIp      goaliyun.String
	Role           goaliyun.String
	InstanceType   goaliyun.String
	Cpu            goaliyun.Integer
	Memory         goaliyun.Integer
	Status         goaliyun.String
	Type           goaliyun.String
	HostInstanceId goaliyun.String
	SerialNumber   goaliyun.String
}

type ListClusterHostGroupHostList []ListClusterHostGroupHost

func (list *ListClusterHostGroupHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostGroupHost)
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
