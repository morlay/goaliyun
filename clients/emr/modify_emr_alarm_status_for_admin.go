package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyEmrAlarmStatusForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	UniqueKey       string `position:"Query" name:"UniqueKey"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyEmrAlarmStatusForAdminRequest) Invoke(client goaliyun.Client) (*ModifyEmrAlarmStatusForAdminResponse, error) {
	resp := &ModifyEmrAlarmStatusForAdminResponse{}
	err := client.Request("emr", "ModifyEmrAlarmStatusForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyEmrAlarmStatusForAdminResponse struct {
	RequestId goaliyun.String
	Success   bool
}
