package ehpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteJobsRequest struct {
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteJobsRequest) Invoke(client goaliyun.Client) (*DeleteJobsResponse, error) {
	resp := &DeleteJobsResponse{}
	err := client.Request("ehpc", "DeleteJobs", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteJobsResponse struct {
	RequestId goaliyun.String
}
