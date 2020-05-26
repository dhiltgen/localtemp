# localtemp
Use geolocation data to get temperature data for your location

A pre-built container image can be used with:

```
docker run --rm dhiltgen/localtemp
```

If you want to specify a more exact location, try something like the following
with your lat/lon coordinates.

```
docker run --rm -e LAT=33.9206811 -e LON=-118.3304733 dhiltgen/localtemp
```
