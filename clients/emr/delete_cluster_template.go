package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteClusterTemplateRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteClusterTemplateRequest) Invoke(client goaliyun.Client) (*DeleteClusterTemplateResponse, error) {
	resp := &DeleteClusterTemplateResponse{}
	err := client.Request("emr", "DeleteClusterTemplate", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteClusterTemplateResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
}
