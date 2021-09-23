# Coin converter api

## Technologies used in this api ##

1. gorm
1. gorilla/mux
1. vipernev
1. Some packages installed indirectly.

## How to run this app ##

1. install and configure local mysql server.
1. create a database and name it as you wish.
1. set .env variables to connect local mysql database.
    1. admin `mysql user`
    1. password `mysql user password`
    1. database `database name`
    1. host `usually tcp(1277.0.0.1:3306)`
1. run the main.go and hit endpoints in your rest client app.

## Avaiables endpoints to use ##

1. `/exchange/<amount>/<from>/<to>/<rate>`
1. `/exchange/conversions`
1. `/exchange/results`