![Workflow status](https://github.com/cheat/cheat/actions/workflows/build.yml/badge.svg)

[English](./README.en.md) | [简体中文](./README.md) 

cheat
=====

`cheat` is a command-line interactive cheatsheet tool designed to help system administrators and developers quickly find command usages. With AI-enhanced features, it not only provides basic command references but also intelligently explains and optimizes the output content.

![The obligatory xkcd](http://imgs.xkcd.com/comics/tar.png 'The obligatory xkcd')

Main Features
-------
- Create and view command-line cheatsheets
- Supports a hierarchically organized cheatsheet system
- Tag-based management and search
- **AI-Enhanced Output**: Improves content readability and comprehension through AI
- Fully customizable configuration system

AI Features
------
The new AI features make cheatsheets smarter and easier to understand:

- **Smart Explanation**: Automatically adds explanations for technical terms
- **Optimized Formatting**: Improves the structure and layout of the output
- **Contextual Relevance**: Provides suggestions for related commands and use cases
- **Enhanced Interaction**: Provides more contextual information while maintaining the accuracy of the commands

### AI Configuration

Configure AI features in `~/.config/cheat/conf.yml`:

```yaml
# AI feature configuration
ai_enabled: true                     # Enable AI processing
ai_url: "YOUR_AI_SERVICE_URL"        # AI service endpoint
ai_key: "YOUR_API_KEY"               # API key
ai_model: "gpt-3.5-turbo"            # AI model to use

# Custom instructions for AI processing
ai_system_prompt: |
  You are an AI assistant responsible for processing the output of the `cheat` command. Please:
  1. Enhance the readability of the content
  2. Add brief explanations for technical terms
  3. Maintain the accuracy of the original commands and examples
  4. If it's a code example, preserve the formatting and add comments
  5. Organize the output in the following format:
     - Command Description
     - Parameter Explanation (if any)
     - Usage Examples
     - Notes (if any)

# Maximum tokens for AI response
ai_max_tokens: 2000
```

Use `cheat` with [cheatsheets][].


Example
-------
For example, if you want to see the usage for the `tar` command, just run:

```sh
cheat tar
```

With the AI feature enabled, you will see an optimized output with better structure and explanations:

You will be presented with a cheatsheet resembling the following:

```sh
To extract an uncompressed archive:
tar -xvf /path/to/foo.tar

To extract a .tar file in a specified directory:
tar -xvf /path/to/foo.tar -C /path/to/destination/

To create an uncompressed archive:
tar -cvf /path/to/foo.tar /path/to/foo/

To extract a .tgz or .tar.gz archive:
tar -xzvf /path/to/foo.tgz
tar -xzvf /path/to/foo.tar.gz

To create a .tgz or .tar.gz archive:
tar -czvf /path/to/foo.tgz /path/to/foo/
tar -czvf /path/to/foo.tar.gz /path/to/foo/

To list the contents of a .tgz or .tar.gz archive:
tar -tzvf /path/to/foo.tgz
tar -tzvf /path/to/foo.tar.gz

To extract a .tar.bz2 archive:
tar -xjvf /path/to/foo.tar.bz2

To create a .tar.bz2 archive:
tar -cjvf /path/to/foo.tar.bz2 /path/to/foo/

To list the contents of a .tar.bz2 archive:
tar -tjvf /path/to/foo.tar.bz2

To create a .tgz archive and exclude all jpg, gif, etc. files:
tar -czvf /path/to/foo.tgz --exclude=\*.{jpg,gif,png,wmv,flv,tar.gz,zip} /path/to/foo/

To use a parallel (multi-threaded) implementation of the compression algorithm:
Replace tar -z ... with tar -Ipigz ...
Replace tar -j ... with tar -Ipbzip2 ...
Replace tar -J ... with tar -Ipixz ...

To append new files to an old tar archive:
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



Installation
----------

### Build from Source (Recommended)

To get the version with AI support, clone and build from my forked repository:

```bash
# Clone the repository
git clone https://github.com/carry00/cheat

# Change into the project directory
cd cheat

# Compile the project
go build -o cheat ./cmd/cheat

# Move the compiled binary to a location in your PATH
sudo mv cheat /usr/local/bin/
```

### First-run Configuration

1.  Create the configuration file on the first run:
    ```bash
    mkdir -p ~/.config/cheat && cheat --init > ~/.config/cheat/conf.yml
    ```

2.  Edit the configuration file to add AI features:
    ```bash
    vim ~/.config/cheat/conf.yml
    ```

3.  Set the relevant parameters according to the "AI Configuration" section above.

For more installation options, please refer to [INSTALLING.md][].

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

1.  Ensure that `fzf` is available on your `$PATH`
2.  Set an envvar: `export CHEAT_USE_FZF=true`

[INSTALLING.md]: INSTALLING.md
[Releases]:      https://github.com/cheat/cheat/releases
[cheatsheets]:   https://github.com/cheat/cheatsheets
[completions]:   https://github.com/cheat/cheat/tree/master/scripts
[fzf]:           https://github.com/junegunn/fzf
[go]:            https://golang.org
