package templates

//TraefikFiles is a map having files to be created to use traefik
var TraefikFiles map[string]string

func init() {
	TraefikFiles = make(map[string]string)

	//The docker compose file
	TraefikFiles["docker-compose.yml"] = `
version: '2'
services:
  proxy:
    image: traefik
    restart: unless-stopped
    command: --logLevel=DEBUG
    networks:
      - webgateway
    ports:
      - "80:80"
      - "443:443"
    labels:
      - "traefik.enable=true"
      - "traefik.frontend.rule=Host:{{ .TraefikDashboardURL }}" # traefik will request SSL from letsencrypt
      - "traefik.port=8080" # enable this port in so traefik can proxy to it
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./traefik.toml:/traefik.toml

networks:
  webgateway:
    driver: bridge
`

	//Traefik config file
	TraefikFiles["traefik.toml"] = `
defaultEntryPoints = ["http", "https"]

[web]
address = ":8080"

[entryPoints]

[entryPoints.http]
address = ":80"
[entryPoints.http.redirect]
entryPoint = "https"

[entryPoints.https]
address = ":443"
compress = true
[entryPoints.https.tls]

[acme]
email = "{{ .AcmeEmail }}"
storage = "acme.json"
entryPoint = "https"
onDemand = false
OnHostRule = true
[acme.httpChallenge]
entryPoint = "http"

[docker]
endpoint = "unix:///var/run/docker.sock"
domain = "traefik.example.com"
watch = true
exposedbydefault = false
`

}

//TraefikFilesData is the data that is required for the templates to render
type TraefikFilesData struct {
	TraefikDashboardURL string
	AcmeEmail           string
}
