# prerequisites
You have to install `docker`, `kubectl`, `kind`

# usage
`make start` to start the cluster in KinD (Kubernetes in Docker). This will create two deployments service-a and service-b.
Those services expose a `/health` route you can access them via `http://localhost:80/service_a/health` and `http://localhost:80/service_b/health`

 * check service-a health status `curl curl http://localhost/service_a/health`
 * check service-b health status `curl curl http://localhost/service_b/health`