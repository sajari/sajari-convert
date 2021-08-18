#!/usr/bin/env sh

main() {
  export NAME=docd
  export VERSION=alpine
  
  echo "Building ${NAME} for ${VERSION}..."

  GOOS=linux GOARCH=amd64 go build -o $NAME || exit 1
  cp docd $VERSION
  cd $VERSION || exit 1
  docker build -t "$NAME:$1" -t "$NAME:latest" . || exit 1
  rm docd  
}

main "$@"