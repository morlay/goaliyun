package rds

import (
	"github.com/morlay/goaliyun"
)

type CancelImportRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int64  `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelImportRequest) Invoke(client goaliyun.Client) (*CancelImportResponse, error) {
	resp := &CancelImportResponse{}
	err := client.Request("rds", "CancelImport", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelImportResponse struct {
	RequestId goaliyun.String
}
