package rds

import (
	"github.com/morlay/goaliyun"
)

type SwitchDBInstanceHARequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
	Operation            string `position:"Query" name:"Operation"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SwitchDBInstanceHARequest) Invoke(client goaliyun.Client) (*SwitchDBInstanceHAResponse, error) {
	resp := &SwitchDBInstanceHAResponse{}
	err := client.Request("rds", "SwitchDBInstanceHA", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SwitchDBInstanceHAResponse struct {
	RequestId goaliyun.String
}
