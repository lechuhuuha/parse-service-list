# Fedora / Redhat Systemd Service List processor.

This is a module that can extract the running services and their status from 

```bash
systemctl list-units 
```

or 

```bash
ps -eo pid,user,ni,%cpu,%mem,args --sort=-%cpu,-%mem
```

into data structures.



### To use, add as a module:

```
$ go get -u github.com/lechuhuuha/parse-service-list
```

Sample Application in the parse-service-list_test.go

Output

```json
[
{"pid":1546,"user":"root","ni":0,"cpu":0.5,"mem":1.7,"command":"/root/go/bin/gopls"},{"pid":1489,"user":"root","ni":0,"cpu":0.3,"mem":2,"command":"/root/.vscode-server/bin/74f6148eb9ea00507ec113ec51c489d6ffb4b771/node"},{"pid":502,"user":"root","ni":0,"cpu":0.1,"mem":0.9,"command":"/snap/ubuntu-desktop-installer/1151/usr/bin/python3.10"},{"pid":1419,"user":"root","ni":0,"cpu":0.1,"mem":0.8,"command":"/root/.vscode-server/bin/74f6148eb9ea00507ec113ec51c489d6ffb4b771/node"}
]
[
{"serviceName":"apport.service","loaded":"loaded","state":"active","status":"exited","description":"LSB: automatic crash report generation"},{"serviceName":"console-getty.service","loaded":"loaded","state":"active","status":"running","description":"Console Getty"},{"serviceName":"console-setup.service","loaded":"loaded","state":"active","status":"exited","description":"Set console font and keymap"},{"serviceName":"cron.service","loaded":"loaded","state":"active","status":"running","description":"Regular background program processing daemon"}

]
```

