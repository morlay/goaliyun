package examples

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/morlay/goaliyun"
	"github.com/morlay/goaliyun/clients/cms"
)

func TestCms(t *testing.T) {
	client := goaliyun.NewClientWithCredential(
		os.Getenv("ALIYUN_ACCESS_KEY"),
		os.Getenv("ALIYUN_SECRET_KEY"),
		goaliyun.ClientOptionWithTransports(goaliyun.NewLogTransportWithBody()),
	)

	req := cms.QueryMetricListRequest{
		Project:  "acs_ecs_dashboard",
		RegionId: "cn-beijing",
	}

	_, err := req.Invoke(client)
	spew.Dump(err)
}
