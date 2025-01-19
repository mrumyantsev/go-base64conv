# base64conv

base64conv is a Go library for encoding and decoding numbers to base64 format. It might be used in certain cases, such as generating URLs for end network resources.

## Prerequisites

- [Go](https://go.dev/dl) >= v1.15

## Installation

Download/update the library.

```
go get -u "github.com/mrumyantsev/base64conv-go"
```

## Usage

Encode **int64 number** id to **base64 string** URL.

``` Go
base64url := base64conv.ItobRawUrl(id)
```

Decode **base64 string** URL to **int64 number** id.

``` Go
id, err := base64conv.BtoiRawUrl(base64url)
if err != nil {
    return err
}
```

## Testing

Run the unit-tests.

```
make test
```
