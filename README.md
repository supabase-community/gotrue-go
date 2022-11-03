# gotrue-go

A Golang client library for the [Supabase GoTrue](https://github.com/supabase/gotrue) API.

> ⚠️ Using [`netlify/gotrue`](https://github.com/netlify/gotrue)?
>
> The types in this library assume you are interacting with a Supabase GoTrue server, so it is unlikely to work.

## Project status

This library is a pre-release work in progress. It has not been thoroughly tested, and the API may be subject to breaking changes, and so it should not be used in production.

## Quick start

### Install

```sh
go get github.com/kwoodhouse93/gotrue-go
```

### Usage
```go
package main

import "github.com/kwoodhouse93/gotrue-go"

const (
    projectReference = "<your_supabase_project_reference>"
    apiKey = "<your_supabase_anon_key>"
)

func main() {
    // Initialise client
    client := gotrue.New(
        projectReference,
        apiKey,
    )

    // Log in a user (get access and refresh tokens)
    resp, err := client.Token(gotrue.TokenRequest{
        GrantType: "password",
        Email: "<user_email>",
        Password: "<user_password>",
    })
    if err != nil {
        log.Fatal(err.Error())
    }
    log.Printf("%+v", resp)
}
```

## Options

The client can be customized with the options below.

In all cases, **these functions return a copy of the client**. To use the configured value, you must use the returned client. For example:

```go
client := gotrue.New(
    projectRefernce,
    apiKey,
)

token, err := client.Token(gotrue.TokenRequest{
        GrantType: "password",
        Email: email,
        Password: password,
})
if err != nil {
    // Handle error...
}

authedClient := client.WithToken(
    token.AccessToken,
)
user, err := authedClient.GetUser()
if err != nil {
    // Handle error...
}
```

### WithToken
```go
func (*Client) WithToken(token string) *Client
```

Returns a client that will use the provided token in the `Authorization` header on all requests.

### WithCustomGotrueURL
```go
func (*Client) WithCustomGotrueURL(url string) *Client
```

Returns a client that will use the provided URL instead of `https://<project_ref>.supabase.com/auth/v1/`. This allows you to use the client with your own deployment of the Gotrue server without relying on a Supabase-hosted project.

### WithClient
```go
func (*Client) WithClient(client http.Client) *Client
```

By default, the library uses a default http.Client. If you want to configure your own, pass one in using `WithClient` and it will be used for all requests made with the returned `*gotrue.Client`.
