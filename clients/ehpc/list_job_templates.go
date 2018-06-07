package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobTemplatesRequest struct {
	Name       string `position:"Query" name:"Name"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListJobTemplatesRequest) Invoke(client goaliyun.Client) (*ListJobTemplatesResponse, error) {
	resp := &ListJobTemplatesResponse{}
	err := client.Request("ehpc", "ListJobTemplates", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobTemplatesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Templates  ListJobTemplatesJobTemplatesList
}

type ListJobTemplatesJobTemplates struct {
	Id                 goaliyun.String
	Name               goaliyun.String
	CommandLine        goaliyun.String
	RunasUser          goaliyun.String
	Priority           goaliyun.Integer
	PackagePath        goaliyun.String
	StdoutRedirectPath goaliyun.String
	StderrRedirectPath goaliyun.String
	ReRunable          bool
	ArrayRequest       goaliyun.String
	Variables          goaliyun.String
}

type ListJobTemplatesJobTemplatesList []ListJobTemplatesJobTemplates

func (list *ListJobTemplatesJobTemplatesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobTemplatesJobTemplates)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
