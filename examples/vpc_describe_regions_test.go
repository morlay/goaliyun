package examples

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/morlay/goaliyun"
	"github.com/morlay/goaliyun/clients/vpc"
)

func TestVPCDescribeRegions(t *testing.T) {
	tt := assert.New(t)

	client := goaliyun.NewClientWithCredential(
		os.Getenv("ALIYUN_ACCESS_KEY"),
		os.Getenv("ALIYUN_SECRET_KEY"),
		goaliyun.ClientOptionWithTransports(goaliyun.NewLogTransportWithBody()),
	)

	req := vpc.DescribeRegionsRequest{
		RegionId: "cn-beijing",
	}
	resp, err := req.Invoke(client)
	tt.NoError(err)

	tt.True(len(resp.Regions) > 1)

	for _, region := range resp.Regions {
		t.Log(region)
	}
}

func TestVPCDescribeRegionsWithError(t *testing.T) {
	tt := assert.New(t)

	client := goaliyun.NewClientWithCredential(
		os.Getenv("ALIYUN_ACCESS_KEY"),
		"",
	)

	req := vpc.DescribeRegionsRequest{}
	_, err := req.Invoke(client)
	tt.Error(err)
}
