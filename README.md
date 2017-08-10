# http2shell
execute shell through http

## Usage

```
Usage of ./http2shell:
  -base string
        base dir
  -p string
        listening port (default "9001")
  -t    enable jaeger tracing
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

## Run jaeger

in case if you want to enable tracing

```
sudo docker run -d -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
  -p5778:5778 -p16686:16686 -p14268:14268 jaegertracing/all-in-one:latest
```