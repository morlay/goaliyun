package ehpc

import (
	"github.com/morlay/goaliyun"
)

type EditJobTemplateRequest struct {
	StderrRedirectPath string `position:"Query" name:"StderrRedirectPath"`
	Variables          string `position:"Query" name:"Variables"`
	RunasUser          string `position:"Query" name:"RunasUser"`
	ReRunable          string `position:"Query" name:"ReRunable"`
	TemplateId         string `position:"Query" name:"TemplateId"`
	Priority           int64  `position:"Query" name:"Priority"`
	CommandLine        string `position:"Query" name:"CommandLine"`
	ArrayRequest       string `position:"Query" name:"ArrayRequest"`
	PackagePath        string `position:"Query" name:"PackagePath"`
	Name               string `position:"Query" name:"Name"`
	StdoutRedirectPath string `position:"Query" name:"StdoutRedirectPath"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *EditJobTemplateRequest) Invoke(client goaliyun.Client) (*EditJobTemplateResponse, error) {
	resp := &EditJobTemplateResponse{}
	err := client.Request("ehpc", "EditJobTemplate", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EditJobTemplateResponse struct {
	RequestId  goaliyun.String
	TemplateId goaliyun.String
}
