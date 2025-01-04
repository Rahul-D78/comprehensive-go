# Comprehensive Go

This repository has the source code for Comprehensive Go, an unofficial Go cource which covers all aspects of Go, from basic to advancie level concepts. It also includes DSA and Network programming.

The Prometheus server will not automatically scrape your custom metrics unless you configure it to do so. Here's what you need to do:

### Install and configure prometheus for monitoring:

1. Add prometheus helm repo

```console
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

2. Fetch prometheus values File

```console
helm inspect values prometheus-community/kube-prometheus-stack > /tmp/kube-prometheus-stack.values
```

3. Add Scrape Configurations in Prometheus Configuration File

```yaml
## The scrape configuration example below will find master nodes, provided they have the name .*mst.*, relabel the
## port to 2379 and allow etcd scraping provided it is running on all Kubernetes master nodes
##
additionalScrapeConfigs:
  - job_name: 'demo-app'
    scrape_interval: 15s
    metrics_path: /metrics
    static_configs:
      - targets: [demo-app-metrics-service.golang-system.svc.cluster.local:8080]
```

4. Deploy prometheus in the monitoring namespace

```console
helm upgrade -i prometheus prometheus-community/kube-prometheus-stack -f kube-prometheus-stack.values --namespace monitoring --debug
```