# Go web application boilerplate

This is a structre for a go web application, inspired by [project-layout](https://github.com/golang-standards/project-layout). Please feel free to open issues or pull request to improve this boilerplate.

Use `make init` to install githooks and download dependencies.

The default http router is [gorilla mux](https://github.com/gorilla/mux).

The logging framework is [Zap](https://github.com/uber-go/zap) and can be replaced in `/pkg/infra/logger.go`.

## Central folder structure

| folder | description |
| ------- | ----------- |
| /api | OpenAPI/Swagger specs, JSON schema files, protocol definition files. |
| /cmd | Main applications for this project. |
| /config | Configuration file templates or default configs. |
| /pkg | Library code. More information inside this folder. |
| /scripts | Scripts to perform various build, install, analysis, etc operations. |
| /test | Additional external test apps and test data. |
| /vendor | Application dependencies. |
| /web | Web application specific components: static web assets, server side templates and SPAs. |
