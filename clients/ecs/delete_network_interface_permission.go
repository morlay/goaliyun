package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteNetworkInterfacePermissionRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	NetworkInterfacePermissionId string `position:"Query" name:"NetworkInterfacePermissionId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	Force                        string `position:"Query" name:"Force"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *DeleteNetworkInterfacePermissionRequest) Invoke(client goaliyun.Client) (*DeleteNetworkInterfacePermissionResponse, error) {
	resp := &DeleteNetworkInterfacePermissionResponse{}
	err := client.Request("ecs", "DeleteNetworkInterfacePermission", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNetworkInterfacePermissionResponse struct {
	RequestId goaliyun.String
}
