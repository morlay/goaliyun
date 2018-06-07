package emr

import (
	"github.com/morlay/goaliyun"
)

type MigrateInstanceForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MigrateInstanceForAdminRequest) Invoke(client goaliyun.Client) (*MigrateInstanceForAdminResponse, error) {
	resp := &MigrateInstanceForAdminResponse{}
	err := client.Request("emr", "MigrateInstanceForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MigrateInstanceForAdminResponse struct {
	RequestId goaliyun.String
	WfId      goaliyun.String
}
