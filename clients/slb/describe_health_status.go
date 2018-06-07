package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeHealthStatusRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeHealthStatusRequest) Invoke(client goaliyun.Client) (*DescribeHealthStatusResponse, error) {
	resp := &DescribeHealthStatusResponse{}
	err := client.Request("slb", "DescribeHealthStatus", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeHealthStatusResponse struct {
	RequestId      goaliyun.String
	BackendServers DescribeHealthStatusBackendServerList
}

type DescribeHealthStatusBackendServer struct {
	ListenerPort       goaliyun.Integer
	ServerId           goaliyun.String
	Port               goaliyun.Integer
	ServerHealthStatus goaliyun.String
	ServerIp           goaliyun.String
	Type               goaliyun.String
}

type DescribeHealthStatusBackendServerList []DescribeHealthStatusBackendServer

func (list *DescribeHealthStatusBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHealthStatusBackendServer)
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
