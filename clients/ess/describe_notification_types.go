package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNotificationTypesRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeNotificationTypesRequest) Invoke(client goaliyun.Client) (*DescribeNotificationTypesResponse, error) {
	resp := &DescribeNotificationTypesResponse{}
	err := client.Request("ess", "DescribeNotificationTypes", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNotificationTypesResponse struct {
	RequestId         goaliyun.String
	NotificationTypes DescribeNotificationTypesNotificationTypeList
}

type DescribeNotificationTypesNotificationTypeList []goaliyun.String

func (list *DescribeNotificationTypesNotificationTypeList) UnmarshalJSON(data []byte) error {
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
