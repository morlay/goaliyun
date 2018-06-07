package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type PopApiTestRequest struct {
	UserName string `position:"Query" name:"UserName"`
	Pwd      string `position:"Query" name:"Pwd"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *PopApiTestRequest) Invoke(client goaliyun.Client) (*PopApiTestResponse, error) {
	resp := &PopApiTestResponse{}
	err := client.Request("phoenixsp-inner", "PopApiTest", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PopApiTestResponse struct {
	Code      goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
}
