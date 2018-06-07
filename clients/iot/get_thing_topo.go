package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetThingTopoRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	PageNo     int64  `position:"Query" name:"PageNo"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetThingTopoRequest) Invoke(client goaliyun.Client) (*GetThingTopoResponse, error) {
	resp := &GetThingTopoResponse{}
	err := client.Request("iot", "GetThingTopo", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetThingTopoResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         GetThingTopoData
}

type GetThingTopoData struct {
	Total       goaliyun.Integer
	CurrentPage goaliyun.Integer
	PageSize    goaliyun.Integer
	PageCount   goaliyun.Integer
	List        GetThingTopoDeviceInfoList
}

type GetThingTopoDeviceInfo struct {
	IotId      goaliyun.String
	ProductKey goaliyun.String
	DeviceName goaliyun.String
}

type GetThingTopoDeviceInfoList []GetThingTopoDeviceInfo

func (list *GetThingTopoDeviceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThingTopoDeviceInfo)
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
