package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeUserCustomLogConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserCustomLogConfigRequest) Invoke(client goaliyun.Client) (*DescribeUserCustomLogConfigResponse, error) {
	resp := &DescribeUserCustomLogConfigResponse{}
	err := client.Request("cdn", "DescribeUserCustomLogConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserCustomLogConfigResponse struct {
	RequestId goaliyun.String
	ConfigIds DescribeUserCustomLogConfigConfigIdList
}

type DescribeUserCustomLogConfigConfigIdList []goaliyun.String

func (list *DescribeUserCustomLogConfigConfigIdList) UnmarshalJSON(data []byte) error {
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
