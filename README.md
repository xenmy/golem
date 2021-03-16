# golem
## Azure Storage Blob Head Server
You can access blob files via http.
```
http://golem.example.com/{blob path}
ex)
http://golem.example.com/hogehoge/fugafuga.html -> ${AZURE_STORAGE_ACCOUNT}/${AZURE_STORAGE_CONTAINER}hogehoge/fugafuga.html
```

## Environment Variables
|No.|Name|Short Description|Example|
|---|---|---|---|
|1|SERVER_ADDRESS|Listen Address|0.0.0.0|
|2|SERVER_PORT|Listen Port|8080|
|3|AZURE_STORAGE_ACCOUNT|Azure Storage Account Name|xmyblob|
|4|AZURE_STORAGE_ACCOUNTKEY|Azore Storage Account Key||
|5|AZURE_STORAGE_CONTAINER|Container Name|web|

## Limitation
- Not support TLS termination
- Not support index file serving
- Experimental quality
- 
