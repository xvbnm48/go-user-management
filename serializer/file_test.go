package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/xvbnm48/go-user-management/pb"
	"github.com/xvbnm48/go-user-management/sample"
	"github.com/xvbnm48/go-user-management/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(laptop2, binaryFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
}
