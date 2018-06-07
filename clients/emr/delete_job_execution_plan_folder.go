package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteJobExecutionPlanFolderRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              int64  `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteJobExecutionPlanFolderRequest) Invoke(client goaliyun.Client) (*DeleteJobExecutionPlanFolderResponse, error) {
	resp := &DeleteJobExecutionPlanFolderResponse{}
	err := client.Request("emr", "DeleteJobExecutionPlanFolder", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteJobExecutionPlanFolderResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
	FolderId  goaliyun.String
}
