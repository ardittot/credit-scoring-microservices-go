# credit-scoring-microservices-go
Build credit scoring microservices app with Go

Install required packages
```
go get -u github.com/gin-gonic/gin
```

Compile & run
```
go build
./credit-scoring-go
```

API Specs
```
curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d @./json/example.json http://localhost:8000/crs
```
