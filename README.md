# cronicle

A journal for your terminal.

As developers, it's easy to get lost in the weeds. `cronicle` helps you keep track of your todo list, your daily work log, and your [brag doc](https://jvns.ca/blog/brag-documents/).

Keep track of your developer journey in the command line.

## Features

- Manage your todos
- Log your daily work
- Keep track of [work you're proud of](https://jvns.ca/blog/brag-documents/)
- All your data is stored in plain text; toss it in git or take it elsewhere

## Data Formats

All plain text data in `cronicle` is stored as a combination of Markdown and YAML. We use YAML front matter to store metadata related to daily logs and brag doc items.

`tags` are optional for all data formats.

### Todo

```yaml
---
date: 2022-03-01 05:03:30
due: 2022-04-01 05:03:30
type: todo
tags:
- bug
---

Debug infinite loop on /status page
```

### Daily Log

Daily logs are stored in a file that has the date (eg. `2022-03-1`) but each individual log item is stored as a separate file. This allows you to write as much detail as you'd like for an item.

```yaml
---
date: 2022-03-01 05:03:30
type: log
tags:
- pr
---

Reviewed [PR](https://github.com/12products/warroom-frontend/pull/46) for updating UI during create events
```

### Brag Item

```yaml
---
date: 2022-03-01 05:03:30
type: brag
tags:
- talk
---

Gave talk to entire company on performance improvements.
```
