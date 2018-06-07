package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVulListRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Product         string `position:"Query" name:"Product"`
	StatusList      string `position:"Query" name:"StatusList"`
	Level           string `position:"Query" name:"Level"`
	Resource        string `position:"Query" name:"Resource"`
	OrderBy         string `position:"Query" name:"OrderBy"`
	Dealed          string `position:"Query" name:"Dealed"`
	CurrentPage     int64  `position:"Query" name:"CurrentPage"`
	Type            string `position:"Query" name:"Type"`
	LastTsEnd       int64  `position:"Query" name:"LastTsEnd"`
	BatchName       string `position:"Query" name:"BatchName"`
	PatchId         int64  `position:"Query" name:"PatchId"`
	AliasName       string `position:"Query" name:"AliasName"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Name            string `position:"Query" name:"Name"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	Lang            string `position:"Query" name:"Lang"`
	LastTsStart     int64  `position:"Query" name:"LastTsStart"`
	Necessity       string `position:"Query" name:"Necessity"`
	Uuids           string `position:"Query" name:"Uuids"`
	Direction       string `position:"Query" name:"Direction"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeVulListRequest) Invoke(client goaliyun.Client) (*DescribeVulListResponse, error) {
	resp := &DescribeVulListResponse{}
	err := client.Request("aegis", "DescribeVulList", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVulListResponse struct {
	RequestId   goaliyun.String
	Count       goaliyun.Integer
	PageSize    goaliyun.Integer
	CurrentPage goaliyun.Integer
	TotalCount  goaliyun.Integer
	VulRecords  DescribeVulListVulRecordList
}

type DescribeVulListVulRecordList []goaliyun.String

func (list *DescribeVulListVulRecordList) UnmarshalJSON(data []byte) error {
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
