# sendsms

A simple command line to send SMS over a selected provider.

Right now, and because it was my primary need, only OVH provider is implemented.
It should be easy to add more provider. The code is designed with that in mind.

sendsms use a configuration file to define profiles. A profile is the way to use
multiple credentials for the same provider, in instance.

## Configuration

Configuration sample (config.yml):

```YAML
---
app:
  logLevel: "INFO"

profiles:
  default:
    provider:              "ovh"
    providerConfig:
      api:
        location:          "ovh-eu"
        appKey:            "azertyuiop"
        appSecret:         "qsdfghjklm"
        consumerKey:       "wxcvbn"
        servicename:       "aqwokn"
      smsOptions:
        sender:            "MyCorp"
      smsOptionsCaps:
        nostopclause:      "noStopClause"
        servicename:       "serviceName"
        senderforresponse: "senderForResponse"
```

## Usage Examples

Send a message to a phone number:

```BASH
sendsms send --phone +330612345678 --message "Hello my friend !"
```

Add other phone number:

```BASH
sendsms send --phone +330612345678 --phone +330698547621 --message "Hello my friend !"
```

Read message from stdin:

```BASH
echo "Hello my friend !" | sendsms send --phone +330612345678 --stdin
```

List available providers:

```BASH
sendsms provider list
```

Get info about a provider:

```BASH
sendsms provider info <provider>
```
