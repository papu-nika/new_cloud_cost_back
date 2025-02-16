# new_cloud_cost_back

## Update


### Schema Update
```
oapi-codegen -package api -generate types,gin,spec,models -o api/api.gen.go openapi.yaml
swag init -g main.go
```


## Setup

### Swagger Setup

```
curl -L https://github.com/swagger-api/swagger-ui/archive/refs/tags/v4.15.5.zip -o swagger-ui.zip

unzip swagger-ui.zip -d docs/
mv docs/swagger-xx.xx docs/swager
```