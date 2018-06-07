package ehpc

import (
	"github.com/morlay/goaliyun"
)

type StopJobsRequest struct {
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *StopJobsRequest) Invoke(client goaliyun.Client) (*StopJobsResponse, error) {
	resp := &StopJobsResponse{}
	err := client.Request("ehpc", "StopJobs", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopJobsResponse struct {
	RequestId goaliyun.String
}
