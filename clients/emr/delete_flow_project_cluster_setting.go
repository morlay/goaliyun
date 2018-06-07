package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteFlowProjectClusterSettingRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteFlowProjectClusterSettingRequest) Invoke(client goaliyun.Client) (*DeleteFlowProjectClusterSettingResponse, error) {
	resp := &DeleteFlowProjectClusterSettingResponse{}
	err := client.Request("emr", "DeleteFlowProjectClusterSetting", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFlowProjectClusterSettingResponse struct {
	RequestId goaliyun.String
	Data      bool
}
