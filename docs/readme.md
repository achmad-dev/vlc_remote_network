## vlc_remote_network
[demo.webm](https://github.com/user-attachments/assets/fb296376-3c0d-43a1-bfe0-a327c7d88bd1)
![demo](https://github.com/user-attachments/assets/b9e8ec97-9fdc-400b-bea9-3b041e892fc7)

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
