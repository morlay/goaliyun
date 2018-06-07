package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeCustinsKernelReleaseNotesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCustinsKernelReleaseNotesRequest) Invoke(client goaliyun.Client) (*DescribeCustinsKernelReleaseNotesResponse, error) {
	resp := &DescribeCustinsKernelReleaseNotesResponse{}
	err := client.Request("rds", "DescribeCustinsKernelReleaseNotes", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCustinsKernelReleaseNotesResponse struct {
	RequestId                 goaliyun.String
	DBInstanceId              goaliyun.String
	DBInstanceDiffReleaseNote goaliyun.String
}
