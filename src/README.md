# Google Agenda Plugin for M.A.R.C.E.L.
[https://github.com/Zenika/MARCEL]
 
This plugin displays the next 'n' entries from your agenda.
 
##Attributes
 See "config.json" file


## Setup

### Backend
```
go get github.com/dghubble/oauth1
go get github.com/rs/cors
go get -u github.com/gorilla/mux
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/...
```

### Frontend
...

##Build cross architecture
```
env GOOS=linux GOARCH=arm go build -o ./dist/back/arm_linux
env GOOS=windows GOARCH=arm go build -o ./dist/back/arm_windows
env GOOS=linux GOARCH=amd64 go build -o ./dist/back/amd64_linux
```
List of all GOOS and GOARCH values : https://golang.org/doc/install/source#environment
