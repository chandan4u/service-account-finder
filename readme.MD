# Service Account Finder

## Overview
Service account finder is a command line tool that find service account of organisation of projects. It lists all the service account from organisation projects. 

## Development setup
###### Git clone 
```base
git clone https://github.com/chandan4u/service-account-finder.git
```
###### Install project Go dependencies
```base
go mod vendor
``` 
###### Install project
The go install command behaves almost identically to go build , but instead of leaving the executable 
in the current directory, or a directory specified by the -o flag, it places the executable into 
the $GOPATH/bin directory.
```base
go install .
```
###### Run project locally
```base
 --help
```

## Installation

###### Download and move service account finder go binary into bin folder
```base
sudo cp service-account-finder /usr/local/bin/
```
###### Test deployer command
```base
service-account-finder
```
It should return given below value
```base
Welcome | ServiceAccountFinder is command line tool for finder all service account inside organisation projects.
```

## Usage

###### Help command
To see the all option inside ServiceAccountFinder command line tool.
```base
service-account-finder --help
```

###### Version Command
To see the current version of ServiceAccountFinder command line tool.
```base
service-account-finder --version
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what 
you would like to change.

Please make sure to update tests as appropriate.

## License
[Public Cloud]