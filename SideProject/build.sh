
#!/bin/bash

module=${1}

if [ "${module}" = "restapi" ];
then
	cd cmd/${module}
	go build -o ../../bin/${module} restapi.go
else
	echo "nope"
fi
