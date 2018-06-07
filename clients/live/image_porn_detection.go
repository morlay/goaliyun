package live

import (
	"github.com/morlay/goaliyun"
)

type ImagePornDetectionRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ImageUrl      string `position:"Query" name:"ImageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ImagePornDetectionRequest) Invoke(client goaliyun.Client) (*ImagePornDetectionResponse, error) {
	resp := &ImagePornDetectionResponse{}
	err := client.Request("live", "ImagePornDetection", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImagePornDetectionResponse struct {
	RequestId goaliyun.String
	Label     goaliyun.String
	Rate      goaliyun.Float
}
