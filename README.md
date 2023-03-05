# kuberta
let's practice kubectl short-coding!

# About

Don't waste your time in CKAD by lengthy commands! kuberta helps you to get used to short aliases.

# Install

```bash
# download binary
$ wget https://github.com/Syuparn/kuberta/releases/download/${version}/kuberta_linux_x86_64.tar.gz
$ tar -xvzf kuberta_linux_x86_64.tar.gz
# or using Go
$ go install github.com/syuparn/kuberta@latest
```

NOTE: `kubectl` should be installed in advance

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

Tips: add `alias k=kuberta` to `~/.bashrc` so that you can run this command by `k` (ex: `k get po`)

```bash
# set alias
echo 'alias k=kuberta' >>~/.bashrc
# enable command auto-completion
echo 'complete -o default -F __start_kubectl k' >>~/.bashrc
```

(See [Official install guide](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/) for details)

# Contributions

`kuberta` uses many hacks and may have some edge-case bugs. Any contributions are always welcome!
