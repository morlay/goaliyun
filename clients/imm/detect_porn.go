package imm

import (
	"github.com/morlay/goaliyun"
)

type DetectPornRequest struct {
	SrcType  string `position:"Query" name:"SrcType"`
	Engine   string `position:"Query" name:"Engine"`
	Project  string `position:"Query" name:"Project"`
	Url      string `position:"Query" name:"Url"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DetectPornRequest) Invoke(client goaliyun.Client) (*DetectPornResponse, error) {
	resp := &DetectPornResponse{}
	err := client.Request("imm", "DetectPorn", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetectPornResponse struct {
	RequestId goaliyun.String
	Score     goaliyun.Float
}
