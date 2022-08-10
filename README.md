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
GOLang server is launched as systemd daemon (service). Path to service: `/etc/systemd/system/academy-go-server`
It should be COMPILED with `go build` and placed into directory `/home/academy/VRisingAcademySite` as `main`
Possible to change the directory by editing `ExecStart` in systemd service

### Commands related to go server
To enable (Add go server as systemd service): systemctl enable academy-go-server
To disable (Disable go server as systemd service): systemctl disable academy-go-server
To start GO server: systemctl start academy-go-server
To stop GO server: systemctl stop academy-go-server
