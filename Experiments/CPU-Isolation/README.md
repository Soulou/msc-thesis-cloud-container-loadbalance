# CPU Isolation experiment

Small CPU burning software written in go to test Linux Container CPU Cgroup isolation features

* Build executable: `go build`
* Build image: `docke build -t 'cranfield/cpu-isolation' .`
* Run image `docker run cranfield/cpu-isolation -nb-cpus=#`

With # the number of CPU you want the container to use.
The container will stop cleanly with any of SIGINT, SIGSTOP, SIGTERM or SIGQUIT
