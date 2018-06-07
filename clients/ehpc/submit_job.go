package ehpc

import (
	"github.com/morlay/goaliyun"
)

type SubmitJobRequest struct {
	StderrRedirectPath string `position:"Query" name:"StderrRedirectPath"`
	Variables          string `position:"Query" name:"Variables"`
	RunasUserPassword  string `position:"Query" name:"RunasUserPassword"`
	RunasUser          string `position:"Query" name:"RunasUser"`
	ClusterId          string `position:"Query" name:"ClusterId"`
	ReRunable          string `position:"Query" name:"ReRunable"`
	Priority           int64  `position:"Query" name:"Priority"`
	CommandLine        string `position:"Query" name:"CommandLine"`
	ArrayRequest       string `position:"Query" name:"ArrayRequest"`
	PackagePath        string `position:"Query" name:"PackagePath"`
	Name               string `position:"Query" name:"Name"`
	StdoutRedirectPath string `position:"Query" name:"StdoutRedirectPath"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *SubmitJobRequest) Invoke(client goaliyun.Client) (*SubmitJobResponse, error) {
	resp := &SubmitJobResponse{}
	err := client.Request("ehpc", "SubmitJob", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
