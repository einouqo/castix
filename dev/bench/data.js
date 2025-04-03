window.BENCHMARK_DATA = {
  "lastUpdate": 1743721082850,
  "repoUrl": "https://github.com/einouqo/castix",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "name": "einouqo",
            "username": "einouqo"
          },
          "committer": {
            "name": "einouqo",
            "username": "einouqo"
          },
          "id": "e3a305617650388a70e60a27a262b4506759f1fc",
          "message": "feat: synchronous unsubscribe",
          "timestamp": "2025-04-01T23:24:25Z",
          "url": "https://github.com/einouqo/castix/pull/6/commits/e3a305617650388a70e60a27a262b4506759f1fc"
        },
        "date": 1743720483422,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCastix/notify/no_one",
            "value": 10.97,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/one",
            "value": 176.3,
            "unit": "ns/op",
            "extra": "6940269 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/few",
            "value": 774.6,
            "unit": "ns/op",
            "extra": "1550506 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/more",
            "value": 2531,
            "unit": "ns/op",
            "extra": "471303 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/even_more",
            "value": 35403,
            "unit": "ns/op",
            "extra": "33806 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/a_lot",
            "value": 24878477,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "einouqo",
            "username": "einouqo"
          },
          "committer": {
            "name": "einouqo",
            "username": "einouqo"
          },
          "id": "8cde636d9af16065785c6203d28499e217839481",
          "message": "feat: synchronous unsubscribe",
          "timestamp": "2025-04-01T23:24:25Z",
          "url": "https://github.com/einouqo/castix/pull/6/commits/8cde636d9af16065785c6203d28499e217839481"
        },
        "date": 1743721082464,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCastix/notify/no_one",
            "value": 11.08,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/one",
            "value": 171.2,
            "unit": "ns/op",
            "extra": "6760158 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/few",
            "value": 737.4,
            "unit": "ns/op",
            "extra": "1618168 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/more",
            "value": 2491,
            "unit": "ns/op",
            "extra": "413520 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/even_more",
            "value": 34962,
            "unit": "ns/op",
            "extra": "34819 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/notify/a_lot",
            "value": 22736070,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          }
        ]
      }
    ]
  }
}