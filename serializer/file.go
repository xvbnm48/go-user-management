package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteProtoBuffToJSONFile
// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON format: %w", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON to file: %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile write protobuf to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary to file: %w", err)
	}

	return nil
}

// ReadProtobuffFromBinaryFile read protobuf from binary file
func ReadProtobufFromBinaryFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}
