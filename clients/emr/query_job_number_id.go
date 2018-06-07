package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryJobNumberIdRequest struct {
	JobBizId        string `position:"Query" name:"JobBizId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryJobNumberIdRequest) Invoke(client goaliyun.Client) (*QueryJobNumberIdResponse, error) {
	resp := &QueryJobNumberIdResponse{}
	err := client.Request("emr", "QueryJobNumberId", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryJobNumberIdResponse struct {
	RequestId   goaliyun.String
	JobNumberId goaliyun.String
}
