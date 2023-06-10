### ZoneMapDB
ZoneMap DB is a geo database consisting of consumer and commercial businesses along Interstate I-90 in the northern region of the US.
This data will be used to map geospatial data and provide logistic and innovative support, and applications to businesses worldwide.


### Premliminary Architecture
a Go module serves up backend GeoJson files to be consumed on the front-end by a modern vue js appication....

# NY  PA  OH  90 Express from East to West  

### Structure

- Root
    - go.mod - 90Scanner go module
    - go.work - go workspace folder
    - go.sum - dependency tree
    - main.go - executable file and main function to run application

- scanner
    - go.mod - scanner package
    - go.sum - dependency tree
    - scanner.go - IsPointInsidePolygon function (consumed by root/main.go)

- structs
    - structs.go - defines structs used in application

- data
    - jsonfiles - stores json documents


### Run Application
- from root directory:
- go run main.go
