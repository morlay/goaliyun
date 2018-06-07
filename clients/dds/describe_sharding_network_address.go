package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeShardingNetworkAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeShardingNetworkAddressRequest) Invoke(client goaliyun.Client) (*DescribeShardingNetworkAddressResponse, error) {
	resp := &DescribeShardingNetworkAddressResponse{}
	err := client.Request("dds", "DescribeShardingNetworkAddress", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeShardingNetworkAddressResponse struct {
	RequestId        goaliyun.String
	NetworkAddresses DescribeShardingNetworkAddressNetworkAddressList
}

type DescribeShardingNetworkAddressNetworkAddress struct {
	NetworkAddress goaliyun.String
	IPAddress      goaliyun.String
	NetworkType    goaliyun.String
	Port           goaliyun.String
	VPCId          goaliyun.String
	VswitchId      goaliyun.String
	NodeId         goaliyun.String
	ExpiredTime    goaliyun.String
}

type DescribeShardingNetworkAddressNetworkAddressList []DescribeShardingNetworkAddressNetworkAddress

func (list *DescribeShardingNetworkAddressNetworkAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeShardingNetworkAddressNetworkAddress)
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
