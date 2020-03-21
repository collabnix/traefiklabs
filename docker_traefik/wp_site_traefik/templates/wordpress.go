package templates

//WordPressFiles is a map having files to be created for the wordpress deployment
var WordPressFiles map[string]string

func init() {
	WordPressFiles = make(map[string]string)

	//The docker compose file
	WordPressFiles["docker-compose.yml"] = `
version: "2"
services:
  mariadb:
    image: mariadb
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: {{ .MySQLRootPassword }}
    volumes:
      - ./mysql-data:/var/lib/mysql # I want to manage volumes manually.

  wordpress:
    image: nithinbose/wordpress
    restart: unless-stopped
    environment:
      WORDPRESS_DB_HOST: mariadb
      WORDPRESS_DB_PASSWORD: {{ .MySQLRootPassword }}
      WORDPRESS_TABLE_PREFIX: {{ .WPTablePrefix }}
    volumes:
      - ./wp-content:/var/www/html/wp-content
    depends_on:
      - mariadb
    networks:
      - server
      - default
    labels:
      - 'traefik.enable=true'
      - 'traefik.backend={{ .TraefikBackend }}'
      - 'traefik.frontend.rule=Host:{{ .WPHosts }}'
      - 'traefik.docker.network=pekka-traefik_webgateway'

networks:
  server:
    external:
      name: pekka-traefik_webgateway
`
}

//WordPressFilesData is the data that is required for the templates to render
type WordPressFilesData struct {
	MySQLRootPassword string
	WPTablePrefix     string
	TraefikBackend    string
	WPHosts           string
}
