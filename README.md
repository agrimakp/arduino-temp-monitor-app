# arduino-temp-monitor-app
View temperature, and humidity readings captured from Arduino on a mobile App

An innovative project utilizing the DHT11 sensor module to create a real-time environmental monitoring system. This system continuously tracks and displays temperature and humidity levels in the surrounding environment, providing accurate and up-to-date data. 

## Goals

- [ ] Read temperature and humidity with DHT11 module
- [ ] Send readings to backend via HTTP
- [ ] Store time series readings data
- [ ] Query latest readings for a source
- [ ] Query analytical data for statistics over time
- [ ] Mobile and Web apps for convenience

## Database

### Table structure

| Column Name | Type     |
| ----------- | -------- |
| Time        | DATETIME |
| Temperature | DOUBLE   |
| Humidity    | DOUBLE   |
| Source      | VARCHAR  |

### Example Data

| Time                | Temperature | Humidity | Source |
| ------------------- | ----------- | -------- | ------ |
| 2024-04-01-01:12:33 | 23.3        | 57.5     | Hall   |

### Access patterns

**1. Latest reading per source**

```
 SELECT *
FROM   readings
WHERE  source = '<source>'
ORDER  BY datetime DESC
LIMIT  1;  
```

## Development

### Start database

https://hub.docker.com/_/postgres

```sh
docker run -p 5432:5432  -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```

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
