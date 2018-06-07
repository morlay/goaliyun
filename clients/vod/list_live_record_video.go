package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListLiveRecordVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DomainName           string `position:"Query" name:"DomainName"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AppName              string `position:"Query" name:"AppName"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	SortBy               string `position:"Query" name:"SortBy"`
	StreamName           string `position:"Query" name:"StreamName"`
	QueryType            string `position:"Query" name:"QueryType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListLiveRecordVideoRequest) Invoke(client goaliyun.Client) (*ListLiveRecordVideoResponse, error) {
	resp := &ListLiveRecordVideoResponse{}
	err := client.Request("vod", "ListLiveRecordVideo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListLiveRecordVideoResponse struct {
	RequestId           goaliyun.String
	Total               goaliyun.Integer
	LiveRecordVideoList ListLiveRecordVideoLiveRecordVideoList
}

type ListLiveRecordVideoLiveRecordVideo struct {
	StreamName      goaliyun.String
	DomainName      goaliyun.String
	AppName         goaliyun.String
	PlaylistId      goaliyun.String
	RecordStartTime goaliyun.String
	RecordEndTime   goaliyun.String
	PlayInfoList    ListLiveRecordVideoPlayInfoList
	Video           ListLiveRecordVideoVideo
}

type ListLiveRecordVideoPlayInfo struct {
	Width      goaliyun.Integer
	Height     goaliyun.Integer
	Size       goaliyun.Integer
	PlayURL    goaliyun.String
	Bitrate    goaliyun.String
	Definition goaliyun.String
	Duration   goaliyun.String
	Format     goaliyun.String
	Fps        goaliyun.String
	Encrypt    goaliyun.Integer
	Plaintext  goaliyun.String
	Complexity goaliyun.String
	StreamType goaliyun.String
	Rand       goaliyun.String
	JobId      goaliyun.String
}

type ListLiveRecordVideoVideo struct {
	VideoId         goaliyun.String
	Title           goaliyun.String
	Tags            goaliyun.String
	Status          goaliyun.String
	Size            goaliyun.Integer
	Privilege       goaliyun.Integer
	Duration        goaliyun.Float
	Description     goaliyun.String
	CustomerId      goaliyun.Integer
	CreateTime      goaliyun.String
	CreationTime    goaliyun.String
	ModifyTime      goaliyun.String
	CoverURL        goaliyun.String
	CateId          goaliyun.Integer
	CateName        goaliyun.String
	DownloadSwitch  goaliyun.String
	TemplateGroupId goaliyun.String
	Snapshots       ListLiveRecordVideoSnapshotList
}

type ListLiveRecordVideoLiveRecordVideoList []ListLiveRecordVideoLiveRecordVideo

func (list *ListLiveRecordVideoLiveRecordVideoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListLiveRecordVideoLiveRecordVideo)
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

type ListLiveRecordVideoPlayInfoList []ListLiveRecordVideoPlayInfo

func (list *ListLiveRecordVideoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListLiveRecordVideoPlayInfo)
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

type ListLiveRecordVideoSnapshotList []goaliyun.String

func (list *ListLiveRecordVideoSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
