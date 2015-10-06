# Elasticsearch monitoring plugin for Atoll

CLI creates JSON that's used by Atoll.

Monitoring stats are focussed on core ES KPIs:

* Cluster status
* Cache sizes
* Heap usage

## Build

```bash
make
```

## Run

```bash
./bin/atoll-elasticsearch
```

## Test

```bash
make test
make test.verbose
```
