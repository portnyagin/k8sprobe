# Деплой в миникуб
#point shell  to minikube's docker demon
eval $(minikube -p minikube docker-env)
docker build --no-cache -t k8sprobe/base .
docker build --no-cache -t k8sprobe/alice -f .\Alice.Dockerfile .
docker build -t k8sprobe/bob -f .\Bob.Dockerfile .

kubectl create -f ./build/deploy.yaml
