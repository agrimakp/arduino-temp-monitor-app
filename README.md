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


