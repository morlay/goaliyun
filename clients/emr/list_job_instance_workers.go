package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobInstanceWorkersRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListJobInstanceWorkersRequest) Invoke(client goaliyun.Client) (*ListJobInstanceWorkersResponse, error) {
	resp := &ListJobInstanceWorkersResponse{}
	err := client.Request("emr", "ListJobInstanceWorkers", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobInstanceWorkersResponse struct {
	RequestId          goaliyun.String
	JobInstanceWorkers ListJobInstanceWorkersJobInstanceWorkerInfoList
}

type ListJobInstanceWorkersJobInstanceWorkerInfo struct {
	ApplicationId goaliyun.String
	InstanceInfo  goaliyun.String
	ContainerInfo goaliyun.String
}

type ListJobInstanceWorkersJobInstanceWorkerInfoList []ListJobInstanceWorkersJobInstanceWorkerInfo

func (list *ListJobInstanceWorkersJobInstanceWorkerInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobInstanceWorkersJobInstanceWorkerInfo)
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
