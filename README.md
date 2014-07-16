## twt
---

What It Is?

> * **stack**: Golang, TwitterAPI

just some bird trials on Twitter API, right now it allows you to

* stalk
> * followers
> [TBD] * friends

* tickr
> [TBD] * timeline
> [TBD] * topic


![twt :)](/twt.jpg "twitter twt")

---

### How To Use

* Run, to create a local isolated GoEnv, install required Go libs and create a binary in project's bin directory "./bin/twt"

```bash
./go-tasks.sh bin
```

* Create an app by using your Twitter A/c @ [http://dev.twitter.com/apps](http://dev.twitter.com/apps).

* Then copy creds.json.sample to creds.json and correct the values with API Key&Secret and Access Token&Secret granted to that app.

* Run following command to test it out

```bash

./bin/twt -config creds.sample
```

---
