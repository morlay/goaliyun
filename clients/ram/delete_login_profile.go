package ram

import (
	"github.com/morlay/goaliyun"
)

type DeleteLoginProfileRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteLoginProfileRequest) Invoke(client goaliyun.Client) (*DeleteLoginProfileResponse, error) {
	resp := &DeleteLoginProfileResponse{}
	err := client.Request("ram", "DeleteLoginProfile", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLoginProfileResponse struct {
	RequestId goaliyun.String
}
