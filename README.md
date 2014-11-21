GoSensor
============================
A simple monitoring program

## Build
```
cd script
bash build.sh
```

## Installation

### On Ubuntu/Debian
``` sh
dpkg -i gosensor_{version}_{arch}.deb
```

### On CentOS/RHEL/Fedora
``` sh
yum localinstall gosensor-{version}.{arch}.rpm
```

## Usage

### Show help
```
gosensor -h

Usage of gosensor:
  -debug=false: enable debug log
  -error=true: enable error log
  -info=true: enable info log
  -logfile="": logging to file
  -rc="/etc/gosensor/monitor.json": project config file path
  -warning=true: enable warning log
```

## Run
### Debug
``` sh
gosensor -debug=true -rc="/path/to/your/monitor.json"
```
### Run on production enviroment
``` sh
gosensor -rc="/path/to/your/monitor.json" -logFile="/var/log/gosensor.log"

# disable info log
gosensor -info=false -rc="/path/to/your/monitor.json" -logFile="/var/log/gosensor.log" 

# Running on background
gosensor -info=false -rc="/path/to/your/monitor.json" -logFile="/var/log/gosensor.log" &
```

## Configuration
GoSensor read default Configuration from `/etc/gosensor/monitor.json`,
but you can use `-rc` arguments to specify other configuration file.

Read more about [monitor.json.example](https://github.com/lowstz/gosensor/blob/master/monitor.json.example)

## License
Copyright (c) 2013 Zhanpeng Chen

[MIT License](https://github.com/lowstz/gosensor/blob/master/LICENSE)