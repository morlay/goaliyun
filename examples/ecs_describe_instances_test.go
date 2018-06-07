package examples

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/morlay/goaliyun"
	"github.com/morlay/goaliyun/clients/ecs"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestEcsDescribeInstances(t *testing.T) {
	tt := assert.New(t)

	client := goaliyun.NewClientWithCredential(
		os.Getenv("ALIYUN_ACCESS_KEY"),
		os.Getenv("ALIYUN_SECRET_KEY"),
		goaliyun.ClientOptionWithTransports(goaliyun.NewLogTransportWithBody()),
	)

	req := ecs.DescribeInstancesRequest{
		RegionId: "cn-beijing",
	}
	resp, err := req.Invoke(client)
	tt.NoError(err)

	tt.True(len(resp.Instances) > 1)

	for _, region := range resp.Instances {
		t.Log(region)
	}
}
