#CICIS K8S Prometheus Testbed


Started zwei dummy cicis Prozess in einen Kubernetes Cluster. Diese stellen jeweil zwei Metric Endpoints für Prometheus zur Verfügung.

## Start - Stop

Jeweils im prom-test Ordner des Root Nutzers

Zum initialien installieren und starten
```
bin/start.sh
```

Zum Löschen des Clusters
```
bin/delete.sh
````

Zum reinen Stoppen des Clusters
```
k3d cluster stop cmmc-test
```

Zum reinen Stoppen des Clusters
```
k3d cluster start cmmc-test
```

## Grafana und Prometheus
Das Cluster exportiert Grafana und Prometheus über Port 80 auf folgen Routen

| App      | Route |
| ----------- | ----------- |
| grafana     | http://localhost/grafana      |
| prometheus  | http://localhost/prometheus        |

***Für grafana muss der bei root Nutzer installierte Firefox benutzt werden***

## Grafan Dashboard

Ein sehr rudimentäres Dashboard kann in grafana geöffnet werden. In grafana über das Icon mit den view Quadraten das Daschboard Menü öffnen. Dann unter Browse das Cicis Dashboard nutzen.

## Metriken

Die beiden cicis dummy Prozesse exportieren jeweils zwei Metriken

* _cicis_processed_messages_ ist letztlich eine Zahl die alle zwei Sekunden hochzählt
* _cicis_message_queue_size_ springt alles zwei Sekunden zwischen 80 und 100 hin und her

## Promql query

Um zum Beispiel die verfügbaren Repliklas des deployments cicics_cicis_process-1 zu erfragen 

```
curl 'http://localhost/prometheus/api/v1/query?query=kube_deployment_status_replicas_available\{deployment="cicis-cicis-process-1"\}'
```



