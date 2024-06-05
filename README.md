# urlShortner

URL Shortening: Generate a short URL from a long URL.
URL Redirect: Retrieve the original long URL from a short URL.
Domain Count: Return the count of hits per domain.


## Build

To build the application, run:

```bash
make build
```

## Run

To run the application locally, use:

```bash
make run
```

## Test

To run the tests, execute:

```bash
make test
```

### How to use?

To use this application, first start the service locally by running:

```bash
make run
```

Then, use the following endpoint:

#### Shorten the URL

```bash
Endpoint: http://localhost:8020/v1/
```

Make a POST request to the above endpoint with the following payload:

```bash
{
    "url": "www.manoj.com"
}
```

The request will return a shortened URL like this:"9cfb054a4b9e7a"

```bash
http://localhost:8020/v1/url/9cfb054a4b9e7a
```

### Retrive origanal URL

Paste the received shortened URL into any browser to be redirected to the original URL.

### Metrics

For getting the top 3 domain names:

```bash
http://localhost:8020/v1/getmetrics
```
