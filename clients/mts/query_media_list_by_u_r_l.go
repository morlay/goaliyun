package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMediaListByURLRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	FileURLs             string `position:"Query" name:"FileURLs"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryMediaListByURLRequest) Invoke(client goaliyun.Client) (*QueryMediaListByURLResponse, error) {
	resp := &QueryMediaListByURLResponse{}
	err := client.Request("mts", "QueryMediaListByURL", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMediaListByURLResponse struct {
	RequestId        goaliyun.String
	MediaList        QueryMediaListByURLMediaList
	NonExistFileURLs QueryMediaListByURLNonExistFileURLList
}

type QueryMediaListByURLMedia struct {
	MediaId      goaliyun.String
	Title        goaliyun.String
	Description  goaliyun.String
	CoverURL     goaliyun.String
	CateId       goaliyun.Integer
	Duration     goaliyun.String
	Format       goaliyun.String
	Size         goaliyun.String
	Bitrate      goaliyun.String
	Width        goaliyun.String
	Height       goaliyun.String
	Fps          goaliyun.String
	PublishState goaliyun.String
	CreationTime goaliyun.String
	PlayList     QueryMediaListByURLPlayList
	SnapshotList QueryMediaListByURLSnapshotList
	Tags         QueryMediaListByURLTagList
	RunIdList    QueryMediaListByURLRunIdListList
	File         QueryMediaListByURLFile
	MediaInfo    QueryMediaListByURLMediaInfo
}

type QueryMediaListByURLPlay struct {
	ActivityName      goaliyun.String
	MediaWorkflowId   goaliyun.String
	MediaWorkflowName goaliyun.String
	Duration          goaliyun.String
	Format            goaliyun.String
	Size              goaliyun.String
	Bitrate           goaliyun.String
	Width             goaliyun.String
	Height            goaliyun.String
	Fps               goaliyun.String
	Encryption        goaliyun.String
	File1             QueryMediaListByURLFile1
}

type QueryMediaListByURLFile1 struct {
	URL   goaliyun.String
	State goaliyun.String
}

type QueryMediaListByURLSnapshot struct {
	Type              goaliyun.String
	MediaWorkflowId   goaliyun.String
	MediaWorkflowName goaliyun.String
	ActivityName      goaliyun.String
	Count             goaliyun.String
	File2             QueryMediaListByURLFile2
}

type QueryMediaListByURLFile2 struct {
	URL   goaliyun.String
	State goaliyun.String
}

type QueryMediaListByURLFile struct {
	URL   goaliyun.String
	State goaliyun.String
}

type QueryMediaListByURLMediaInfo struct {
	Streams QueryMediaListByURLStreams
	Format  QueryMediaListByURLFormat
}

type QueryMediaListByURLStreams struct {
	VideoStreamList    QueryMediaListByURLVideoStreamList
	AudioStreamList    QueryMediaListByURLAudioStreamList
	SubtitleStreamList QueryMediaListByURLSubtitleStreamList
}

type QueryMediaListByURLVideoStream struct {
	Index          goaliyun.String
	CodecName      goaliyun.String
	CodecLongName  goaliyun.String
	Profile        goaliyun.String
	CodecTimeBase  goaliyun.String
	CodecTagString goaliyun.String
	CodecTag       goaliyun.String
	Width          goaliyun.String
	Height         goaliyun.String
	HasBFrames     goaliyun.String
	Sar            goaliyun.String
	Dar            goaliyun.String
	PixFmt         goaliyun.String
	Level          goaliyun.String
	Fps            goaliyun.String
	AvgFPS         goaliyun.String
	Timebase       goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Bitrate        goaliyun.String
	NumFrames      goaliyun.String
	Lang           goaliyun.String
	Rotate         goaliyun.String
	NetworkCost    QueryMediaListByURLNetworkCost
}

type QueryMediaListByURLNetworkCost struct {
	PreloadTime   goaliyun.String
	CostBandwidth goaliyun.String
	AvgBitrate    goaliyun.String
}

type QueryMediaListByURLAudioStream struct {
	Index          goaliyun.String
	CodecName      goaliyun.String
	CodecTimeBase  goaliyun.String
	CodecLongName  goaliyun.String
	CodecTagString goaliyun.String
	CodecTag       goaliyun.String
	SampleFmt      goaliyun.String
	Samplerate     goaliyun.String
	Channels       goaliyun.String
	ChannelLayout  goaliyun.String
	Timebase       goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Bitrate        goaliyun.String
	NumFrames      goaliyun.String
	Lang           goaliyun.String
}

type QueryMediaListByURLSubtitleStream struct {
	Index goaliyun.String
	Lang  goaliyun.String
}

type QueryMediaListByURLFormat struct {
	NumStreams     goaliyun.String
	NumPrograms    goaliyun.String
	FormatName     goaliyun.String
	FormatLongName goaliyun.String
	StartTime      goaliyun.String
	Duration       goaliyun.String
	Size           goaliyun.String
	Bitrate        goaliyun.String
}

type QueryMediaListByURLMediaList []QueryMediaListByURLMedia

func (list *QueryMediaListByURLMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLMedia)
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

type QueryMediaListByURLNonExistFileURLList []goaliyun.String

func (list *QueryMediaListByURLNonExistFileURLList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLPlayList []QueryMediaListByURLPlay

func (list *QueryMediaListByURLPlayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLPlay)
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

type QueryMediaListByURLSnapshotList []QueryMediaListByURLSnapshot

func (list *QueryMediaListByURLSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLSnapshot)
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

type QueryMediaListByURLTagList []goaliyun.String

func (list *QueryMediaListByURLTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLRunIdListList []goaliyun.String

func (list *QueryMediaListByURLRunIdListList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLVideoStreamList []QueryMediaListByURLVideoStream

func (list *QueryMediaListByURLVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLVideoStream)
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

type QueryMediaListByURLAudioStreamList []QueryMediaListByURLAudioStream

func (list *QueryMediaListByURLAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLAudioStream)
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

type QueryMediaListByURLSubtitleStreamList []QueryMediaListByURLSubtitleStream

func (list *QueryMediaListByURLSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLSubtitleStream)
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
