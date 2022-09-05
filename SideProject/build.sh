
#!/bin/bash

module=${1}

if [ "${module}" = "httpserver" ];
then
	cd cmd/${module}
	go build -o ../../bin/${module} main.go
else
	echo "nope"
fi
