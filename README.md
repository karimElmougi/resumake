# resumake

CLI app to generate your resume from a YAML file using any Go template

## Install

Assuming you have $GOPATH/bin or $GOBIN added to your path:
```bash
go install github.com/karimElmougi/resumake
```


## Building from source

```bash
./build.sh
```

or simply
```bash
go build -o resumake ./main.go
```

## Usage

```bash
resumake plaintext resume.yaml > resume.txt

resumake latex resume.yaml > resume.tex && pdflatex resume.tex

resumake custom template.tmpl resume.yaml > resume.pdf
```

See [test_resume.yaml](test_resume.yaml) for an example resume.

Template files are any valid Go template. While `{{` and `}}` are the default delimiters, any delimiters can be used simply by specifying them like so:

```bash
resumake custom --delimiters [[]] template.tmpl resume.yaml
```

The delimiters argument only needs to be a string with an even length.
