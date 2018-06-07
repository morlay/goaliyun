package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateMainVersionForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
	Content         string `position:"Query" name:"Content"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateMainVersionForAdminRequest) Invoke(client goaliyun.Client) (*CreateMainVersionForAdminResponse, error) {
	resp := &CreateMainVersionForAdminResponse{}
	err := client.Request("emr", "CreateMainVersionForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMainVersionForAdminResponse struct {
	RequestId goaliyun.String
}
