package utils

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	path := "../configs/apiserver.yaml"
	sampleConfig := &Config{
		BindAddr:         ":8081",
		DatabaseURL:      "mongodb://127.0.0.1:27017",
		DbName:           "elCharge",
		DbUserCollection: "user",
		JWTKey:           "21d5680b6abb9cda96860971d05c8d9459b3f08b87a930ddf9b298148780b4234d5b6098593523e73bc3a0cc569270477afe033f834cb4b856071864ecad6f47aed5b68cff79434e1277c3cc31d3d5b0a49afa7b436e943b08c9e5a73abbce687ddcc8bd7bb070096358446f89123cfdf1ab35d10118b3f863da5a4631f0ebdf3cbe847b543477b800da5ac3e2b47debe9494322435a8c977e535afb46f0d0ddd1af65cbdb1eef565c69222bcaa382aed0e92f645185ff3777b7b9b8ef049687764aee479667f331359b431cae4cdae6644c5eb7f1f3990de2a50cc9cf334b50e192e9beb834b63e9f77d47fe3ff2e4c08bcb4ee1b5dffaaacca91ce0fd9da18182a6403694726f60ad4e44cfcfed098fd7fc861217c399a52fee55d52b1f0a0",
		RedisDB:          "localhost:6379",
	}
	testConfig, err := NewConfig(path)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(testConfig, sampleConfig) {
		t.Error("Not equal")
	}
}

func TestNewConfig2(t *testing.T) {
	path := "../configs/apiserver1.yaml"
	testConfig, err := NewConfig(path)
	if err == nil {
		t.Error(err)
	}
	if testConfig != nil {
		t.Error(testConfig)
	}
}
