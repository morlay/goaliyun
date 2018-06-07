package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLoadBalancersRequest struct {
	Access_key_id         string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	Address               string `position:"Query" name:"Address"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	NetworkType           string `position:"Query" name:"NetworkType"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ServerId              string `position:"Query" name:"ServerId"`
	MasterZoneId          string `position:"Query" name:"MasterZoneId"`
	PageNumber            int64  `position:"Query" name:"PageNumber"`
	Tags                  string `position:"Query" name:"Tags"`
	ServerIntranetAddress string `position:"Query" name:"ServerIntranetAddress"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	LoadBalancerName      string `position:"Query" name:"LoadBalancerName"`
	LoadBalancerId        string `position:"Query" name:"LoadBalancerId"`
	InternetChargeType    string `position:"Query" name:"InternetChargeType"`
	VpcId                 string `position:"Query" name:"VpcId"`
	PageSize              int64  `position:"Query" name:"PageSize"`
	AddressType           string `position:"Query" name:"AddressType"`
	SlaveZoneId           string `position:"Query" name:"SlaveZoneId"`
	PayType               string `position:"Query" name:"PayType"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *DescribeLoadBalancersRequest) Invoke(client goaliyun.Client) (*DescribeLoadBalancersResponse, error) {
	resp := &DescribeLoadBalancersResponse{}
	err := client.Request("slb", "DescribeLoadBalancers", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLoadBalancersResponse struct {
	RequestId     goaliyun.String
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	TotalCount    goaliyun.Integer
	LoadBalancers DescribeLoadBalancersLoadBalancerList
}

type DescribeLoadBalancersLoadBalancer struct {
	LoadBalancerId     goaliyun.String
	LoadBalancerName   goaliyun.String
	LoadBalancerStatus goaliyun.String
	Address            goaliyun.String
	AddressType        goaliyun.String
	RegionId           goaliyun.String
	RegionIdAlias      goaliyun.String
	VSwitchId          goaliyun.String
	VpcId              goaliyun.String
	NetworkType        goaliyun.String
	MasterZoneId       goaliyun.String
	SlaveZoneId        goaliyun.String
	InternetChargeType goaliyun.String
	CreateTime         goaliyun.String
	CreateTimeStamp    goaliyun.Integer
	PayType            goaliyun.String
	ResourceGroupId    goaliyun.String
}

type DescribeLoadBalancersLoadBalancerList []DescribeLoadBalancersLoadBalancer

func (list *DescribeLoadBalancersLoadBalancerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersLoadBalancer)
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
