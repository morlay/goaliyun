package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateNetworkInterfacePermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountId            int64  `position:"Query" name:"AccountId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Permission           string `position:"Query" name:"Permission"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateNetworkInterfacePermissionRequest) Invoke(client goaliyun.Client) (*CreateNetworkInterfacePermissionResponse, error) {
	resp := &CreateNetworkInterfacePermissionResponse{}
	err := client.Request("ecs", "CreateNetworkInterfacePermission", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNetworkInterfacePermissionResponse struct {
	RequestId                  goaliyun.String
	NetworkInterfacePermission CreateNetworkInterfacePermissionNetworkInterfacePermission
}

type CreateNetworkInterfacePermissionNetworkInterfacePermission struct {
	AccountId                    goaliyun.Integer
	ServiceName                  goaliyun.String
	NetworkInterfaceId           goaliyun.String
	NetworkInterfacePermissionId goaliyun.String
	Permission                   goaliyun.String
	PermissionState              goaliyun.String
}
