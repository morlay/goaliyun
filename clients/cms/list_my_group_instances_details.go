package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMyGroupInstancesDetailsRequest struct {
	Total       string `position:"Query" name:"Total"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
	GroupId     int64  `position:"Query" name:"GroupId"`
	PageSize    int64  `position:"Query" name:"PageSize"`
	Category    string `position:"Query" name:"Category"`
	Keyword     string `position:"Query" name:"Keyword"`
	PageNumber  int64  `position:"Query" name:"PageNumber"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListMyGroupInstancesDetailsRequest) Invoke(client goaliyun.Client) (*ListMyGroupInstancesDetailsResponse, error) {
	resp := &ListMyGroupInstancesDetailsResponse{}
	err := client.Request("cms", "ListMyGroupInstancesDetails", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMyGroupInstancesDetailsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	Total        goaliyun.Integer
	Resources    ListMyGroupInstancesDetailsResourceList
}

type ListMyGroupInstancesDetailsResource struct {
	AliUid       goaliyun.Integer
	InstanceName goaliyun.String
	InstanceId   goaliyun.String
	Desc         goaliyun.String
	NetworkType  goaliyun.String
	Category     goaliyun.String
	Tags         ListMyGroupInstancesDetailsTagList
	Region       ListMyGroupInstancesDetailsRegion
	Vpc          ListMyGroupInstancesDetailsVpc
}

type ListMyGroupInstancesDetailsTag struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type ListMyGroupInstancesDetailsRegion struct {
	RegionId         goaliyun.String
	AvailabilityZone goaliyun.String
}

type ListMyGroupInstancesDetailsVpc struct {
	VpcInstanceId     goaliyun.String
	VswitchInstanceId goaliyun.String
}

type ListMyGroupInstancesDetailsResourceList []ListMyGroupInstancesDetailsResource

func (list *ListMyGroupInstancesDetailsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupInstancesDetailsResource)
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

type ListMyGroupInstancesDetailsTagList []ListMyGroupInstancesDetailsTag

func (list *ListMyGroupInstancesDetailsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupInstancesDetailsTag)
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
