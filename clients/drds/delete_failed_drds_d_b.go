package drds

import (
	"github.com/morlay/goaliyun"
)

type DeleteFailedDrdsDBRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DeleteFailedDrdsDBRequest) Invoke(client goaliyun.Client) (*DeleteFailedDrdsDBResponse, error) {
	resp := &DeleteFailedDrdsDBResponse{}
	err := client.Request("drds", "DeleteFailedDrdsDB", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFailedDrdsDBResponse struct {
	RequestId goaliyun.String
	Success   bool
}
