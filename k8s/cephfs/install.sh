#!/bin/bash 

timer(){
	sleep 2
}

kubectl create ns cephfs
kubectl create ns private-disk
sleep 3

kubectl apply -f ./rbac/
timer

kubectl apply -f storageclass-cephfs.yaml
timer
kubectl apply -f 6ingress-mandatory.yaml
timer
kubectl apply -f 7disk-deployment.yaml
timer
kubectl apply -f 8service-nodeport.yaml
timer
kubectl apply -f 9ingress.yaml
