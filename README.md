# Mauritius Emergency Services Go

Mauritius Emergency Services is an API that provides a list of emergency phone services in Mauritius.

MES Go is written in Go language using the Gin framework.

## Installation

To run the API locally, you need to have GO binaries installed along side the following dependencies:

`go get -u github.com/gin-gonic/gin`

## Base URL

`https://mesg.th3pl4gu3.com`

## Usage

All responses will have the form

```json
{
    "services": "A JSON list of all available emergency services",
    "message" : "A description of what happened",
    "success" : "A boolean value to whether request succeeded or failed"
}
```

### Get services by language

***Available Languages***

English: `GET /en/{other parameters}`

French: `GET /fr/{other parameters}`

### List all services

**Definition**
`GET /en/services`

***Responses***

- `200 OK` on success

```json
{
    "services": [
        {
            "identifier": "police-direct-line-1",
            "name": "Police Direct Line 1",
            "type": "SECURITY",
            "icon": "https://www.icon.com/link_to_icon",
            "number": 999
        },
        {
            "identifier": "police-direct-line-2",
            "name": "Police Direct Line 2",
            "type": "SECURITY",
            "icon": "https://www.icon.com/link_to_icon",
            "number": 112
        }
    ],
    "message" : "",
    "success" : true
}
```

- `404 NOT FOUND` on failure

```json
{
    "services": [],
    "message" : "Wrong routes used. Please read the docs on https://github.com/mervin16/Mauritius-Emergency-Services-Api-Go",
    "success" : false
}
```

### Get a single service

**Definition**
`GET /en/service/{identifier}`

***Responses***

- `200 OK` on success

```json
{
    "services": [
        {
            "identifier": "police-direct-line-1",
            "name": "Police Direct Line 1",
            "type": "SECURITY",
            "icon": "https://www.icon.com/link_to_icon",
            "number": 999
        }
    ],
    "message" : "",
    "success" : true
}
```

- `404 NOT FOUND` if no services found

```json
{
    "services": [],
    "message" : "No services found under id {ID}",
    "success" : false
}
```

### Search for services

**Definition**
`GET /en/services/search?query=searchQuery`

***Responses***

- `200 OK` on success

```json
{
    "services": [
        {
            "identifier": "police-direct-line-1",
            "name": "Police Direct Line 1",
            "type": "SECURITY",
            "icon": "https://www.icon.com/link_to_icon",
            "number": 999
        },
        {
            "identifier": "police-direct-line-2",
            "name": "Police Direct Line 2",
            "type": "SECURITY",
            "icon": "https://www.icon.com/link_to_icon",
            "number": 112
        }
    ],
    "message" : "",
    "success" : true
}
```

## License

```
Copyright Mervin Hemaraju

```
