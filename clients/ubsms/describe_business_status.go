package ubsms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBusinessStatusRequest struct {
	CallerBid string `position:"Query" name:"CallerBid"`
	Password  string `position:"Query" name:"Password"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DescribeBusinessStatusRequest) Invoke(client goaliyun.Client) (*DescribeBusinessStatusResponse, error) {
	resp := &DescribeBusinessStatusResponse{}
	err := client.Request("ubsms", "DescribeBusinessStatus", "2015-06-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBusinessStatusResponse struct {
	RequestId              goaliyun.String
	Success                bool
	UserBusinessStatusList DescribeBusinessStatusUserBusinessStatusList
}

type DescribeBusinessStatusUserBusinessStatus struct {
	Uid         goaliyun.String
	ServiceCode goaliyun.String
	Statuses    DescribeBusinessStatusStatusList
}

type DescribeBusinessStatusStatus struct {
	StatusKey   goaliyun.String
	StatusValue goaliyun.String
}

type DescribeBusinessStatusUserBusinessStatusList []DescribeBusinessStatusUserBusinessStatus

func (list *DescribeBusinessStatusUserBusinessStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBusinessStatusUserBusinessStatus)
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

type DescribeBusinessStatusStatusList []DescribeBusinessStatusStatus

func (list *DescribeBusinessStatusStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBusinessStatusStatus)
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
