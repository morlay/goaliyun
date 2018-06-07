package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteJobRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteJobRequest) Invoke(client goaliyun.Client) (*DeleteJobResponse, error) {
	resp := &DeleteJobResponse{}
	err := client.Request("emr", "DeleteJob", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteJobResponse struct {
	RequestId goaliyun.String
}
