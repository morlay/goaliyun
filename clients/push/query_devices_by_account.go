package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDevicesByAccountRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	Account  string `position:"Query" name:"Account"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryDevicesByAccountRequest) Invoke(client goaliyun.Client) (*QueryDevicesByAccountResponse, error) {
	resp := &QueryDevicesByAccountResponse{}
	err := client.Request("push", "QueryDevicesByAccount", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDevicesByAccountResponse struct {
	RequestId goaliyun.String
	DeviceIds QueryDevicesByAccountDeviceIdList
}

type QueryDevicesByAccountDeviceIdList []goaliyun.String

func (list *QueryDevicesByAccountDeviceIdList) UnmarshalJSON(data []byte) error {
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
