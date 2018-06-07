package drds

import (
	"github.com/morlay/goaliyun"
)

type CreateDrdsDBRequest struct {
	Encode         string `position:"Query" name:"Encode"`
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	RdsInstances   string `position:"Query" name:"RdsInstances"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreateDrdsDBRequest) Invoke(client goaliyun.Client) (*CreateDrdsDBResponse, error) {
	resp := &CreateDrdsDBResponse{}
	err := client.Request("drds", "CreateDrdsDB", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDrdsDBResponse struct {
	RequestId goaliyun.String
	Success   bool
}
