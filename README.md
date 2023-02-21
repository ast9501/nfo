# RAN nfo

## Usage
### Build app
>You may need to install `make` first
```
make all
```

### Run nfo
```
./nfo -c config
```
- `-c`: path of config directory

### Test
* Health check
```
curl https://<IP>:<port>/nfo/v1/
```

## Support
### Swagger api
```
http://<IP>:<port>/swagger/index.html
```

### Project Archtecture
1. [Golang clean arch](https://github.com/bxcodec/go-clean-arch)
2. [Golang project layout](https://github.com/golang-standards/project-layout)
