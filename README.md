# Gotemplate Test

This is a sample project for testing gotemplate. Edit the main fn to your requirements.
Paste your gotemplate in `content` string and set the values used inside the gotemplate in `values`.
If you want to use `.instance.metadata.name` inside your gotemplate, use `values.set(".instance.metadata.name", "hello")` to set the value as `hello`.


## Build and run:
```
$ go build -o gotemplate
$ ./gotemplate
```