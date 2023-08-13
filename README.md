# url-shortener

A URL shortner microservice

## How to use

- Ensure you have PostgreSQL set up on your local machine.
- Replace the values of dbUser, dbName and dbPassword in models/db.go to match your PostgreSQL configuration credentials.

## To run main REST component

```
go run cmd/main.go
```

## To get shortened URL

After running main REST component (see above)

Open a web browser and navigate to following url:

- Remember to replace <url_to_be_shortened> with your URL of choice

```
http://localhost:8080/shorten?url=<url_to_be_shortened>
```

## To use new shortened URL

Open a web browser and navigate to shortened URL received from previous step
