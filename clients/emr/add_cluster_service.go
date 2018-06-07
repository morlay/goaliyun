package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddClusterServiceRequest struct {
	ResourceOwnerId int64                         `position:"Query" name:"ResourceOwnerId"`
	Services        *AddClusterServiceServiceList `position:"Query" type:"Repeated" name:"Service"`
	ClusterId       string                        `position:"Query" name:"ClusterId"`
	RegionId        string                        `position:"Query" name:"RegionId"`
}

func (req *AddClusterServiceRequest) Invoke(client goaliyun.Client) (*AddClusterServiceResponse, error) {
	resp := &AddClusterServiceResponse{}
	err := client.Request("emr", "AddClusterService", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddClusterServiceService struct {
	ServiceName string `name:"ServiceName"`
}

type AddClusterServiceResponse struct {
	RequestId goaliyun.String
}

type AddClusterServiceServiceList []AddClusterServiceService

func (list *AddClusterServiceServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddClusterServiceService)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
