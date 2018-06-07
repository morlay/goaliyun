package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterHostGroupForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string `position:"Query" name:"HostGroupId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterHostGroupForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterHostGroupForAdminResponse, error) {
	resp := &ListClusterHostGroupForAdminResponse{}
	err := client.Request("emr", "ListClusterHostGroupForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterHostGroupForAdminResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	HostList   ListClusterHostGroupForAdminHostList
}

type ListClusterHostGroupForAdminHost struct {
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

type ListClusterHostGroupForAdminHostList []ListClusterHostGroupForAdminHost

func (list *ListClusterHostGroupForAdminHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHostGroupForAdminHost)
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
