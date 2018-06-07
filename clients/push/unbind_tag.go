package push

import (
	"github.com/morlay/goaliyun"
)

type UnbindTagRequest struct {
	TagName   string `position:"Query" name:"TagName"`
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *UnbindTagRequest) Invoke(client goaliyun.Client) (*UnbindTagResponse, error) {
	resp := &UnbindTagResponse{}
	err := client.Request("push", "UnbindTag", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindTagResponse struct {
	RequestId goaliyun.String
}
