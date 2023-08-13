# url-shortener

A URL shortner microservice

## How to use

Ensure

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
