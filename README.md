# Rest Countries Go
Get information about countries via a RESTful API made by Go. 

_**Build:**_ 
- Clone repository.
- Edit .env file and main.go to change `@host` var.
- `go mod download`
- `swag init .`
- `CGO_ENABLED=0 go build -ldflags "-s -w" .`

**_Usage:_** `./rest-countries`

**_Demo:_** https://rest-countries-go.herokuapp.com/


<hr>

## Endpoints:

| Method  | Route  |
|---|---|
| GET  | /api/v1/all  |
| GET  | /api/v1/name/:name  |
| GET  | /api/v1/code/:code  |
| GET  | /api/v1/currency/:currency  |
| GET  | /api/v1/language/:language  |
| GET  | /api/v1/capital/:capital  |
| GET  | /api/v1/calling/:code  |
| GET  | /api/v1/region/:region  |
| GET  | /api/v1/regional-block/:block  |


### **Note:** Look at swagger docs for more information.
