package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterForInternalRequest struct {
	ResourceOwnerId int64                                    `position:"Query" name:"ResourceOwnerId"`
	ClusterIdLists  *ListClusterForInternalClusterIdListList `position:"Query" type:"Repeated" name:"ClusterIdList"`
	UserId          string                                   `position:"Query" name:"UserId"`
	RegionId        string                                   `position:"Query" name:"RegionId"`
}

func (req *ListClusterForInternalRequest) Invoke(client goaliyun.Client) (*ListClusterForInternalResponse, error) {
	resp := &ListClusterForInternalResponse{}
	err := client.Request("emr", "ListClusterForInternal", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterForInternalResponse struct {
	RequestId                   goaliyun.String
	DescribeClusterResponseList ListClusterForInternalDescribeClusterResponseList
}

type ListClusterForInternalDescribeClusterResponse struct {
	RequestId   goaliyun.String
	Id          goaliyun.String
	BizId       goaliyun.String
	Name        goaliyun.String
	StartTime   goaliyun.Integer
	StopTime    goaliyun.Integer
	LogEnable   bool
	LogPath     goaliyun.String
	UserId      goaliyun.String
	RunningTime goaliyun.Integer
	Status      goaliyun.String
	ExpiredTime goaliyun.Integer
	FailReason  ListClusterForInternalFailReason
}

type ListClusterForInternalFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type ListClusterForInternalClusterIdListList []string

func (list *ListClusterForInternalClusterIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListClusterForInternalDescribeClusterResponseList []ListClusterForInternalDescribeClusterResponse

func (list *ListClusterForInternalDescribeClusterResponseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterForInternalDescribeClusterResponse)
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
