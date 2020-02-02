# Weather microservice
[![Build Status](https://travis-ci.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw.svg?branch=master)](https://travis-ci.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw)

| Resource                 | Query params| Description|
|:-------------------------|:------------|:-----------|
| `api/v1/weather/current` | **q**: string, required=true, example=warsaw,london; **page**: int, required=false, default=1; **count**: int, required=false, default=20 |returns a list of weather items for given city/country names separated by commas, optionally with pagination|
                                                                                                                    
                                                                                                             

## Setup

Copy example config file to .env and set the values in it.

```shell
$ cp .env.dist .env
```

Build & run

```shell
$ make prod
```

or if you want to debug

```shell
$ make dev
```

## Usage
Example request
```shell
$ curl http://localhost:8080/api/v1/weather/current?q=warsaw
```
With pagination
```shell
$ curl http://localhost:8080/api/v1/weather/current?q=warsaw
london,paris,praga,york,amsterdam,budapest,cracow,phoenix,
columbus?limit=5&page=2
```

## Available configuration
```.env
# Http port on which server will listen
HTTP_PORT=8080

# Debugger port
DEBUG_PORT=40000

# Open weather API
OPEN_WEATHER_API_KEY=api-key
OPEN_WEATHER_API_BASE_URL=https://api.openweathermap.org/data/2.5/

# Cache item expiration time
CACHE_EXPIRATION=30
# After what time remove expired items
CACHE_INTERVAL=60
```
