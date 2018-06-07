package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AssignPrivateIpAddressesRequest struct {
	ResourceOwnerId                int64                                         `position:"Query" name:"ResourceOwnerId"`
	SecondaryPrivateIpAddressCount int64                                         `position:"Query" name:"SecondaryPrivateIpAddressCount"`
	ResourceOwnerAccount           string                                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string                                        `position:"Query" name:"OwnerAccount"`
	OwnerId                        int64                                         `position:"Query" name:"OwnerId"`
	PrivateIpAddresss              *AssignPrivateIpAddressesPrivateIpAddressList `position:"Query" type:"Repeated" name:"PrivateIpAddress"`
	NetworkInterfaceId             string                                        `position:"Query" name:"NetworkInterfaceId"`
	RegionId                       string                                        `position:"Query" name:"RegionId"`
}

func (req *AssignPrivateIpAddressesRequest) Invoke(client goaliyun.Client) (*AssignPrivateIpAddressesResponse, error) {
	resp := &AssignPrivateIpAddressesResponse{}
	err := client.Request("ecs", "AssignPrivateIpAddresses", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssignPrivateIpAddressesResponse struct {
	RequestId goaliyun.String
}

type AssignPrivateIpAddressesPrivateIpAddressList []string

func (list *AssignPrivateIpAddressesPrivateIpAddressList) UnmarshalJSON(data []byte) error {
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
