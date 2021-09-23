# Converter coin api

## Technologies used in this api ##

1. gorm
1. gorilla/mux
1. vipernev
1. Some packages installed indirectly.

## How to run this app ##

1. install and configure local mysql database.
1. create a database called `conversiondb`
1. set .env variables to local mysql database.
1. run the main.go and hit endpoints in your rest client app.

## Avaiables endpoints to use ##

1. `/exchange/<amount>/<from>/<to>/<rate>`
1. `/exchange/conversions`
1. `/exchange/results`