Plain
=====

Want to read a web article without all the cruft? `plain` grabs the text you care about and leave out the rest.

Installation
------------

You'll need to have Go installed on your machine. The [getting started](https://golang.org/doc/install) guide is a good walkthrough of how to do this.

To install the `plain` binary, run this command:

```shell
go install github.com/qsymmachus/plain
```

Usage
-----

The `plain` command accepts a `-url` flag. It will download the page at that address, extract the article's text, and print it out as plaintext. Here's an example:

```shell
plain -url https://henrikwarne.com/2020/03/22/secure-by-design/
```

This just prints the article text to standard output. Optionally, you may output the text to a file using the `-file` flag:

```shell
plain -url https://henrikwarne.com/2020/03/22/secure-by-design/ -file secure-by-design.txt
```

__A few caveats:__ the program is not very clever. Currently it just extracts the text from all paragraph and header tags. This means that other text that might be relevant will be skipped, and it can't handle non-HTML content at all. I'll keep working on it, but as-is, it works well for most text articles you'll encounter.
