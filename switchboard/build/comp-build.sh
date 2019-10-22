deployname=switch

echo "Building component $deployname"

# Cleanup
mkdir -p bin
rm -rf bin/*
cd bin


# Build the project
# git config --global url."ssh://git@github.hpe.com".insteadOf "https://github.hpe.com" 
export CGO_ENABLED=0
export GO111MODULE=on
# export GOPRIVATE=github.hpe.com/*
go mod download
go build ../../cmd/Start.go

# Build the docker
docker rmi $deployname -f
docker build -t $deployname ../../../ -f ../dockerImage
