#!/bin/bash

check_status_code() {
	if [ $? -eq 0 ]; then
		echo "返回状态码：$?，创建成功！"
	else
		echo "返回状态码：$?，创建异常，请检查！"
	fi
}

time_sleep() {
	sleep 1
}

kubectl create ns private-disk
kubectl apply -f 1disk-storageclass-rbac.yaml
time_sleep

kubectl apply -f 2disk-storageclass-deployment.yaml
time_sleep

kubectl apply -f 3disk-storageclass.yaml
check_status_code
time_sleep

kubectl apply -f 6ingress-mandatory.yaml
time_sleep

kubectl apply -f 7disk-deployment.yaml
time_sleep

kubectl apply -f 8service-nodeport.yaml
time_sleep

kubectl apply -f 9ingress.yaml
check_status_code