# Run goback with systemd  

## Setup  
+ build goback   
```
go build cmd/server/main.go 
```
+ fix `PATH_TO_BUILT_MAIN` and `PATH_TO_PROD_CONFIG` in `goback.service` and save it as /lib/systemd/system/goback.service  
+ sudo systemctl enable goback.service  

## Start  
```
sudo systemctl start goback
```

## Stop  
```
sudo service goback stop
```

## View logs
```
journalctl -u goback
```
