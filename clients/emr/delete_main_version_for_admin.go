package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteMainVersionForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteMainVersionForAdminRequest) Invoke(client goaliyun.Client) (*DeleteMainVersionForAdminResponse, error) {
	resp := &DeleteMainVersionForAdminResponse{}
	err := client.Request("emr", "DeleteMainVersionForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMainVersionForAdminResponse struct {
	RequestId goaliyun.String
}
