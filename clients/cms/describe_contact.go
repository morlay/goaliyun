package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeContactRequest struct {
	ContactName string `position:"Query" name:"ContactName"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DescribeContactRequest) Invoke(client goaliyun.Client) (*DescribeContactResponse, error) {
	resp := &DescribeContactResponse{}
	err := client.Request("cms", "DescribeContact", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeContactResponse struct {
	RequestId  goaliyun.String
	Success    bool
	Code       goaliyun.Integer
	Message    goaliyun.String
	Datapoints DescribeContactDatapoints
}

type DescribeContactDatapoints struct {
	Name     goaliyun.String
	Channels DescribeContactChannelList
}

type DescribeContactChannel struct {
	Type  goaliyun.String
	Value goaliyun.String
}

type DescribeContactChannelList []DescribeContactChannel

func (list *DescribeContactChannelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeContactChannel)
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
