# github-user-activity-cli

Simple CLI to fetch recent activity of users from Github API and display it in terminal.

---

## Installation

### Clone repository

```bash
git clone https://github.com/grizlier/github-user-activity-cli.git
cd github-user-activity-cli
```

### Build

```bash
go build -o github-activity
```

---

## Usage

Run the CLI with a GitHub username:

```bash
./github-activity <username>
```
Example:

```bash
./github-activity kamranahmedse
```
---

## Future improvements

- Better formatting (colors, grouping)
- Support for private activity via token
- Caching responses
- Error handling improvements

---

## About project

This project was created to practice working with the GitHub API and JSON parsing
