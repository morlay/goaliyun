package drds

import (
	"github.com/morlay/goaliyun"
)

type DeleteDrdsDBRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DeleteDrdsDBRequest) Invoke(client goaliyun.Client) (*DeleteDrdsDBResponse, error) {
	resp := &DeleteDrdsDBResponse{}
	err := client.Request("drds", "DeleteDrdsDB", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDrdsDBResponse struct {
	RequestId goaliyun.String
	Success   bool
}
