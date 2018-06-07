package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListServiceLogRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	LogstoreName    string `position:"Query" name:"LogstoreName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListServiceLogRequest) Invoke(client goaliyun.Client) (*ListServiceLogResponse, error) {
	resp := &ListServiceLogResponse{}
	err := client.Request("emr", "ListServiceLog", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListServiceLogResponse struct {
	RequestId   goaliyun.String
	LogFileList ListServiceLogLogFileList
}

type ListServiceLogLogFile struct {
	FileName     goaliyun.String
	Size         goaliyun.Integer
	HostName     goaliyun.String
	LastModified goaliyun.Integer
}

type ListServiceLogLogFileList []ListServiceLogLogFile

func (list *ListServiceLogLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListServiceLogLogFile)
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
