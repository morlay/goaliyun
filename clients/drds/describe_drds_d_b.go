package drds

import (
	"github.com/morlay/goaliyun"
)

type DescribeDrdsDBRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDrdsDBRequest) Invoke(client goaliyun.Client) (*DescribeDrdsDBResponse, error) {
	resp := &DescribeDrdsDBResponse{}
	err := client.Request("drds", "DescribeDrdsDB", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDrdsDBResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeDrdsDBData
}

type DescribeDrdsDBData struct {
	DbName     goaliyun.String
	Status     goaliyun.Integer
	CreateTime goaliyun.String
	Msg        goaliyun.String
	Mode       goaliyun.String
}
