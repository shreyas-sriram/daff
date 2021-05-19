# daff - Don't Ask For Flags

A Discord bot to check the health of CTF challenges.

CTF participants do not have to pester organizers about the challenge's health :rocket:

> This bot internally makes HTTP requests, hence works only for challenges expecting HTTP requests.

![up](https://github.com/shreyas-sriram/daff/blob/main/docs/up.png)

![down](https://github.com/shreyas-sriram/daff/blob/main/docs/down.png)

## Usage

### Bot command

```
!daff <challenge-name>
```

### Server

```
Usage of ./bin/daff:
  -f string
    	Config file (default "config.yaml")
  -t string
    	Bot Token
```

### Configuration file

#### Format

```
challenges:
  chall-get:
    url: http://localhost:5000/get
    request:
      method: GET
      headers:
        - "Authorization:Bearer foobar"
      cookies:
        - "admin:1"
      body: '{"username":"guest","password":"guest"}'
    response:
      status: 200
```

| Field       | Description                                                                   | Type              |
| ----------- | ----------------------------------------------------------------------------- | ----------------- |
| chall-get   | Name of the challenge                                                         | string            |
| url         | URL of the challenge                                                          | string            |
| method      | HTTP method (GET / POST)                                                      | string            |
| headers     | Headers to be added in the request. Header key and value are separated by `:` | array of strings  |
| cookies     | Cookies to be added in the request. Cookie key and value are separated by `:` | array of strings  |
| body        | Raw body to be added in the request                                           | string            |
| status      | Expected response status                                                      | int               |
