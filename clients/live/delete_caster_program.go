package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterProgramRequest struct {
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterProgramRequest) Invoke(client goaliyun.Client) (*DeleteCasterProgramResponse, error) {
	resp := &DeleteCasterProgramResponse{}
	err := client.Request("live", "DeleteCasterProgram", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterProgramResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
}
