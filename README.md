<!-- togo-brand -->
<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo" /></p>
<h1 align="center">compute-flink</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a></sub></p>

A togo **compute** backend that submits jobs to **Apache Flink** via `flink run`.

```bash
togo install togo-framework/compute-flink
togo provider:use compute flink
```

| Key | Default | Purpose |
|---|---|---|
| `FLINK_BIN` | `flink` | flink binary |
| `FLINK_OPTS` | — | extra `flink run` opts |

MIT © fadymondy
