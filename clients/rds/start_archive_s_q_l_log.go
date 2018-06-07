package rds

import (
	"github.com/morlay/goaliyun"
)

type StartArchiveSQLLogRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Database             string `position:"Query" name:"Database"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	User                 string `position:"Query" name:"User"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StartArchiveSQLLogRequest) Invoke(client goaliyun.Client) (*StartArchiveSQLLogResponse, error) {
	resp := &StartArchiveSQLLogResponse{}
	err := client.Request("rds", "StartArchiveSQLLog", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartArchiveSQLLogResponse struct {
	RequestId goaliyun.String
}
