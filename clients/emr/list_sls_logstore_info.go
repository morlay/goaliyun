package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListSlsLogstoreInfoRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ComponentName   string `position:"Query" name:"ComponentName"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListSlsLogstoreInfoRequest) Invoke(client goaliyun.Client) (*ListSlsLogstoreInfoResponse, error) {
	resp := &ListSlsLogstoreInfoResponse{}
	err := client.Request("emr", "ListSlsLogstoreInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListSlsLogstoreInfoResponse struct {
	RequestId           goaliyun.String
	SlsLogstoreInfoList ListSlsLogstoreInfoSlsLogstoreInfoList
}

type ListSlsLogstoreInfoSlsLogstoreInfo struct {
	Id            goaliyun.Integer
	ServiceName   goaliyun.String
	ComponentName goaliyun.String
	LogstoreName  goaliyun.String
	LogType       goaliyun.String
}

type ListSlsLogstoreInfoSlsLogstoreInfoList []ListSlsLogstoreInfoSlsLogstoreInfo

func (list *ListSlsLogstoreInfoSlsLogstoreInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSlsLogstoreInfoSlsLogstoreInfo)
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
