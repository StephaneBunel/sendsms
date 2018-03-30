# sendsms

A simple command line (and a library) to send SMS messages throught various providers (gateways).

Currently and because it was my primary goal, only OVH provider is implemented.
It is easy to add more provider. The code is designed with that in mind.

## Configuration

sendsms use a configuration file (YAML formated) to define profiles.
Profiles is the way to configure multiple providers/credentials to use.

Example: (config.yml)

```YAML
---
sendsms:
  logLevel: "INFO"

  profiles:
    default:                       # <-- profile name
      provider:              "ovh" # <-- provider to use for this profile
      providerConfig:              # <-- provider configuration for this profile
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

Send a message to a phone number using default profile:

```BASH
sendsms send --phone +336123456789 --message "Hello my friend !"
```

You can specify multiple recipient at a time:

```BASH
sendsms send --phone +336123456789 --phone +336987654321 --message "Hello my friend !"
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
sendsms provider info <provider name>
```
