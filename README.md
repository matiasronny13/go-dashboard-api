# go-dashboard-api

Component | Library
--- | --- 
Web Framework|gin-gonic/gin + HTMX
ORM|gorm.io/driver/postgres
Dependency Injection|google/wire 
Swagger|go-openapi/swag
Database|PostgreSQL

### Build Dependency Injection by using Wire
```
wire .\api\bookmark
wire .\api\notes
```

### Generate Swagger
```
swag init -g .\api\bookmark\handler.go
swag init -g .\api\notes\handler.go
```

### Create Simbolic Link to [lit-web-component](https://github.com/matiasronny13/lit-web-component) build output
```
mklink /J .\html\assets\lit <lit-web-component>\build
```
