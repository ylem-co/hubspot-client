Hubspot REST client
===========

![GitHub branch check runs](https://img.shields.io/github/check-runs/datamin-io/hubspot-client/main?color=green)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/datamin-io/hubspot-client?color=blue)
<a href="https://github.com/datamin-io/hubspot-client?tab=Apache-2.0-1-ov-file">![Static Badge](https://img.shields.io/badge/license-Apache%202.0-blue)</a>
<a href="https://datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/website-datamin.io-blue)</a>
<a href="https://docs.datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/documentation-docs.datamin.io-blue)</a>
<a href="https://join.slack.com/t/datamincommunity/shared_invite/zt-2nawzl6h0-qqJ0j7Vx_AEHfnB45xJg2Q" target="_blank">![Static Badge](https://img.shields.io/badge/community-join%20Slack-blue)</a>

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
