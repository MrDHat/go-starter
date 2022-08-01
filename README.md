# go-starter
A starter template for golang based web projects. The template follows convention over configuration pattern with focus on performance & ease of use


## Notes about dependencies

The following packages are used to build the template:
- [Echo Web Framework](https://github.com/labstack/echo) & a bunch of echo plugins
- [Zap Logging Library](https://github.com/uber-go/zap)
- [Cleanenv Config Library](https://github.com/ilyakaznacheev/cleanenv)

### Generate Swagger Docs
- Install [Swag CLI](github.com/swaggo/swag/cmd/swag@latest)
- Run `swag init`
> Best to add it as a part of your CI pipeline

### Hot Reload
- If you want to enable hot reload, install [air](https://github.com/cosmtrek/air#prefer-installsh)
- Run `air start` to start the api server

## Using the template
>TODO
