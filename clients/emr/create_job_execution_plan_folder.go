package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateJobExecutionPlanFolderRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	ParentId        int64  `position:"Query" name:"ParentId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateJobExecutionPlanFolderRequest) Invoke(client goaliyun.Client) (*CreateJobExecutionPlanFolderResponse, error) {
	resp := &CreateJobExecutionPlanFolderResponse{}
	err := client.Request("emr", "CreateJobExecutionPlanFolder", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateJobExecutionPlanFolderResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
	FolderId  goaliyun.String
	FolderId1 goaliyun.String
}
