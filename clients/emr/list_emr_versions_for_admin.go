package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListEmrVersionsForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListEmrVersionsForAdminRequest) Invoke(client goaliyun.Client) (*ListEmrVersionsForAdminResponse, error) {
	resp := &ListEmrVersionsForAdminResponse{}
	err := client.Request("emr", "ListEmrVersionsForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListEmrVersionsForAdminResponse struct {
	RequestId               goaliyun.String
	PageNumber              goaliyun.Integer
	PageSize                goaliyun.Integer
	TotalCount              goaliyun.Integer
	MainVersionKeyValueList ListEmrVersionsForAdminMainVersionKeyValueList
}

type ListEmrVersionsForAdminMainVersionKeyValue struct {
	Id                   goaliyun.Integer
	Key                  goaliyun.String
	Value                goaliyun.String
	Status               goaliyun.Integer
	UtcCreateTimestamp   goaliyun.Integer
	UtcModifiedTimestamp goaliyun.Integer
}

type ListEmrVersionsForAdminMainVersionKeyValueList []ListEmrVersionsForAdminMainVersionKeyValue

func (list *ListEmrVersionsForAdminMainVersionKeyValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmrVersionsForAdminMainVersionKeyValue)
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
