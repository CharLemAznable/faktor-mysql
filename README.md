### faktor-mysql

[![Build Status](https://app.travis-ci.com/CharLemAznable/faktor-mysql.svg?branch=main)](https://app.travis-ci.com/CharLemAznable/faktor-mysql)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/CharLemAznable/faktor-mysql)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
![GitHub code size](https://img.shields.io/github/languages/code-size/CharLemAznable/faktor-mysql)

伪装mysql服务, 用于在开发/测试阶段本机使用, 轻量化本机环境.

项目名: fa**k**~~e~~ ~~vi~~**k**tor

Viktor(维克托): <Arcane: League of Legends>(<英雄联盟：双城之战>)中的角色, 未来的The Machine Herald(机械先驱).

#### 获取可执行文件

1. 下载最新的可执行文件压缩包并解压

    下载地址: [faktor-mysql releases](https://github.com/CharLemAznable/faktor-mysql/releases)

    *Mac系统如无法执行bin文件, 请在 ```系统偏好设置 -> 安全性与隐私``` 中允许程序运行*

```bash
$ tar -xvJf faktor-[version].[arch].[os].tar.xz
```

2. git clone到本地, 编译打包

    获取打包后的二进制文件```faktor.bin```

```bash
$ ./build.sh
```

#### 启动

启动伪数据库服务后, 进入命令行模式

```bash
$ ./faktor.bin
########    ###    ##    ## ########  #######  ########  
##         ## ##   ##   ##     ##    ##     ## ##     ##
##        ##   ##  ##  ##      ##    ##     ## ##     ##
######   ##     ## #####       ##    ##     ## ########  
##       ######### ##  ##      ##    ##     ## ##   ##   
##       ##     ## ##   ##     ##    ##     ## ##    ##  
##       ##     ## ##    ##    ##     #######  ##     ##

faktor »  
```

可选参数:
* --log: 指定伪数据库日志输出路径, 默认为faktor.db.log
* --port/-p: 指定伪数据库端口号, 默认9900
* --username/-u: 指定伪数据库用户名, 默认auser
* --password/-p: 指定伪数据库用户密码, 默认sa

#### 加载数据库

```bash
faktor » load [schema name] [init sql script file path, optional]
```

#### 卸载数据库

```bash
faktor » unload [schema name]
```

#### 查看帮助

```bash
faktor » help
########    ###    ##    ## ########  #######  ########  
##         ## ##   ##   ##     ##    ##     ## ##     ##
##        ##   ##  ##  ##      ##    ##     ## ##     ##
######   ##     ## #####       ##    ##     ## ########  
##       ######### ##  ##      ##    ##     ## ##   ##   
##       ##     ## ##   ##     ##    ##     ## ##    ##  
##       ##     ## ##    ##    ##     #######  ##     ##


fake mysql server

Commands:
  clear         clear the screen
  exit          exit the shell
  help          use 'help [command]' for command help
  load, create  load a schema [ with init sql script ]
  unload, drop  unload a schema
```

```bash
faktor » help load

load a schema [ with init sql script ]

Usage:
  load [flags] schema [script]

Args:
  schema  string    schema name
  script  string    schema init script

Flags:
  -h, --help     display help
```

```bash
faktor » help unload

unload a schema

Usage:
  unload [flags] schema

Args:
  schema  string    schema name

Flags:
  -h, --help     display help
```
