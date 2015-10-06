package main

import (
  "testing"
  "fmt"
  "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "net/http"
  "github.com/jeffail/gabs"
)

func TestElasticsearchReport(t *testing.T) {
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"timestamp":1444099819856,"cluster_name":"elasticsearch_dominiek","status":"green","indices":{"count":0,"shards":{},"docs":{"count":0,"deleted":0},"store":{"size_in_bytes":0,"throttle_time_in_millis":0},"fielddata":{"memory_size_in_bytes":0,"evictions":0},"filter_cache":{"memory_size_in_bytes":0,"evictions":0},"id_cache":{"memory_size_in_bytes":0},"completion":{"size_in_bytes":0},"segments":{"count":0,"memory_in_bytes":0,"index_writer_memory_in_bytes":0,"index_writer_max_memory_in_bytes":0,"version_map_memory_in_bytes":0,"fixed_bit_set_memory_in_bytes":0},"percolate":{"total":0,"time_in_millis":0,"current":0,"memory_size_in_bytes":-1,"memory_size":"-1b","queries":0}},"nodes":{"count":{"total":1,"master_only":0,"data_only":0,"master_data":1,"client":0},"versions":["1.7.2"],"os":{"available_processors":4,"mem":{"total_in_bytes":8589934592},"cpu":[{"vendor":"Intel","model":"MacBook8,1","mhz":1100,"total_cores":4,"total_sockets":4,"cores_per_socket":16,"cache_size_in_bytes":256,"count":1}]},"process":{"cpu":{"percent":0},"open_file_descriptors":{"min":156,"max":156,"avg":156}},"jvm":{"max_uptime_in_millis":49103653,"versions":[{"version":"1.8.0_60","vm_name":"Java HotSpot(TM) 64-Bit Server VM","vm_version":"25.60-b23","vm_vendor":"Oracle Corporation","count":1}],"mem":{"heap_used_in_bytes":67845896,"heap_max_in_bytes":1038876672},"threads":46},"fs":{"total_in_bytes":249678528512,"free_in_bytes":91290705920,"available_in_bytes":91028561920,"disk_reads":0,"disk_writes":0,"disk_io_op":0,"disk_read_size_in_bytes":0,"disk_write_size_in_bytes":0,"disk_io_size_in_bytes":0},"plugins":[]}}`)
  })
  ts := httptest.NewServer(handler)
  defer ts.Close();

  elasticsearch := Elasticsearch{"localhost", 9200};
  data, err := elasticsearch.Monitor();
  assert.Equal(t, err, nil)

  t.Logf("Report: %v", data)

  jsonParsed, err := gabs.ParseJSON([]byte(data))
  assert.Equal(t, err, nil)

  status, _ := jsonParsed.S("report").S("status").S("state").Data().(string);
  assert.Equal(t, status, "ok")
}

func TestElasticsearchClusterStats(t *testing.T) {
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"timestamp":1444099819856,"cluster_name":"elasticsearch_dominiek","status":"green","indices":{"count":0,"shards":{},"docs":{"count":0,"deleted":0},"store":{"size_in_bytes":0,"throttle_time_in_millis":0},"fielddata":{"memory_size_in_bytes":0,"evictions":0},"filter_cache":{"memory_size_in_bytes":0,"evictions":0},"id_cache":{"memory_size_in_bytes":0},"completion":{"size_in_bytes":0},"segments":{"count":0,"memory_in_bytes":0,"index_writer_memory_in_bytes":0,"index_writer_max_memory_in_bytes":0,"version_map_memory_in_bytes":0,"fixed_bit_set_memory_in_bytes":0},"percolate":{"total":0,"time_in_millis":0,"current":0,"memory_size_in_bytes":-1,"memory_size":"-1b","queries":0}},"nodes":{"count":{"total":1,"master_only":0,"data_only":0,"master_data":1,"client":0},"versions":["1.7.2"],"os":{"available_processors":4,"mem":{"total_in_bytes":8589934592},"cpu":[{"vendor":"Intel","model":"MacBook8,1","mhz":1100,"total_cores":4,"total_sockets":4,"cores_per_socket":16,"cache_size_in_bytes":256,"count":1}]},"process":{"cpu":{"percent":0},"open_file_descriptors":{"min":156,"max":156,"avg":156}},"jvm":{"max_uptime_in_millis":49103653,"versions":[{"version":"1.8.0_60","vm_name":"Java HotSpot(TM) 64-Bit Server VM","vm_version":"25.60-b23","vm_vendor":"Oracle Corporation","count":1}],"mem":{"heap_used_in_bytes":67845896,"heap_max_in_bytes":1038876672},"threads":46},"fs":{"total_in_bytes":249678528512,"free_in_bytes":91290705920,"available_in_bytes":91028561920,"disk_reads":0,"disk_writes":0,"disk_io_op":0,"disk_read_size_in_bytes":0,"disk_write_size_in_bytes":0,"disk_io_size_in_bytes":0},"plugins":[]}}`)
  })
  ts := httptest.NewServer(handler)
  defer ts.Close();

  elasticsearch := Elasticsearch{"localhost", 9200};
  data, err := elasticsearch.ClusterStats();
  assert.Equal(t, err, nil)

  jsonParsed, err := gabs.ParseJSON([]byte(data))
  assert.Equal(t, err, nil)

  status, _ := jsonParsed.S("status").Data().(string);
  assert.Equal(t, status, "green")
}
