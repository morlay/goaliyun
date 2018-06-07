package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDevicesByAliasRequest struct {
	Alias    string `position:"Query" name:"Alias"`
	AppKey   int64  `position:"Query" name:"AppKey"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryDevicesByAliasRequest) Invoke(client goaliyun.Client) (*QueryDevicesByAliasResponse, error) {
	resp := &QueryDevicesByAliasResponse{}
	err := client.Request("push", "QueryDevicesByAlias", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDevicesByAliasResponse struct {
	RequestId goaliyun.String
	DeviceIds QueryDevicesByAliasDeviceIdList
}

type QueryDevicesByAliasDeviceIdList []goaliyun.String

func (list *QueryDevicesByAliasDeviceIdList) UnmarshalJSON(data []byte) error {
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
