package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryBatchRegisterDeviceStatusRequest struct {
	ApplyId    int64  `position:"Query" name:"ApplyId"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryBatchRegisterDeviceStatusRequest) Invoke(client goaliyun.Client) (*QueryBatchRegisterDeviceStatusResponse, error) {
	resp := &QueryBatchRegisterDeviceStatusResponse{}
	err := client.Request("iot", "QueryBatchRegisterDeviceStatus", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryBatchRegisterDeviceStatusResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryBatchRegisterDeviceStatusData
}

type QueryBatchRegisterDeviceStatusData struct {
	Status      goaliyun.String
	ValidList   QueryBatchRegisterDeviceStatusValidListList
	InvalidList QueryBatchRegisterDeviceStatusInvalidListList
}

type QueryBatchRegisterDeviceStatusValidListList []goaliyun.String

func (list *QueryBatchRegisterDeviceStatusValidListList) UnmarshalJSON(data []byte) error {
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

type QueryBatchRegisterDeviceStatusInvalidListList []goaliyun.String

func (list *QueryBatchRegisterDeviceStatusInvalidListList) UnmarshalJSON(data []byte) error {
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
