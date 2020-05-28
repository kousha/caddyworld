# caddyworld
My playground for messing around with Caddy v2

Barebones caddy app.

---

## Build and Run
* `go build`
* `./caddyworld run`
* in a new terminal
  * `curl localhost:2019/load -X POST -H "Content-Type: application/json" -d @caddy.json  #load config`

