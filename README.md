## usage


### and note

__-e ：是否可执行__

```bash
o hello -e echo -n Hello World! 
o gov -e go version
o passwd 12345678
```

生成的配置文件

```toml
[gov]
  Val = "'go' 'version'"
  Exced = true

[hello]
  Val = "'echo' -n 'Hello' 'World!'"
  Exced = true

[passwd]
  Val = "'12345678'"
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
passwd
hello
gov
```