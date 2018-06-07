package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RenewClusterRequest struct {
	ResourceOwnerId int64                       `position:"Query" name:"ResourceOwnerId"`
	RenewEcsDos     *RenewClusterRenewEcsDoList `position:"Query" type:"Repeated" name:"RenewEcsDo"`
	ClusterId       string                      `position:"Query" name:"ClusterId"`
	RegionId        string                      `position:"Query" name:"RegionId"`
}

func (req *RenewClusterRequest) Invoke(client goaliyun.Client) (*RenewClusterResponse, error) {
	resp := &RenewClusterResponse{}
	err := client.Request("emr", "RenewCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewClusterRenewEcsDo struct {
	EcsId     string `name:"EcsId"`
	EcsPeriod string `name:"EcsPeriod"`
	EmrPeriod string `name:"EmrPeriod"`
}

type RenewClusterResponse struct {
	RequestId      goaliyun.String
	EcsOrderIdList goaliyun.String
	EmrOrderIdList goaliyun.String
}

type RenewClusterRenewEcsDoList []RenewClusterRenewEcsDo

func (list *RenewClusterRenewEcsDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RenewClusterRenewEcsDo)
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
