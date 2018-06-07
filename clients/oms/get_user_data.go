package oms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetUserDataRequest struct {
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	OwnerAccount string `position:"Query" name:"OwnerAccount"`
	ProductName  string `position:"Query" name:"ProductName"`
	DataType     string `position:"Query" name:"DataType"`
	StartTime    string `position:"Query" name:"StartTime"`
	EndTime      string `position:"Query" name:"EndTime"`
	NextToken    string `position:"Query" name:"NextToken"`
	MaxResult    int64  `position:"Query" name:"MaxResult"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *GetUserDataRequest) Invoke(client goaliyun.Client) (*GetUserDataResponse, error) {
	resp := &GetUserDataResponse{}
	err := client.Request("oms", "GetUserData", "2015-02-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserDataResponse struct {
	RequestId   goaliyun.String
	ProductName goaliyun.String
	DataType    goaliyun.String
	NextToken   goaliyun.String
	DataList    GetUserDataDataList
}

type GetUserDataData struct {
	DataItems GetUserDataDataItemList
}

type GetUserDataDataItem struct {
	Name  goaliyun.String
	Value goaliyun.String
}

type GetUserDataDataList []GetUserDataData

func (list *GetUserDataDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserDataData)
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

type GetUserDataDataItemList []GetUserDataDataItem

func (list *GetUserDataDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserDataDataItem)
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
