# How to use Traefik

![](https://raw.githubusercontent.com/collabnix/traefiklabs/master/docker_traefik/Quick_Start/quick_start.png)

```
git clone https://github.com/collabnix/traefiklabs/
cd docker_traefik/Quick_Start
docker-compose up -d 


```


## Create `docker-compose.yml` with following content: 
```
version: '3'

services:
  reverse-proxy:
    image: traefik:alpine # The official Traefik docker image
    command: --api --docker # Enables the web UI and tells Traefik to listen to docker
    ports:
      - "80:80"     # The HTTP port
      - "8080:8080" # The Web UI (enabled by --api)
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # So that Traefik can listen to the Docker events

```
## Docker-compose up 

```
docker-compose up
Creating network "quick_start_default" with the default driver
Pulling reverse-proxy (traefik:alpine)...
alpine: Pulling from library/traefik
9d48c3bd43c5: Pull complete
1be319f51f9f: Pull complete
aacb2071e114: Pull complete
c6c61df92853: Pull complete
Creating quick_start_reverse-proxy_1 ... done
Attaching to quick_start_reverse-proxy_1
````
## Open Port 8080 for Traefik dashboard
![](https://raw.githubusercontent.com/collabnix/traefiklabs/master/docker_traefik/Quick_Start/click_8080.png)
![](https://raw.githubusercontent.com/collabnix/traefiklabs/master/docker_traefik/Quick_Start/dashboard.png)


## Create `traefik.toml.quickstart` with following content: 

```
# quick-start stuff:

defaultEntryPoints = ["http"]

[docker]
  endpoint = "unix:///var/run/docker.sock"

[api]
  dashboard = true
  entrypoint = "dashboard"

[entryPoints]
  [entryPoints.http]
  address = ":80"

  [entryPoints.dashboard]
  address = ":8080"
```
## stop runing conntainer and remove it 
```
$ docker-compose rm -f
Going to remove quick_start_reverse-proxy_1
Removing quick_start_reverse-proxy_1 ... done
[node1] (local) root@192.168.0.18 ~/traefiklabs/docker_traefik/Quick_Start
$ 
```
