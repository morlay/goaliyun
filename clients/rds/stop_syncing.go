package rds

import (
	"github.com/morlay/goaliyun"
)

type StopSyncingRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int64  `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StopSyncingRequest) Invoke(client goaliyun.Client) (*StopSyncingResponse, error) {
	resp := &StopSyncingResponse{}
	err := client.Request("rds", "StopSyncing", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopSyncingResponse struct {
	RequestId goaliyun.String
}
