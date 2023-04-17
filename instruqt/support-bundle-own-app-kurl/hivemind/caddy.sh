podman run -d --rm --name caddy -p 8080:80 \
        --label "io.containers.autoupdate=registry" \
        --mount type=bind,src=/home/dan/Projects/wss/Caddyfile,dst=/etc/caddy/Caddyfile,ro \
        --mount type=bind,src=/home/dan/Projects/wss/html,dst=/var/www/html,ro \
        docker.io/library/caddy
