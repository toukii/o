## usage

```
no go -e go help
no passwd 12345678
```

```toml
[go]
  Val = " 'go' 'help'"
  Exced = true

[gv]
  Val = " 'go' 'version'"
  Exced = true

[hello]
  Val = " 'echo' 'Hello World! '"
  Exced = true

[passwd]
  Val = " '12345678'"
  Exced = false
```



```
➜  no git:(master) ✗ no hello
 'echo' 'Hello World! '
Hello World!

➜  no git:(master) ✗ no -l
*********** keys **********
gv
hello
go
```