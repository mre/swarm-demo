# Grafana-InfluxDB Metrics Backend

This is a stand-alone docker-compose setup for metrics.  
It starts a container with InfluxDB 0.9.x as a backend for time-series data.  
After that it starts Grafana, to graph the data.

You can send data to the endpoint using the InfluxDB HTTP API on port 8086 or Graphite data
on port 2003.
