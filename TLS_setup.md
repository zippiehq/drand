### TLS setup: Nginx with Let's Encrypt
Such a setup greatly simplify TLS management issues (renewal of certificates, etc). We provide here the
minimum setup using [nginx](https://www.nginx.com/) and [certbot](https://certbot.eff.org/lets-encrypt/) - make sure you have both binaries installed with the latest version.

+ First, add an entry in the nginx configuration for drand:
```bash
# /etc/nginx/sites-available/default
server {
  server_name drand.nikkolasg.xyz;
  listen 443 ssl;

  location / {
    grpc_pass grpc://localhost:8080;
  }
  location /api/ {
    proxy_pass http://localhost:8080;
    proxy_set_header Host $host;
  }
}  
```
**Note**: you can change
1. the port on which you want drand to be accessible by changing the line `listen 443 ssl` to use any port.
2. the port on which the drand binary will listen locally by changing the line `proxy_pass http://localhost:8080; ` and ` grpc_pass grpc://localhost:8080;` to use any local port.

+ Run certbot to get a TLS certificate:
```bash
sudo certbot --nginx
```

+ **Running** drand now requires to add the following options:
```bash
drand start --tls-disable --listen 127.0.0.1:8080
```

The `--listen` flag tells drand to listen on the given address instead of the public address generated during the setup phase (see below).
