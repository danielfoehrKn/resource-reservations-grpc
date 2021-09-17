# TODO: does not work. Look at https://github.com/kubernetes/kubernetes/blob/260b5ef1554af1adf149a4fbcd4658dcb40814e4/hack/lib/protoc.sh


# Generates $1/api.pb.go from the protobuf file $1/api.proto
# $1: Full path to the directory where the api.proto file is
function generate() {
  local package=${1}
  gogopath=/Users/d060239/go/bin/protoc-gen-gogo

  export PATH="${gogopath}:${PATH}"
  protoc --go_opt=paths=source_relative -I ../pkg/proto/ ../pkg/proto/pod-resources.proto --go_out=plugins=grpc:../pkg/proto/gen
}

generate