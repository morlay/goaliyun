package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UnassignPrivateIpAddressesRequest struct {
	ResourceOwnerId      int64                                           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                          `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                           `position:"Query" name:"OwnerId"`
	PrivateIpAddresss    *UnassignPrivateIpAddressesPrivateIpAddressList `position:"Query" type:"Repeated" name:"PrivateIpAddress"`
	NetworkInterfaceId   string                                          `position:"Query" name:"NetworkInterfaceId"`
	RegionId             string                                          `position:"Query" name:"RegionId"`
}

func (req *UnassignPrivateIpAddressesRequest) Invoke(client goaliyun.Client) (*UnassignPrivateIpAddressesResponse, error) {
	resp := &UnassignPrivateIpAddressesResponse{}
	err := client.Request("ecs", "UnassignPrivateIpAddresses", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnassignPrivateIpAddressesResponse struct {
	RequestId goaliyun.String
}

type UnassignPrivateIpAddressesPrivateIpAddressList []string

func (list *UnassignPrivateIpAddressesPrivateIpAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
