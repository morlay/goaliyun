package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAuditFilesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAuditFilesRequest) Invoke(client goaliyun.Client) (*DescribeAuditFilesResponse, error) {
	resp := &DescribeAuditFilesResponse{}
	err := client.Request("dds", "DescribeAuditFiles", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAuditFilesResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	DBInstanceId     goaliyun.String
	Items            DescribeAuditFilesLogFileList
}

type DescribeAuditFilesLogFile struct {
	FileID         goaliyun.Integer
	LogStatus      goaliyun.String
	LogStartTime   goaliyun.String
	LogEndTime     goaliyun.String
	LogDownloadURL goaliyun.String
	LogSize        goaliyun.Integer
}

type DescribeAuditFilesLogFileList []DescribeAuditFilesLogFile

func (list *DescribeAuditFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuditFilesLogFile)
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
