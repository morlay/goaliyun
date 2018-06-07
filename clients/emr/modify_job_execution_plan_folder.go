package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyJobExecutionPlanFolderRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              int64  `position:"Query" name:"Id"`
	ParentId        int64  `position:"Query" name:"ParentId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyJobExecutionPlanFolderRequest) Invoke(client goaliyun.Client) (*ModifyJobExecutionPlanFolderResponse, error) {
	resp := &ModifyJobExecutionPlanFolderResponse{}
	err := client.Request("emr", "ModifyJobExecutionPlanFolder", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyJobExecutionPlanFolderResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
	FolderId  goaliyun.String
}
