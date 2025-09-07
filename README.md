![Workflow status](https://github.com/cheat/cheat/actions/workflows/build.yml/badge.svg)

cheat
=====

`cheat` 是一个命令行交互式速查表（cheatsheet）工具，专为帮助系统管理员和开发者快速查找命令用法而设计。通过AI增强功能，它不仅提供基础的命令参考，还能智能解释和优化输出内容。

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

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

Use `cheat` with [cheatsheets][].


示例
-------
比如你想查看 tar 命令的用法，只需运行：

```sh
cheat tar
```

启用 AI 功能后，你会看到优化后的输出，包含更好的结构和解释：

You will be presented with a cheatsheet resembling the following:

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

Usage
-----
To view a cheatsheet:

```sh
cheat tar      # a "top-level" cheatsheet
cheat foo/bar  # a "nested" cheatsheet
```

To edit a cheatsheet:

```sh
cheat -e tar     # opens the "tar" cheatsheet for editing, or creates it if it does not exist
cheat -e foo/bar # nested cheatsheets are accessed like this
```

To view the configured cheatpaths:

```sh
cheat -d
```

To list all available cheatsheets:

```sh
cheat -l
```

To list all cheatsheets that are tagged with "networking":

```sh
cheat -l -t networking
```

To list all cheatsheets on the "personal" path:

```sh
cheat -l -p personal
```

To search for the phrase "ssh" among cheatsheets:

```sh
cheat -s ssh
```

To search (by regex) for cheatsheets that contain an IP address:

```sh
cheat -r -s '(?:[0-9]{1,3}\.){3}[0-9]{1,3}'
```

Flags may be combined in intuitive ways. Example: to search sheets on the
"personal" cheatpath that are tagged with "networking" and match a regex:

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

Cheatsheets
-----------
Cheatsheets are plain-text files with no file extension, and are named
according to the command used to view them:

```sh
cheat tar     # file is named "tar"
cheat foo/bar # file is named "bar", in a "foo" subdirectory
```

Cheatsheet text may optionally be preceeded by a YAML frontmatter header that
assigns tags and specifies syntax:

```
---
syntax: javascript
tags: [ array, map ]
---
// To map over an array:
const squares = [1, 2, 3, 4].map(x => x * x);
```

The `cheat` executable includes no cheatsheets, but [community-sourced
cheatsheets are available][cheatsheets]. You will be asked if you would like to
install the community-sourced cheatsheets the first time you run `cheat`.

Cheatpaths
----------
Cheatsheets are stored on "cheatpaths", which are directories that contain
cheatsheets. Cheatpaths are specified in the `conf.yml` file.

It can be useful to configure `cheat` against multiple cheatpaths. A common
pattern is to store cheatsheets from multiple repositories on individual
cheatpaths:

```yaml
# conf.yml:
# ...
cheatpaths:
  - name: community                   # a name for the cheatpath
    path: ~/documents/cheat/community # the path's location on the filesystem
    tags: [ community ]               # these tags will be applied to all sheets on the path
    readonly: true                    # if true, `cheat` will not create new cheatsheets here

  - name: personal
    path: ~/documents/cheat/personal  # this is a separate directory and repository than above
    tags: [ personal ]
    readonly: false                   # new sheets may be written here
# ...
```

The `readonly` option instructs `cheat` not to edit (or create) any cheatsheets
on the path. This is useful to prevent merge-conflicts from arising on upstream
cheatsheet repositories.

If a user attempts to edit a cheatsheet on a read-only cheatpath, `cheat` will
transparently copy that sheet to a writeable directory before opening it for
editing.

### Directory-scoped Cheatpaths ###
At times, it can be useful to closely associate cheatsheets with a directory on
your filesystem. `cheat` facilitates this by searching for a `.cheat` folder in
the current working directory. If found, the `.cheat` directory will
(temporarily) be added to the cheatpaths.

Autocompletion
--------------
Shell autocompletion is currently available for `bash`, `fish`, and `zsh`. Copy
the relevant [completion script][completions] into the appropriate directory on
your filesystem to enable autocompletion. (This directory will vary depending
on operating system and shell specifics.)

Additionally, `cheat` supports enhanced autocompletion via integration with
[fzf][]. To enable `fzf` integration:

1. Ensure that `fzf` is available on your `$PATH`
2. Set an envvar: `export CHEAT_USE_FZF=true`

[INSTALLING.md]: INSTALLING.md
[Releases]:      https://github.com/cheat/cheat/releases
[cheatsheets]:   https://github.com/cheat/cheatsheets
[completions]:   https://github.com/cheat/cheat/tree/master/scripts
[fzf]:           https://github.com/junegunn/fzf
[go]:            https://golang.org
