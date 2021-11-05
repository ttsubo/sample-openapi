# Practice OpenAPI using OpenAPI Generator

This page indicates how to use [OpenAPI Generator](https://openapi-generator.tech) as initial practice

## Initial confirming

### Step0) Installing OpenAPI Generator on MacOS

```
$ brew install openapi-generator
```
```
% openapi-generator version
5.3.0
```

### Step1) Preparing openapi.yml

```
openapi: 3.0.1

info:
  title: "sample API"
  description: "This is sample."
  version: "1.0.0"

paths:
  /pets/{id}:
    get:
      description: Returns a single pet
      parameters:
        - name: id
          in: path
          description: ID of pet to fetch
          required: true
          schema:
            type: integer
            format: int64
      responses:
        200:
          description: successful pet response
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pets'

components:
  schemas:
    Pets:
      type: object
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
          example: doggie
        status:
          type: string
          description: pet status in the store
          enum:
          - available
          - pending
          - sold
```
### Step2) Executing openapi-generator for server
```
% openapi-generator generate -i openapi.yml -g go-server -o ./server
[main] INFO  o.o.codegen.DefaultGenerator - Generating with dryRun=false
[main] INFO  o.o.c.ignore.CodegenIgnoreProcessor - Output directory (/Users/ttsubo/source/sample-openapi/./server) does not exist, or is inaccessible. No file (.openapi-generator-ignore) will be evaluated.
[main] INFO  o.o.codegen.DefaultGenerator - OpenAPI Generator: go-server (server)
[main] INFO  o.o.codegen.DefaultGenerator - Generator 'go-server' is considered stable.
[main] INFO  o.o.c.languages.AbstractGoCodegen - Environment variable GO_POST_PROCESS_FILE not defined so Go code may not be properly formatted. To define it, try `export GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"` (Linux/Mac)
[main] INFO  o.o.c.languages.AbstractGoCodegen - NOTE: To enable file post-processing, 'enablePostProcessFile' must be set to `true` (--enable-post-process-file for CLI).
[main] INFO  o.o.codegen.utils.URLPathUtils - 'host' (OAS 2.0) or 'servers' (OAS 3.0) not defined in the spec. Default to [http://localhost] for server URL [http://localhost/]
[main] INFO  o.o.codegen.utils.URLPathUtils - 'host' (OAS 2.0) or 'servers' (OAS 3.0) not defined in the spec. Default to [http://localhost] for server URL [http://localhost/]
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/model_pets.go
[main] WARN  o.o.codegen.DefaultCodegen - Empty operationId found for path: get /pets/{id}. Renamed to auto-generated operationId: petsIdGet
[main] INFO  o.o.codegen.utils.URLPathUtils - 'host' (OAS 2.0) or 'servers' (OAS 3.0) not defined in the spec. Default to [http://localhost] for server URL [http://localhost/]
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/api_default.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/api_default_service.go
[main] INFO  o.o.codegen.utils.URLPathUtils - 'host' (OAS 2.0) or 'servers' (OAS 3.0) not defined in the spec. Default to [http://localhost] for server URL [http://localhost/]
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/api/openapi.yaml
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/main.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/Dockerfile
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go.mod
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/routers.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/logger.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/impl.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/helpers.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/api.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/go/error.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/README.md
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/.openapi-generator-ignore
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/.openapi-generator/VERSION
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/ttsubo/source/sample-openapi/./server/.openapi-generator/FILES
################################################################################
# Thanks for using OpenAPI Generator.                                          #
# Please consider donation to help us maintain this project 🙏                 #
# https://opencollective.com/openapi_generator/donate                          #
################################################################################
```

### Step3) Try to confirm how it works
After customizing go.mod environment, try to run openapi server

```
% go run main.go      
2021/11/05 10:25:05 Server started
```

When accessing to openapi server from another terminal ...
```
% curl http://localhost:8080/pets/1
"PetsIdGet method not implemented"
```
```
% curl http://localhost:8080/pets/1 -w '%{http_code}\n'
"PetsIdGet method not implemented"
501
```
