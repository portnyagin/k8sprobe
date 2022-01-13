# k8sprobe
trivial service for probe internal communications


Сервис должен иметь два метода api
- При вызове первого в консоль пишется имя сервиса
- ПРи вызове второго метода, вызывается первый метод соседнего сервиса

Параметры: 
имя сервиса
порт
урлы соседей


docker build -t k8sprobe/base .
docker build -t k8sprobe/alice -f .\Alice.Dockerfile .
docker build -t k8sprobe/bob -f .\Bob.Dockerfile .

docker run --rm -d -p 3000:8080 k8sprobe/alice
docker run --rm -d -p 3005:8080 k8sprobe/bob

user definew bridge
docker network create k8sprobe_network
docker run -d -p 3000:8080 --net k8sprobe_network --name alice -d k8sprobe/alice
docker run -d -p 3005:8080 --net k8sprobe_network --name bob -d k8sprobe/bob