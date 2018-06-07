package oms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetProductDefineRequest struct {
	ProductName string `position:"Query" name:"ProductName"`
	DataType    string `position:"Query" name:"DataType"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetProductDefineRequest) Invoke(client goaliyun.Client) (*GetProductDefineResponse, error) {
	resp := &GetProductDefineResponse{}
	err := client.Request("oms", "GetProductDefine", "2015-02-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetProductDefineResponse struct {
	RequestId   goaliyun.String
	ProductName goaliyun.String
	DataType    goaliyun.String
	ProductList GetProductDefineProductList
}

type GetProductDefineProduct struct {
	ProductName goaliyun.String
	TypeList    GetProductDefineTypeList
}

type GetProductDefineType struct {
	DataType goaliyun.String
	Fields   GetProductDefineFieldList
}

type GetProductDefineProductList []GetProductDefineProduct

func (list *GetProductDefineProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProductDefineProduct)
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

type GetProductDefineTypeList []GetProductDefineType

func (list *GetProductDefineTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProductDefineType)
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

type GetProductDefineFieldList []goaliyun.String

func (list *GetProductDefineFieldList) UnmarshalJSON(data []byte) error {
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
