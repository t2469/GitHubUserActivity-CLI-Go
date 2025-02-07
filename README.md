# GitHub User Activity CLI

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org)

A simple CLI app that fetches a GitHub user's recent events and displays them in JST.

**Usage:**

```bash
bin/github-activity <username>
```

Example Output:

```plaintext
Event: CreateEvent, Repo: t2469/chat-app, Created At: 2025-02-07T00:13:42+09:00
Event: IssuesEvent, Repo: t2469/chat-app, Created At: 2025-02-06T22:51:52+09:00
Event: DeleteEvent, Repo: t2469/chat-app, Created At: 2025-02-05T00:46:12+09:00
...
```
