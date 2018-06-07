package iot

import (
	"github.com/morlay/goaliyun"
)

type QueryProductRequest struct {
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryProductRequest) Invoke(client goaliyun.Client) (*QueryProductResponse, error) {
	resp := &QueryProductResponse{}
	err := client.Request("iot", "QueryProduct", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryProductResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryProductData
}

type QueryProductData struct {
	GmtCreate   goaliyun.String
	DataFormat  goaliyun.Integer
	Description goaliyun.String
	DeviceCount goaliyun.Integer
	NodeType    goaliyun.Integer
	ProductKey  goaliyun.String
	ProductName goaliyun.String
}
