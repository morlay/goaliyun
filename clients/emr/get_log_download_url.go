package emr

import (
	"github.com/morlay/goaliyun"
)

type GetLogDownloadUrlRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	LogstoreName    string `position:"Query" name:"LogstoreName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	LogFileName     string `position:"Query" name:"LogFileName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetLogDownloadUrlRequest) Invoke(client goaliyun.Client) (*GetLogDownloadUrlResponse, error) {
	resp := &GetLogDownloadUrlResponse{}
	err := client.Request("emr", "GetLogDownloadUrl", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetLogDownloadUrlResponse struct {
	RequestId                goaliyun.String
	DownloadUrlBase64Encoded goaliyun.String
}
