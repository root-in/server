The RootIn app: https://rootin.co.il  
- A web server that runs on Google Cloud Run.  
- Static files are also served from here.  
- The uploaded file is stored to a Google Storage bucket.  
- The form fields are saved into a Google spreadsheet including a link to the uploaded file.  

Running the server locally:
```
go run main.go
```
- To write the result to the google sheet: set an env var `credentials.json` with a service account secret that has access to the google sheet in json 
- Writing the file to the bucket requires a change to the code

Calling the server locally:
```
curl -v \
-F "name=reuven" \
-F "surname=harrison" \
-F "email=my@email.com" \
-F "phone=666" \
-F "file=@/tmp/file" \
http://localhost:8080/register
``````