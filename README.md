# vrisingacademysite

## Tables
[Items](https://docs.google.com/spreadsheets/d/1R-Re2Xszgm2UDUvNVrYwZES5eAN9gIJqCEi6FGA4XPo/edit#gid=1956406025)

[Spells](https://docs.google.com/spreadsheets/d/1YAMSUdtcrsG-mNUE8sqeDZMyww9yjaZ2kHGg6CavXxY/edit#gid=30364497)
## Other
```
npm i axios
```

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

### ---- Linux related ----
GOLang server is launched as systemd daemon (service). Path to service: `/etc/systemd/system/academy-go-server`
It should be COMPILED with `go build` and placed into directory `/home/academy/VRisingAcademySite` as `main`
Possible to change the directory by editing `ExecStart` in systemd service

### Commands related to go server
To enable (Add go server as systemd service): systemctl enable academy-go-server
To disable (Disable go server as systemd service): systemctl disable academy-go-server
To start GO server: systemctl start academy-go-server
To stop GO server: systemctl stop academy-go-server
