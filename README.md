# VRising Database

## Dependencies
[![Vue 3](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)](https://vuejs.org/) 
[![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)](https://vitejs.dev/)
[![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)](https://www.nginx.com/)
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)

## Links
- Prod — https://vrising-academy.info/
- Dev — https://dev.vrising-academy.info/
- [Master Sheet VRDB](https://docs.google.com/spreadsheets/d/1R-Re2Xszgm2UDUvNVrYwZES5eAN9gIJqCEi6FGA4XPo/edit#gid=1956406025)

### ---- Linux related ----
## Frontend
- `npm i`
- `npm run build`
- `sudo systemctl restart nginx.service`

Nginx Config
path `/etc/nginx/sites-available`:
- `vrising-academy.info` prod
```
server {
    # WWW redirect
    if ($host ~* www\.(.*)) {
        set $host_without_www $1;
        rewrite ^(.*)$ http://$host_without_www$1 permanent;
    }

    root /home/academy/VRisingAcademySite/frontend/dist;

    index index.html index.htm;
    server_name vrising-academy.info; # managed by Certbot

    error_page 404 /index.html;

    location / {
        try_files $uri $uri/ =404;
    }

    location /api {
        proxy_pass http://localhost:8087;
    }

    error_log /var/log/nginx/nginx-error.log;
    access_log /var/log/nginx/nginx-access.log;


    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/vrising-academy.info/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/vrising-academy.info/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot


}

server {
    if ($host = vrising-academy.info) {
        return 301 https://$host$request_uri;
        } # managed by Certbot


        listen 80 ;
        listen [::]:80 ;
        server_name vrising-academy.info;
        return 404; # managed by Certbot


}
```

`dev.vrising-academy.info` dev
```
server {
    # WWW redirect
    if ($host ~* www\.(.*)) {
        set $host_without_www $1;
        rewrite ^(.*)$ http://$host_without_www$1 permanent;
    }

    root /home/academy/VRisingAcademySite-dev/frontend/dist;

    index index.html index.htm;
    server_name dev.vrising-academy.info; # managed by Certbot

    error_page 404 /index.html;

    location / {
        try_files $uri $uri/ =404;
    }

    location /api {
        proxy_pass http://localhost:4047;
    }

    error_log /var/log/nginx/nginx-error.log;
    access_log /var/log/nginx/nginx-access.log;


    listen [::]:443 ssl; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/vrising-academy.info/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/vrising-academy.info/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}

server {
    if ($host = dev.vrising-academy.info) {
        return 301 https://$host$request_uri;
        } # managed by Certbot


        server_name dev.vrising-academy.info;

        listen 80 ;
        listen [::]:80 ;
        return 404; # managed by Certbot


}
```
## Backend
- `go build main.go`
- `systemctl restart academy-go-server` prod
- `systemctl restart academy-go-server-dev` dev

GOLang server is launched as systemd daemon (service). Path to service: `/etc/systemd/system/academy-go-server`
It should be COMPILED with `go build` and placed into directory `/home/academy/VRisingAcademySite` as `main`
Possible to change the directory by editing `ExecStart` in systemd service

### Commands related to go server
To enable (Add go server as systemd service): systemctl enable academy-go-server
To disable (Disable go server as systemd service): systemctl disable academy-go-server
To start GO server: systemctl start academy-go-server
To stop GO server: systemctl stop academy-go-server

## Database
- `sudo su postgrest`
- `psql`
- `\c vrisingdb` prod
- `\c vrisingdb_dev` dev
