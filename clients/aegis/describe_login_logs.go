package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLoginLogsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	CurrentPage     int64  `position:"Query" name:"CurrentPage"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeLoginLogsRequest) Invoke(client goaliyun.Client) (*DescribeLoginLogsResponse, error) {
	resp := &DescribeLoginLogsResponse{}
	err := client.Request("aegis", "DescribeLoginLogs", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLoginLogsResponse struct {
	RequestId      goaliyun.String
	PageSize       goaliyun.Integer
	CurrentPage    goaliyun.Integer
	TotalCount     goaliyun.Integer
	HttpStatusCode goaliyun.Integer
	LoginLogs      DescribeLoginLogsLoginLogList
}

type DescribeLoginLogsLoginLogList []goaliyun.String

func (list *DescribeLoginLogsLoginLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
