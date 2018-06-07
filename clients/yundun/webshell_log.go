package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type WebshellLogRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RecordType int64  `position:"Query" name:"RecordType"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *WebshellLogRequest) Invoke(client goaliyun.Client) (*WebshellLogResponse, error) {
	resp := &WebshellLogResponse{}
	err := client.Request("yundun", "WebshellLog", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type WebshellLogResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	LogList    WebshellLogWebshellLogList
}

type WebshellLogWebshellLog struct {
	Id     goaliyun.String
	Path   goaliyun.String
	Status goaliyun.Integer
	Time   goaliyun.String
}

type WebshellLogWebshellLogList []WebshellLogWebshellLog

func (list *WebshellLogWebshellLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WebshellLogWebshellLog)
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
