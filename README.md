# json2mdtbl

JSON to Markdown Table

# Installation

Download from https://github.com/tkuchiki/json2mdtbl/releases

# Usage

```console
$ echo -e '[{"name": "alice", "foo":"bar"},{"name":"bob", "foo":"baz"}]' | ./json2mdtbl
| FOO | NAME  |
|-----|-------|
| bar | alice |
| baz | bob   |

$ echo -e '{"name": "alice", "foo":"bar"}\n{"name":"bob", "foo":"baz"}' | ./json2mdtbl
| FOO | NAME  |
|-----|-------|
| bar | alice |
| baz | bob   |
```