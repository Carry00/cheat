![Workflow status](https://github.com/cheat/cheat/actions/workflows/build.yml/badge.svg)


[English](./README.en.md) | [简体中文](./README.md) | [项目详情](./wiki.zh.md) 


cheat
=====

`cheat` 是一个命令行交互式速查表（cheatsheet）工具，专为帮助系统管理员和开发者快速查找命令用法而设计。通过AI增强功能，它不仅提供基础的命令参考，还能智能解释和优化输出内容。

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

 **关于漫画**

*   **情节 (Plot):** 一个人（Rob）被紧急叫去拆一个炸弹，拆弹的方式是在10秒内输入一个正确的 `tar` 命令。然而，即使是 "use UNIX" 的 Rob，也因为 `tar` 命令复杂的参数组合而当场愣住，最终导致了失败。
*   **笑点 (The Joke):** `tar` 命令的参数（如 `-xvf`, `-cvf`, `-czvf`, `-xzvf` 等）非常多且难以记忆，以至于在巨大压力下，一个经验丰富的用户也可能无法立刻回忆起来。这幅漫画用一种幽默而夸张的方式吐槽了这一点。
如果漫画里的 Rob 当时拥有了 `cheat` 工具，他可能只需要从容地输入 `cheat tar`，就能立即获得正确指令，轻松化解危机。

主要特性
-------
- 创建和查看命令行速查表
- 支持分层组织的速查表系统
- 标签化管理和搜索
- **AI增强输出**：通过AI优化内容的可读性和理解性
- 完全可定制的配置系统

AI功能
------
新增的AI功能让速查表更智能、更易理解：

- **智能解释**：自动为专业术语添加解释
- **优化格式**：改善输出的结构和排版
- **上下文关联**：提供相关命令和用例建议
- **交互增强**：保持命令的准确性的同时提供更多上下文信息

### AI配置

在 `~/.config/cheat/conf.yml` 中配置AI功能：

```yaml
# AI功能配置
ai_enabled: true                     # 启用AI处理
ai_url: "YOUR_AI_SERVICE_URL"        # AI服务端点
ai_key: "YOUR_API_KEY"              # API密钥
ai_model: "gpt-3.5-turbo"          # 使用的AI模型

# 自定义AI处理指令
ai_system_prompt: |
  你是一个AI助手，负责处理cheat命令的输出。请：
  1. 增强内容的可读性
  2. 为专业术语添加简短解释
  3. 保持原始命令和示例的准确性
  4. 如果是代码示例，保持格式并添加注释
  5. 按照以下格式组织输出：
     - 命令说明
     - 参数解释（如果有）
     - 使用示例
     - 注意事项（如果有）

# AI响应的最大token数
ai_max_tokens: 2000
```

配合 [cheatsheets][] 使用 `cheat`。

示例
-------
比如你想查看 tar 命令的用法，只需运行：

```sh
cheat tar
```

启用 AI 功能后，你会看到优化后的输出，包含更好的结构和解释：

你将看到类似下面的速查表：

```sh
要提取一个未压缩的归档文件:
tar -xvf /path/to/foo.tar

要在指定的目录中提取一个 .tar 文件:
tar -xvf /path/to/foo.tar -C /path/to/destination/

要创建一个未压缩的归档文件:
tar -cvf /path/to/foo.tar /path/to/foo/

要提取一个 .tgz 或 .tar.gz 归档文件:
tar -xzvf /path/to/foo.tgz
tar -xzvf /path/to/foo.tar.gz

要创建一个 .tgz 或 .tar.gz 归档文件:
tar -czvf /path/to/foo.tgz /path/to/foo/
tar -czvf /path/to/foo.tar.gz /path/to/foo/

要列出 .tgz 或 .tar.gz 归档文件的内容:
tar -tzvf /path/to/foo.tgz
tar -tzvf /path/to/foo.tar.gz

要提取一个 .tar.bz2 归档文件:
tar -xjvf /path/to/foo.tar.bz2

要创建一个 .tar.bz2 归档文件:
tar -cjvf /path/to/foo.tar.bz2 /path/to/foo/

要列出 .tar.bz2 归档文件的内容:
tar -tjvf /path/to/foo.tar.bz2

要创建一个 .tgz 归档文件并排除所有 jpg、gif 等文件:
tar -czvf /path/to/foo.tgz --exclude=\*.{jpg,gif,png,wmv,flv,tar.gz,zip} /path/to/foo/

要使用压缩算法的并行（多线程）实现:
将 tar -z ... 替换为 tar -Ipigz ...
将 tar -j ... 替换为 tar -Ipbzip2 ...
将 tar -J ... 替换为 tar -Ipixz ...

要将新文件追加到旧的 tar 归档文件中:
tar -rf <archive.tar> <new-file-to-append>
```

使用方法
-----
查看速查表：

```sh
cheat tar      # 查看一个“顶级”速查表
cheat foo/bar  # 查看一个“嵌套”速查表
```

编辑速查表：

```sh
cheat -e tar     # 打开 "tar" 速查表进行编辑，如果不存在则创建
cheat -e foo/bar # 嵌套的速查表也这样访问
```

查看已配置的速查表路径：

```sh
cheat -d
```

列出所有可用的速查表：

```sh
cheat -l
```

列出所有带有 "networking" 标签的速查表：

```sh
cheat -l -t networking
```

列出 "personal" 路径下的所有速查表：

```sh
cheat -l -p personal
```

在速查表中搜索 "ssh" 短语：

```sh
cheat -s ssh
```

通过正则表达式搜索包含 IP 地址的速查表：

```sh
cheat -r -s '(?:[0-9]{1,3}\.){3}[0-9]{1,3}'
```

参数可以直观地组合使用。例如：在 "personal" 路径下搜索带有 "networking" 标签且匹配正则表达式的速查表：

```sh
cheat -p personal -t networking --regex -s '(?:[0-9]{1,3}\.){3}[0-9]{1,3}'
```

安装
----------

### 从源代码构建（推荐）

要获取支持AI功能的版本，请从我的fork仓库克隆和构建：

```bash
# 克隆仓库
git clone https://github.com/carry00/cheat

# 进入项目目录
cd cheat

# 编译项目
go build -o cheat ./cmd/cheat

# 将编译好的二进制文件移动到PATH中的某个位置
sudo mv cheat /usr/local/bin/
```

### 首次运行配置

1. 首次运行时创建配置文件：
```bash
mkdir -p ~/.config/cheat && cheat --init > ~/.config/cheat/conf.yml
```

2. 编辑配置文件添加AI功能：
```bash
vim ~/.config/cheat/conf.yml
```

3. 按照上面的"AI配置"部分设置相关参数

更多安装选项，请参考 [INSTALLING.md][]。

速查表
-----------
速查表是纯文本文件，没有文件扩展名，并根据用于查看它们的命令来命名：

```sh
cheat tar     # 文件名为 "tar"
cheat foo/bar # 文件名为 "bar"，位于 "foo" 子目录中
```

速查表文本可以选择性地在开头添加 YAML frontmatter 头部，用于分配标签和指定语法：

```
---
syntax: javascript
tags: [ array, map ]
---
// To map over an array:
const squares = [1, 2, 3, 4].map(x => x * x);
```

`cheat` 可执行文件本身不包含任何速查表，但你可以使用[社区贡献的速查表][cheatsheets]。首次运行 `cheat` 时，系统会询问你是否要安装社区贡献的速查表。

速查表路径 (Cheatpaths)
----------
速查表存储在 "cheatpaths" 中，这些是包含速查表文件的目录。Cheatpaths 在 `conf.yml` 文件中指定。

为 `cheat` 配置多个 cheatpath 会非常有用。一种常见的模式是将来自多个仓库的速查表存储在各自独立的 cheatpath 中：

```yaml
# conf.yml:
# ...
cheatpaths:
  - name: community                   # cheatpath 的名称
    path: ~/documents/cheat/community # 该路径在文件系统中的位置
    tags: [ community ]               # 这些标签将应用于该路径下的所有速查表
    readonly: true                    # 如果为 true，`cheat` 将不会在此处创建新的速查表

  - name: personal
    path: ~/documents/cheat/personal  # 这是一个与上面不同的独立目录和仓库
    tags: [ personal ]
    readonly: false                   # 可以在此处写入新的速查表
# ...
```

`readonly` 选项指示 `cheat` 不要编辑（或创建）该路径下的任何速查表。这对于防止在上游速查表仓库中产生合并冲突非常有用。

如果用户尝试编辑只读 cheatpath 上的速查表，`cheat` 会在打开编辑前，自动将该速查表复制到一个可写目录中。

### 目录作用域的速查表路径 ###
有时，将速查表与文件系统中的某个特定目录紧密关联会很有用。`cheat` 支持此功能，它会在当前工作目录中查找 `.cheat` 文件夹。如果找到，`.cheat` 目录将被（临时）添加到 cheatpaths 中。

自动补全
--------------
目前支持 `bash`、`fish` 和 `zsh` 的 Shell 自动补全。将相关的[补全脚本][completions]复制到你文件系统中相应的目录即可启用。（具体目录取决于你的操作系统和 Shell 类型。）

此外，`cheat` 通过与 [fzf][] 集成支持增强的自动补全功能。要启用 `fzf` 集成：

1. 确保 `fzf` 在你的 `$PATH` 路径中可用
2. 设置环境变量：`export CHEAT_USE_FZF=true`

[INSTALLING.md]: INSTALLING.md
[Releases]:      https://github.com/cheat/cheat/releases
[cheatsheets]:   https://github.com/cheat/cheatsheets
[completions]:   https://github.com/cheat/cheat/tree/master/scripts
[fzf]:           https://github.com/junegunn/fzf
[go]:            https://golang.org