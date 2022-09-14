Hubspot REST client
===========

How to use?
------

First of initiate the client with stateful oauth credentials. E.g., in init

```go
func init() {
	hubspotclient.Initiate(hubspot.Config{
		ClientID:     "<your client id>",
		ClientSecret: "<your client secret>",
		RedirectUrl:  "<guess what: redirect url>",
		Scopes:       []string{"<sco", "pes>"},
	})
}
```

After the grant & token exchange (it's pretty straight forward), create a client and... use it.

```go
hs, _ := hubspotclient.CreateInstance(ctx, token)
_ = hs.CreateTicket(request) // note, the access token may automatically refresh here if it is expired
token, _ := hs.Token() // here you will always have an up to date access token
```
