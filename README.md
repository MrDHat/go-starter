![GitHub](https://img.shields.io/github/license/mrdhat/go-starter)
[![DeepSource](https://deepsource.io/gh/MrDHat/go-starter.svg/?label=active+issues&show_trend=true&token=DubK1nIGgMWl5FQNtxhx9ux0)](https://deepsource.io/gh/MrDHat/go-starter/?ref=repository-badge)
# Go starter template
A starter template for golang based web projects. The template follows convention over configuration pattern with focus on performance & ease of use


## Notes about dependencies

The following packages are used to build the template:
- [Echo Web Framework](https://github.com/labstack/echo) & a bunch of echo plugins
- [Zap Logging Library](https://github.com/uber-go/zap)
- [Cleanenv Config Library](https://github.com/ilyakaznacheev/cleanenv)

### Generate Swagger Docs
- Install [Swag CLI](https://github.com/swaggo/swag#getting-started)
- Run `swag init`
> Best to add it as a part of your CI pipeline

### Hot Reload
- If you want to enable hot reload, install [air](https://github.com/cosmtrek/air#prefer-installsh)
- Run `air start` to start the api server

## Using the template
>TODO
