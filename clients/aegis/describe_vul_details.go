package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVulDetailsRequest struct {
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

func (req *DescribeVulDetailsRequest) Invoke(client goaliyun.Client) (*DescribeVulDetailsResponse, error) {
	resp := &DescribeVulDetailsResponse{}
	err := client.Request("aegis", "DescribeVulDetails", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVulDetailsResponse struct {
	RequestId      goaliyun.String
	Name           goaliyun.String
	AliasName      goaliyun.String
	Level          goaliyun.String
	VulPublishTs   goaliyun.Integer
	Type           goaliyun.String
	Product        goaliyun.String
	HasPatch       bool
	PatchPublishTs goaliyun.Integer
	PatchSource    goaliyun.String
	Cvss           goaliyun.String
	CveIds         goaliyun.String
	Advice         goaliyun.String
	Description    goaliyun.String
	PendingCount   goaliyun.Integer
	HandledCount   goaliyun.Integer
	CveLists       DescribeVulDetailsCveListList
}

type DescribeVulDetailsCveListList []goaliyun.String

func (list *DescribeVulDetailsCveListList) UnmarshalJSON(data []byte) error {
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
