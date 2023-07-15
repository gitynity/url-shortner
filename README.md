# go-web-Projects.url-shortner

<img width="794" alt="Screenshot 2023-07-16 at 5 11 40 AM" src="https://github.com/gitynity/go-web-Projects/assets/23361845/29e0d851-456f-4af5-96fa-4b97c2b72125">

----------------------------------
## Dependencies:
- make sure you have mysql server installed on your machine.
- To install mysql server on mac: `brew install mysql`

## Running the server:
- `git clone https://github.com/gitynity/go-web-Projects.git`
- `cd go-web-Projects/URLShortner && make build`
- `./bin/url-shortner`

## Requests:
- POST: `curl -X POST -i -v 'localhost:8080/add-url?long_url=https://go.dev/play/'`
- GET: `curl -i -v 'localhost:8080/get-long-url?short_code=rJQgckya'`
- GET: `curl -i -v 'localhost:8080/get-short-url?long_url=https://goplay.tools/'`
- DELETE: `curl -X DELETE -i -v 'localhost:8080/remove-url?long_url=https://go.dev/play/'`
