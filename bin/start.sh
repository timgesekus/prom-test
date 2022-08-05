#!/bin/bash


k3d cluster create cmmc-test --port 80:80@loadbalancer
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/kube-prometheus-stack -f prometheus/values.yaml
helm install cicis cicis
