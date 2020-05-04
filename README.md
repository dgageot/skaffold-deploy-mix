## Goal

Reproduce https://github.com/GoogleContainerTools/skaffold/issues/1753

```
skaffold build -q > build_result.json
skaffold deploy --build-artifacts=build_result.json -p david
skaffold deploy --tail --build-artifacts=build_result.json -p world
```

or

```
make test
```

The logs should show:

```
[getting-started-world getting-started] Hello World!
[getting-started-world getting-started] Hello World!
[getting-started-world getting-started] Hello World!
```

not

```
[getting-started-david getting-started] Hello David!
[getting-started-world getting-started] Hello World!
[getting-started-david getting-started] Hello David!
[getting-started-world getting-started] Hello World!
[getting-started-world getting-started] Hello World!
```

