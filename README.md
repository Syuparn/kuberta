# kuberta
let's practice kubectl short-coding!

# About

Don't waste your time in CKAD by lengthy commands! kuberta helps you to get used to short aliases.

# Usage

`kuberta` works as a thin wrapper of `kubectl`. It raises errors if you forget to use short aliases.

```bash
$ kuberta get replicasets
ERROR: too long! should be `kubectl get rs`

$ kuberta get rs --namespace foo
ERROR: too long! should be `kubectl get rs -n foo`

# if you use short aliases properly, this works just same as kubectl
$ kuberta get rs -n foo
NAME              DESIRED   CURRENT   READY   AGE
nginx-76d6c9b8c   1         1         0       9s
```

# Contributions

`kuberta` uses many hacks and may have some edge-case bugs. Any contributions are always welcome!
