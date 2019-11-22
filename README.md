## Two-step authenticator

Use this script to give me the two-step verification code just in case that
I am too lazy to turn on the cell phone.

### Usage
```
gpg -d secrets.json.asc | go run authenticator.go
```

----

>### Install
>
>```
>$ git clone git@github.com:erning/authenticator.git
>$ cd authenticator
>$ virtualenv .virtualenv
>$ ve pip install -r requirements.txt
>```
>
>### Usage
>
>```
>$ gpg -d secrets.json.asc | ve python authenticator
>```

### secrets.json

Secrets.json is a JSON data file in plain text. for example,

```json
[
  ["${description}",        "${secret_key}"],
  ["zhang@erning.com",      "xxxx xxxx xxxx xxxx xxxx xxxx xxxx xxxx"],
  ["zendragon@hotmail.com", "xxxx xxxx xxxx xxxx"]
]
```

To encrypt secrets.json

```
$ gpg -ase -r zhang@erning.com -r zhang.yining@gmail.com secrets.json
```

Remember to delete the original secrets.json file
