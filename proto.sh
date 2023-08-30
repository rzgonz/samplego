#!/bin/bash

# Find all .proto files in the current directory
proto_files=$(find . -name "*.proto")

# Generate protobuf and gRPC files for each .proto file
for proto_file in $proto_files; do
    # Extract the file name without the extension
    file_name=$(basename -- "$proto_file")
    file_name_no_ext="${file_name%.*}"

    echo "Generating protobuf and gRPC files for $proto_file..."
    
    # Generate the protobuf and gRPC files
    protoc --go_out=. --go-grpc_out=. "$proto_file"

    echo "Generated files for $file_name_no_ext"
done

echo "All protobuf and gRPC files generated."
