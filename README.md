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
docker build -t k8sprobe/Bob -f .\Bob.Dockerfile . 