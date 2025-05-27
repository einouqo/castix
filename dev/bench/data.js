window.BENCHMARK_DATA = {
  "lastUpdate": 1748389155381,
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
          "id": "a2b802161c1f4c949c67d55e507eccf4906de7d8",
          "message": "feat: new goroutine mux",
          "timestamp": "2025-05-19T21:56:48Z",
          "url": "https://github.com/einouqo/castix/pull/11/commits/a2b802161c1f4c949c67d55e507eccf4906de7d8"
        },
        "date": 1748388718834,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCastix/goroutine/inputs/one",
            "value": 650.4,
            "unit": "ns/op",
            "extra": "1803399 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/few",
            "value": 2767,
            "unit": "ns/op",
            "extra": "425005 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/more",
            "value": 11531,
            "unit": "ns/op",
            "extra": "100890 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/even_more",
            "value": 197668,
            "unit": "ns/op",
            "extra": "6189 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/a_lot",
            "value": 811217,
            "unit": "ns/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/one",
            "value": 652.4,
            "unit": "ns/op",
            "extra": "1855341 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/few",
            "value": 1730,
            "unit": "ns/op",
            "extra": "669154 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/more",
            "value": 5444,
            "unit": "ns/op",
            "extra": "226495 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/even_more",
            "value": 81026,
            "unit": "ns/op",
            "extra": "14874 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/a_lot",
            "value": 301993,
            "unit": "ns/op",
            "extra": "6562 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/one",
            "value": 1311,
            "unit": "ns/op",
            "extra": "900349 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/few",
            "value": 8757,
            "unit": "ns/op",
            "extra": "134480 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/more",
            "value": 88064,
            "unit": "ns/op",
            "extra": "13840 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/even_more",
            "value": 12599200,
            "unit": "ns/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/a_lot",
            "value": 226148045,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/one",
            "value": 971.7,
            "unit": "ns/op",
            "extra": "1213232 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/few",
            "value": 4920,
            "unit": "ns/op",
            "extra": "213825 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/more",
            "value": 38475,
            "unit": "ns/op",
            "extra": "31304 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/even_more",
            "value": 7233859,
            "unit": "ns/op",
            "extra": "200 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/a_lot",
            "value": 90739746,
            "unit": "ns/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/one",
            "value": 938.2,
            "unit": "ns/op",
            "extra": "1272154 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/few",
            "value": 2023,
            "unit": "ns/op",
            "extra": "561908 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/more",
            "value": 5670,
            "unit": "ns/op",
            "extra": "210651 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/even_more",
            "value": 80724,
            "unit": "ns/op",
            "extra": "14652 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/a_lot",
            "value": 301276,
            "unit": "ns/op",
            "extra": "7491 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/one",
            "value": 1988,
            "unit": "ns/op",
            "extra": "575796 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/few",
            "value": 12077,
            "unit": "ns/op",
            "extra": "98614 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/more",
            "value": 118948,
            "unit": "ns/op",
            "extra": "9526 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/even_more",
            "value": 23941371,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/a_lot",
            "value": 580348076,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
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
          "id": "c332111ddc444f014fd04b3b667d41c6f33f1827",
          "message": "feat: new goroutine mux",
          "timestamp": "2025-05-19T21:56:48Z",
          "url": "https://github.com/einouqo/castix/pull/11/commits/c332111ddc444f014fd04b3b667d41c6f33f1827"
        },
        "date": 1748388918853,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCastix/goroutine/inputs/one",
            "value": 661.2,
            "unit": "ns/op",
            "extra": "1768378 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/few",
            "value": 2758,
            "unit": "ns/op",
            "extra": "423936 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/more",
            "value": 11668,
            "unit": "ns/op",
            "extra": "102380 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/even_more",
            "value": 196263,
            "unit": "ns/op",
            "extra": "5834 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/a_lot",
            "value": 806838,
            "unit": "ns/op",
            "extra": "1372 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/one",
            "value": 669.9,
            "unit": "ns/op",
            "extra": "1829865 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/few",
            "value": 1754,
            "unit": "ns/op",
            "extra": "665629 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/more",
            "value": 5345,
            "unit": "ns/op",
            "extra": "192747 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/even_more",
            "value": 81924,
            "unit": "ns/op",
            "extra": "14521 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/a_lot",
            "value": 302589,
            "unit": "ns/op",
            "extra": "7723 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/one",
            "value": 1314,
            "unit": "ns/op",
            "extra": "929686 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/few",
            "value": 8791,
            "unit": "ns/op",
            "extra": "136777 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/more",
            "value": 88423,
            "unit": "ns/op",
            "extra": "13652 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/even_more",
            "value": 9077405,
            "unit": "ns/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/a_lot",
            "value": 336209145,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/one",
            "value": 913.6,
            "unit": "ns/op",
            "extra": "1314459 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/few",
            "value": 4856,
            "unit": "ns/op",
            "extra": "239820 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/more",
            "value": 38396,
            "unit": "ns/op",
            "extra": "31339 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/even_more",
            "value": 7151826,
            "unit": "ns/op",
            "extra": "210 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/a_lot",
            "value": 89211791,
            "unit": "ns/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/one",
            "value": 929.5,
            "unit": "ns/op",
            "extra": "1284054 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/few",
            "value": 2016,
            "unit": "ns/op",
            "extra": "591416 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/more",
            "value": 5684,
            "unit": "ns/op",
            "extra": "212229 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/even_more",
            "value": 80839,
            "unit": "ns/op",
            "extra": "14625 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/a_lot",
            "value": 303482,
            "unit": "ns/op",
            "extra": "3334 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/one",
            "value": 1977,
            "unit": "ns/op",
            "extra": "590468 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/few",
            "value": 12169,
            "unit": "ns/op",
            "extra": "98320 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/more",
            "value": 119291,
            "unit": "ns/op",
            "extra": "9391 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/even_more",
            "value": 24063049,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/a_lot",
            "value": 582704142,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
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
          "id": "f860524be7b40531128a53b8e102c83908d3ca41",
          "message": "feat: new goroutine mux",
          "timestamp": "2025-05-19T21:56:48Z",
          "url": "https://github.com/einouqo/castix/pull/11/commits/f860524be7b40531128a53b8e102c83908d3ca41"
        },
        "date": 1748389154901,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCastix/goroutine/inputs/one",
            "value": 668.9,
            "unit": "ns/op",
            "extra": "1849500 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/few",
            "value": 2809,
            "unit": "ns/op",
            "extra": "423105 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/more",
            "value": 11720,
            "unit": "ns/op",
            "extra": "97831 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/even_more",
            "value": 195914,
            "unit": "ns/op",
            "extra": "5718 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/inputs/a_lot",
            "value": 803604,
            "unit": "ns/op",
            "extra": "1446 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/one",
            "value": 649.1,
            "unit": "ns/op",
            "extra": "1847007 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/few",
            "value": 1788,
            "unit": "ns/op",
            "extra": "680887 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/more",
            "value": 5452,
            "unit": "ns/op",
            "extra": "191318 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/even_more",
            "value": 81678,
            "unit": "ns/op",
            "extra": "14754 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/outputs/a_lot",
            "value": 307117,
            "unit": "ns/op",
            "extra": "6639 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/one",
            "value": 1345,
            "unit": "ns/op",
            "extra": "879714 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/few",
            "value": 8724,
            "unit": "ns/op",
            "extra": "133639 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/more",
            "value": 88660,
            "unit": "ns/op",
            "extra": "13826 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/even_more",
            "value": 10910439,
            "unit": "ns/op",
            "extra": "144 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/goroutine/input-output/a_lot",
            "value": 214786792,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/one",
            "value": 911.2,
            "unit": "ns/op",
            "extra": "1300812 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/few",
            "value": 4915,
            "unit": "ns/op",
            "extra": "237376 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/more",
            "value": 38806,
            "unit": "ns/op",
            "extra": "31245 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/even_more",
            "value": 6968647,
            "unit": "ns/op",
            "extra": "200 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/inputs/a_lot",
            "value": 89143060,
            "unit": "ns/op",
            "extra": "12 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/one",
            "value": 924,
            "unit": "ns/op",
            "extra": "1299498 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/few",
            "value": 2057,
            "unit": "ns/op",
            "extra": "570891 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/more",
            "value": 5771,
            "unit": "ns/op",
            "extra": "188752 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/even_more",
            "value": 82777,
            "unit": "ns/op",
            "extra": "14828 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/outputs/a_lot",
            "value": 300704,
            "unit": "ns/op",
            "extra": "3507 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/one",
            "value": 1963,
            "unit": "ns/op",
            "extra": "599535 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/few",
            "value": 12154,
            "unit": "ns/op",
            "extra": "98290 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/more",
            "value": 119732,
            "unit": "ns/op",
            "extra": "9016 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/even_more",
            "value": 24187796,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkCastix/reflection/input-output/a_lot",
            "value": 575434224,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          }
        ]
      }
    ]
  }
}