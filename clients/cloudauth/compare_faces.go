package cloudauth

import (
	"github.com/morlay/goaliyun"
)

type CompareFacesRequest struct {
	SourceImageType  string `position:"Query" name:"SourceImageType"`
	ResourceOwnerId  int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp         string `position:"Query" name:"SourceIp"`
	TargetImageType  string `position:"Query" name:"TargetImageType"`
	SourceImageValue string `position:"Query" name:"SourceImageValue"`
	TargetImageValue string `position:"Query" name:"TargetImageValue"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *CompareFacesRequest) Invoke(client goaliyun.Client) (*CompareFacesResponse, error) {
	resp := &CompareFacesResponse{}
	err := client.Request("cloudauth", "CompareFaces", "2018-05-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CompareFacesResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      CompareFacesData
}

type CompareFacesData struct {
	SimilarityScore      goaliyun.Float
	ConfidenceThresholds goaliyun.String
}
