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
curl -s localhost:9001/ -F path=echo -F args=aa
```