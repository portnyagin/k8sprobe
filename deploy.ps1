docker build --no-cache -t k8sprobe/base .
docker build --no-cache -t k8sprobe/alice -f .\Alice.Dockerfile .
docker build -t k8sprobe/bob -f .\Bob.Dockerfile .


docker run --rm -d -p 3000:8080 k8sprobe/alice
docker run --rm -d -p 3005:8080 k8sprobe/bob