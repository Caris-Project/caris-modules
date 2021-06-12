#!bin/bash

source ~/.bash_profile

# Setup path
base_path=${PWD}
proto_path=$base_path/define

cd $proto_path

# Loop in list subfolders
for d in $(ls -d */) ; do
target_path="$base_path/define/$d"
model_path="$base_path/model/$d"

# Mkdir if not existed
mkdir -p $model_path

# Cd to proto folder and generate
cd $target_path
protoc ./* --proto_path=. --proto_path=$proto_path --go_out=plugins=grpc:$model_path
  echo "- Done $d"
done

echo "Protobuf files generated!"
cd $base_path