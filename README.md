# url-shortener

A URL shortner microservice

# How to use

## To run main REST component

```
go run cmd/main.go
```

## To get shortened URL

After running main REST component (see above)

Visit following URL in browser of choice:

```
http://localhost:8080/shorten?url=<url_to_be_shortened>
```

## To use new shortened URL

Visit following URL in browser of choice:

\*Replace shortened_url with the URL printed to the browser after running the above command

```
http://localhost:8080/redirect?url=<shortened_url>
```
