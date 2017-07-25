# Google Agenda Plugin for M.A.R.C.E.L.
[https://github.com/Zenika/MARCEL]
 
This plugin displays the next 'n' entries from your agenda.
 
##Attributes
 See "description.json" file


## Setup

### Backend
 // in the  src/back/ folder

 #### 1. Setup of the development environment
 ```
  go get github.com/dghubble/oauth1
  go get github.com/rs/cors
  go get github.com/gorilla/mux
  go get google.golang.org/api/calendar/v3
  go get golang.org/x/oauth2/...
 ```
 
 #### 2. Build binary file and the docker image 
 ```
  make
 ```
 
 #### 3.Pull the docker image on public Docker hub
 ```
  make push
 ```
 or
 ```
  docker pull marcel_agenda
 ```

**Build cross architecture**
Please, refer to  official documentation for all GOOS and GOARCH values : https://golang.org/doc/install/source#environment

 #### Run the Docker image
```
docker run --name "marcel_agenda" -p 8080:8080 -v "/home/gwennael.buchet/PROJETS/gopath/src/github.com/Zenika/marcel-plugin-calendar/dist/data":"/tmp" --env="GOOGLE_API_KEY_FILE=MARCEL-6711c4e89cff.json" --env="MARCEL_AGENDA_ID=zenika.com_h6tsmhv6iqs0e3l0vrugi6tbfg@group.calendar.google.com" marcel_agenda
```
### Frontend
...
