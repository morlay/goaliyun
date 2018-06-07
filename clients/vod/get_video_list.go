package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetVideoListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	SortBy               string `position:"Query" name:"SortBy"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetVideoListRequest) Invoke(client goaliyun.Client) (*GetVideoListResponse, error) {
	resp := &GetVideoListResponse{}
	err := client.Request("vod", "GetVideoList", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVideoListResponse struct {
	RequestId goaliyun.String
	Total     goaliyun.Integer
	VideoList GetVideoListVideoList
}

type GetVideoListVideo struct {
	VideoId      goaliyun.String
	Title        goaliyun.String
	Tags         goaliyun.String
	Status       goaliyun.String
	Size         goaliyun.Integer
	Duration     goaliyun.Float
	Description  goaliyun.String
	CreateTime   goaliyun.String
	CreationTime goaliyun.String
	ModifyTime   goaliyun.String
	CoverURL     goaliyun.String
	CateId       goaliyun.Integer
	CateName     goaliyun.String
	Snapshots    GetVideoListSnapshotList
}

type GetVideoListVideoList []GetVideoListVideo

func (list *GetVideoListVideoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetVideoListVideo)
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

type GetVideoListSnapshotList []goaliyun.String

func (list *GetVideoListSnapshotList) UnmarshalJSON(data []byte) error {
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
