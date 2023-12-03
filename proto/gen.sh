for file in $(git ls-files '*.proto'); do
    protoc --go_out=./ ./"$file"
done