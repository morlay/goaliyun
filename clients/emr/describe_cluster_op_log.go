package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterOpLogRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterOpLogRequest) Invoke(client goaliyun.Client) (*DescribeClusterOpLogResponse, error) {
	resp := &DescribeClusterOpLogResponse{}
	err := client.Request("emr", "DescribeClusterOpLog", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterOpLogResponse struct {
	RequestId     goaliyun.String
	ChangeLogList DescribeClusterOpLogChangeLogList
}

type DescribeClusterOpLogChangeLog struct {
	Id          goaliyun.Integer
	GmtCreate   goaliyun.String
	GmtModified goaliyun.String
	TargetKey   goaliyun.String
	Status      goaliyun.String
	ChangeType  goaliyun.String
	Message     goaliyun.String
	TargetType  goaliyun.String
}

type DescribeClusterOpLogChangeLogList []DescribeClusterOpLogChangeLog

func (list *DescribeClusterOpLogChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterOpLogChangeLog)
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
