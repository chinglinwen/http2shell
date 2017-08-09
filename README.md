# http2shell
execute shell through http

## Usage

```
Usage of ./http2shell:
  -base string
        base dir
  -p string
        listening port (default "9001")

```

## Note

base flag used to restrict which command can be execute.

## Example

```
curl localhost:9001 -d run="echo hello world"
curl localhost:9001 -d run='./a.sh a b "c c" d'   # c c as the third argument
curl localhost:9001 -d run="./a.sh a b \"c c\" d"
```

> the single quote or double quote will be trimed for the arguments.
