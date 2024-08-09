# arduino-temp-monitor-app
View temperature, and humidity readings captured from Arduino on a mobile App

## Goals

- [ ] Read temperature and humidity with DHT11 module
- [ ] Send readings to backend via HTTP
- [ ] Store time series readings data
- [ ] Query latest readings for a source
- [ ] Query analytical data for statistics over time
- [ ] Mobile and Web apps for convenience

## Database

### Table structure (Readings)

| Column Name | Type     |
| ----------- | -------- |
| Time        | DATETIME |
| Location    | VARCHAR  |
| Temperature | DOUBLE   |
| Humidity    | DOUBLE   |

### Example Data

| Time                | Temperature | Humidity | Location |
| ------------------- | ----------- | -------- | ------  |
| 2024-04-01-01:12:33 | 23.3        | 57.5     | Hall    |
| 2024-04-01-01:12:33 | 3.0         | 50.5     | Fridge  |

### Access patterns

**1. Latest reading per source**

```
 SELECT *
FROM   readings
WHERE  location = '<location_name>'
ORDER  BY time DESC
LIMIT  1;  
```


**1. Latest reading per source**

```
 SELECT *
FROM   readings
WHERE  location = '<location_name>'
and time > <from> and time < <to>
ORDER  BY time DESC
LIMIT  1;  
```

## Development

### Start database

https://hub.docker.com/_/postgres

```sh
docker run  -e POSTGRES_PASSWORD=mysecretpassword -d postgres

# Optionally connect to db using psql (enter password mysecretpassword when prompted)
psql -h localhost -U postgres
```

#### Configure database

To start working with our database, we first need to create our table based on designed schema
(See db/schema.sql).

Optionally add some sample data using db/seed.sql

### Start server

```sh
cd server
go run .
```

### Testing

**1. Add Reading**

Send POST request to `/readings/add`

> Adds a new reading to the db, temp=21, humidity=50 for location hall

```sh
curl -X POST http://localhost:8090/readings/add -H 'Content-Type: application/json' -d '{
	"Temperature": 21,
	"Humidity": 50,
	"Source": "hall"
}'
```

**2. Get Latest Reading**

Send GET requesty to '/readings/latest'

```sh
curl http://localhost:8090/readings/latest?source=hall
```
## Development Board

Using ESP32 Plus Development board.

### Setup

1. Configure Arduino IDE 
    https://docs.espressif.com/projects/arduino-esp32/en/latest/installing.html
2. Install Board
    Tools -> Board Manager
3. Select Board
    Tools -> Board -> ESP32 Dev
4. Select Port

In linux there might be permission issue on the device port
Run the following command

```sh
# replace /ttyUSB0 with your device port name
sudo chmod a+rw /dev/ttyUSB0
```

## Resources

https://go.dev/doc/tutorial/database-access