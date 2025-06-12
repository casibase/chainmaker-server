package examples

import (
	chainmaker_sdk_go "chainmaker.org/chainmaker/sdk-go/v2"
	"fmt"
	"testing"
)

const (
	sdkConfigOrg1Client1Path = "./sdk_config.yml"
)

func TestChainClientGetChainMakerServerVersion(t *testing.T) {
	client, err := chainmaker_sdk_go.NewChainClient(
		chainmaker_sdk_go.WithConfPath(sdkConfigOrg1Client1Path),
	)
	if err != nil {
		t.Errorf("create chain client failed, err: %s", err)
	}
	version, err := client.GetChainMakerServerVersion()
	if err != nil {
		t.Errorf("get chainserver version failed, err: %s", err)
	}
	fmt.Println("get chainserver version:", version)
}
