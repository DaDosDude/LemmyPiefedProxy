# Deployment tutorial

A full, real deployment: the proxy running behind a reverse proxy,
paired with mlmym as the actual frontend users visit, with automatic
image updates. This is the exact setup this project was tested against
in production, using Caddy specifically, though the same steps apply
with Nginx or any other reverse proxy.

## 1. Decide on a domain

Two options:

- **Keep your Piefed domain, add path routes.** Recommended if you don't
  want a separate subdomain just for the API and image endpoints. Add
  `/api/v3/*` and `/pictrs/image*` as routes on your existing domain that
  forward to the proxy container, while everything else continues going
  to Piefed as normal.

- **Give your frontend its own subdomain.** Your frontend (mlmym or
  whatever you're using) should live on its own subdomain regardless,
  since it renders full pages for real visitors and needs its own public
  HTTPS. A common pattern is `old.yourdomain.com`, with the API and image
  routes added to that same subdomain's block rather than your main
  domain. This keeps your main domain's behavior completely unchanged,
  and nothing on it accidentally looks like a Lemmy server to apps or
  federation tooling that might probe it. This is the approach used in
  testing and described below.

Avoid putting the proxy's API routes on your bare Piefed domain if you
can. A Lemmy app pointed at that domain would log in and see instance
metadata, then fail confusingly on almost everything else, since only
the frontend subdomain actually has the full proxy surface behind it.

## 2. Run the proxy and the frontend

```
docker build -t lemmybeproxy .

docker run -d \
  --name lemmybeproxy \
  --restart unless-stopped \
  -p 127.0.0.1:8050:8080 \
  -e BACKEND_TYPE=piefed \
  -e BACKEND_INSTANCE=yourdomain.com \
  -e FRONTEND_VERSION=0.19 \
  --label "com.centurylinklabs.watchtower.enable=true" \
  lemmybeproxy
```

Then run your frontend. This example uses mlmym:

```
docker run -d \
  --name mlmym \
  --restart unless-stopped \
  -p 127.0.0.1:8051:8080 \
  -e LEMMY_DOMAIN=old.yourdomain.com \
  --add-host old.yourdomain.com:host-gateway \
  --label "com.centurylinklabs.watchtower.enable=true" \
  code.mschae23.de/mschae23/mlmym:latest
```

Two details worth explaining:

`LEMMY_DOMAIN` should be the frontend's own public subdomain, not the
proxy's internal port and not your Piefed domain. mlmym uses this both
to build links and to make its own outbound API calls, so it needs to be
a real, publicly resolvable hostname Caddy routes correctly.

`--add-host old.yourdomain.com:host-gateway` matters more than it looks.
Without it, the frontend container reaches its own API calls by going
out to the public internet and back in, since it's calling its own
public domain name — a genuinely slow round trip added to every page
load. This flag makes the container resolve that domain straight to the
host machine instead, so the connection never leaves the server. Caddy
still sees the same hostname and routes it the same way; only the
network path gets shorter.

## 3. Configure your reverse proxy

Send `/api/v3/*` and `/pictrs/image*` to the proxy container, everything
else to the frontend container.

### Nginx

Add this to the server block for your frontend subdomain, above the
existing `location /` block, since Nginx matches by specificity and
these need priority:

```
location /api/v3/ {
        proxy_pass http://127.0.0.1:8050;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-Proto https;
}

location /pictrs/image {
        proxy_pass http://127.0.0.1:8050;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-Proto https;
}

location / {
        proxy_pass http://127.0.0.1:8051;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-Proto https;
}
```

Adjust the ports to match however you published the containers, then:

```
nginx -t && systemctl reload nginx
```

### Caddy

Add this to your frontend subdomain's block, using `handle` so these
routes take priority over the frontend's own catch-all:

```
old.yourdomain.com {
        handle /api/v3/* {
                reverse_proxy localhost:8050 {
                        header_up Host {host}
                        header_up X-Real-IP {remote_host}
                        header_up X-Forwarded-Proto https
                }
        }
        handle /pictrs/image* {
                reverse_proxy localhost:8050 {
                        header_up Host {host}
                        header_up X-Real-IP {remote_host}
                        header_up X-Forwarded-Proto https
                }
        }
        reverse_proxy localhost:8051 {
                header_up Host {host}
                header_up X-Real-IP {remote_host}
                header_up X-Forwarded-Proto https
        }
}
```

Adjust the ports if published differently, then:

```
caddy validate --config /etc/caddy/Caddyfile && systemctl reload caddy
```

If your reverse proxy runs inside Docker rather than as a host service,
reload through its container instead — e.g. `docker exec <container>
caddy reload --config /etc/caddy/Caddyfile` or `docker exec <container>
nginx -s reload`.

## 4. Verify it

```
curl -s "https://old.yourdomain.com/api/v3/site" | head -c 200
```

Should return real site data from your backend, translated into Lemmy's
response shape. If it does, load the frontend's URL in a browser and
confirm it renders.

## 5. Automatic updates (optional but recommended)

Both containers above are labeled for Watchtower. If you don't already
run it on this host:

```
docker run -d \
  --name watchtower \
  --restart unless-stopped \
  -e DOCKER_API_VERSION=<your docker daemon's api version> \
  -v /var/run/docker.sock:/var/run/docker.sock \
  containrrr/watchtower \
  --label-enable --cleanup --interval 86400
```

Find your daemon's API version with `docker version --format
'{{.Server.APIVersion}}'`. Setting it explicitly avoids a real
compatibility issue in some Watchtower builds, where the bundled Docker
client defaults to an old API version and refuses to talk to a newer
daemon.

With `--label-enable`, Watchtower only touches labeled containers, so it
won't affect anything else running on the host.
