# PayFast Client for Go

A client library written in Go (golang) for the [PayFast](https://www.payfast.co.za/) payment gateway which uses no
external dependencies.

## Installation
```shell script
go get -u github.com/huysamen/payfast-go
```

## Quickstart
Create a new client default client to start accessing your PayFast account.
```go
client, err := api.Default()
``` 

This will create a client with the default settings.  It also expects the following environment variables to be set:
```
PAYFAST_MERCHANT_ID=[your payfast merchant id]
PAYFAST_MERCHANT_PASSPHRASE=[your payfast merchant passphrase]
```

You can also create a more configurable client which accepts an `*http.Client` as well as the merchant ID and passphrase.

```go
client, err := api.New(123, "passphrase", httpClient)
```

## Examples and How To

### Health checks

#### API health check
```go
ok, err := client.Health.Ping()
```

### Subscriptions

#### Fetch subscription
```go
sub, err := client.Subscriptions.Fetch("subscription-token")
```

#### Update subscription
An example of how to use the nullable types:
```go
sub, err := api.Subscriptions.Update(
	"subscription-token",
	subscriptions.UpdateSubscriptionReq{
		Cycles:    types.Numeric{},     // default instance treated as nil and ignored
		Frequency: types.NewNumeric(types.Annual),
		//RunDate: types.Time{},      // excluded field treated as nil and ignored
		Amount: types.NewNumeric(123),
	},
)
```