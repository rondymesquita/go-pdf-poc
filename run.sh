docker build -t pdf .
docker run \
  -ti \
  --rm \
  -v $(pwd):/go/src/go-pdf \
  --cpus=1 \
  --memory=3g \
  pdf bash
