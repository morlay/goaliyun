package push

import (
	"github.com/morlay/goaliyun"
)

type RemoveTagRequest struct {
	TagName  string `position:"Query" name:"TagName"`
	AppKey   int64  `position:"Query" name:"AppKey"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *RemoveTagRequest) Invoke(client goaliyun.Client) (*RemoveTagResponse, error) {
	resp := &RemoveTagResponse{}
	err := client.Request("push", "RemoveTag", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveTagResponse struct {
	RequestId goaliyun.String
}
