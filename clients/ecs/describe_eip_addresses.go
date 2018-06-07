package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeEipAddressesRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Filter2Value           string `position:"Query" name:"Filter.2.Value"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AllocationId           string `position:"Query" name:"AllocationId"`
	Filter1Value           string `position:"Query" name:"Filter.1.Value"`
	Filter2Key             string `position:"Query" name:"Filter.2.Key"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	EipAddress             string `position:"Query" name:"EipAddress"`
	PageNumber             int64  `position:"Query" name:"PageNumber"`
	LockReason             string `position:"Query" name:"LockReason"`
	Filter1Key             string `position:"Query" name:"Filter.1.Key"`
	AssociatedInstanceType string `position:"Query" name:"AssociatedInstanceType"`
	PageSize               int64  `position:"Query" name:"PageSize"`
	ChargeType             string `position:"Query" name:"ChargeType"`
	AssociatedInstanceId   string `position:"Query" name:"AssociatedInstanceId"`
	Status                 string `position:"Query" name:"Status"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *DescribeEipAddressesRequest) Invoke(client goaliyun.Client) (*DescribeEipAddressesResponse, error) {
	resp := &DescribeEipAddressesResponse{}
	err := client.Request("ecs", "DescribeEipAddresses", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeEipAddressesResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	EipAddresses DescribeEipAddressesEipAddressList
}

type DescribeEipAddressesEipAddress struct {
	RegionId           goaliyun.String
	IpAddress          goaliyun.String
	AllocationId       goaliyun.String
	Status             goaliyun.String
	InstanceId         goaliyun.String
	Bandwidth          goaliyun.String
	EipBandwidth       goaliyun.String
	InternetChargeType goaliyun.String
	AllocationTime     goaliyun.String
	InstanceType       goaliyun.String
	ChargeType         goaliyun.String
	ExpiredTime        goaliyun.String
	OperationLocks     DescribeEipAddressesLockReasonList
}

type DescribeEipAddressesLockReason struct {
	LockReason goaliyun.String
}

type DescribeEipAddressesEipAddressList []DescribeEipAddressesEipAddress

func (list *DescribeEipAddressesEipAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipAddressesEipAddress)
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

type DescribeEipAddressesLockReasonList []DescribeEipAddressesLockReason

func (list *DescribeEipAddressesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipAddressesLockReason)
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
