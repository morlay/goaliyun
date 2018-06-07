package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNetworkInterfacePermissionsRequest struct {
	ResourceOwnerId               int64                                                                `position:"Query" name:"ResourceOwnerId"`
	PageNumber                    int64                                                                `position:"Query" name:"PageNumber"`
	PageSize                      int64                                                                `position:"Query" name:"PageSize"`
	NetworkInterfacePermissionIds *DescribeNetworkInterfacePermissionsNetworkInterfacePermissionIdList `position:"Query" type:"Repeated" name:"NetworkInterfacePermissionId"`
	ResourceOwnerAccount          string                                                               `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string                                                               `position:"Query" name:"OwnerAccount"`
	OwnerId                       int64                                                                `position:"Query" name:"OwnerId"`
	NetworkInterfaceId            string                                                               `position:"Query" name:"NetworkInterfaceId"`
	RegionId                      string                                                               `position:"Query" name:"RegionId"`
}

func (req *DescribeNetworkInterfacePermissionsRequest) Invoke(client goaliyun.Client) (*DescribeNetworkInterfacePermissionsResponse, error) {
	resp := &DescribeNetworkInterfacePermissionsResponse{}
	err := client.Request("ecs", "DescribeNetworkInterfacePermissions", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNetworkInterfacePermissionsResponse struct {
	RequestId                   goaliyun.String
	TotalCount                  goaliyun.Integer
	PageNumber                  goaliyun.Integer
	PageSize                    goaliyun.Integer
	NetworkInterfacePermissions DescribeNetworkInterfacePermissionsNetworkInterfacePermissionList
}

type DescribeNetworkInterfacePermissionsNetworkInterfacePermission struct {
	AccountId                    goaliyun.Integer
	ServiceName                  goaliyun.String
	NetworkInterfaceId           goaliyun.String
	NetworkInterfacePermissionId goaliyun.String
	Permission                   goaliyun.String
	PermissionState              goaliyun.String
}

type DescribeNetworkInterfacePermissionsNetworkInterfacePermissionIdList []string

func (list *DescribeNetworkInterfacePermissionsNetworkInterfacePermissionIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacePermissionsNetworkInterfacePermissionList []DescribeNetworkInterfacePermissionsNetworkInterfacePermission

func (list *DescribeNetworkInterfacePermissionsNetworkInterfacePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacePermissionsNetworkInterfacePermission)
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
