## vlc_remote_network

[watch video](../assets/vlc_remote_network_demo.webm)

## How to use

### Install

go to network folder and run

```bash
go mod tidy
```

and then Make sure that your computer’s firewall allows incoming traffic on port 8080 

```bash
sudo ufw allow 8080/tcp
```

### Run

run

```bash
cd cmd
go run main.go
```

### To get local ip address

```bash
ifconfig | grep netmask
```

also make sure that your router don't have ap isolation

## To use flutter app

first build the flutter app to apk

```bash
flutter build apk
```

then install the apk on your phone

and insert the ip address of your computer in the ip field like this
eg: `http://192.168.1.xxx:8080/vlc`

## For note

this is working on ubuntu 22.04 also make sure that the vlc is running on top focused window when the go app running
