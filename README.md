# Gotemplate Test

This is a sample project for testing gotemplate. Edit the `main` function and `template.yaml` to your requirements.
Put your gotemplate in `template.yaml` file and set the values used inside the gotemplate in `values` object in `main` function.
If you want to use `.instance.metadata.name` inside your gotemplate, use `values.set(".instance.metadata.name", "hello")` to set the value as `hello`.


## Build and run:
```
$ go build -o gotemplate
$ ./gotemplate
```