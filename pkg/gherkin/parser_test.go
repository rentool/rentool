package gherkin

import (
	"github.com/spf13/viper"
	"testing"
)

func TestAga(t *testing.T) {
	expectedFilePath := "fixtures/complete.feature-nodes.yaml"
	config := viper.New()
	config.SetConfigFile(expectedFilePath)
	err := config.ReadInConfig()
	if err != nil {
		t.Fatalf("failed to read fixtures file %v: %v", expectedFilePath, err)
	}

	expectedFeature := &Feature{}
	err = config.UnmarshalKey("feature", expectedFeature)
	if err != nil {
		t.Fatalf("failed to unmarshal expected feature to Feature struct: %v", err)
	}

}
