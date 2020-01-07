# profile

![](https://github.com/26huitailang/profile/workflows/Go/badge.svg)

## Config

In the `$PROJECT/config` package. Override configs according to the following order:
- default.yml
- `$GO_ENV`.yml(a env variable: test/develop)
- local.yml (not in the repo, you should create one)

## Packages

- echo: web server
- viper: config
- mongodb: data storage

## Todo

- [ ] Swagger 不是很好配置🙅‍