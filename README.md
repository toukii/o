## usage


### and note

```bash
o go -e go help
o passwd 12345678
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



```bash
➜  o git:(master) ✗ o hello
 'echo' 'Hello World! '
Hello World!

➜  o git:(master) ✗ o hel
hel ≈≈> hello
'echo' -n 'Hello World! '
Hello World!

➜  o git:(master) ✗ o
*********** keys **********
gv
hello
go
```