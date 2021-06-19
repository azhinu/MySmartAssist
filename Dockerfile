# Choose any golang image, just make sure it doesn't have -onbuild
FROM golang:1

RUN apt-get update && apt-get -y install ffmpeg && mkdir -p /go/src/app && go env -w GO111MODULE=auto

# Everything below is copied manually from the official -onbuild image,
# with the ONBUILD keywords removed.
COPY . /go/src/app
WORKDIR /go/src/app
CMD ["go get & go build"]
