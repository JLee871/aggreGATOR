# aggreGATOR
## Description
Guided project from boot.dev that builds a CLI tool for automatically aggregating RSSfeeds. Uses a local database to keep track of users, feeds, and posts.
## Prerequisites
* **Go**
* **PostgreSQL**
* **goose** (Recommended)
## ⚙️ Installation
```bash
go install github.com/JLee871/aggreGATOR
```

Manually create a config file in your home directory `~/.gatorconfig.json` with the following content:
```json
{
  "db_url": "db_connection_url"
}
```
Run the database migrations in the sql/schema directory. This can be done with goose by running `goose postgres <db_connection_url> up` in the directory.

## Commands

### User-Related Commands:

### *register*
- **Usage**: `register {name}`
###### Example:
```bash
aggreGATOR register MrGator
```
###### Output:
```bash
New user created: MrGator
```
- **Description**: Adds user with the given name to the users database
***
### *login*
- **Usage**: `login {name}`
###### Example:
```bash
aggreGATOR login MrGator
```
###### Output:
```bash
User has been set to MrGator.
```
- **Description**: Sets current user to user with given name. Fails if user does not exist in the database.
***
### *users*
- **Usage**: `users`
###### Example:
```bash
aggreGATOR users
```
###### Output:
```bash
* MrFrog
* MrGator (current)
* MrCrocodile
```
- **Description**: Returns all users in the database
***
### Feed-Related Commands:

### *addfeed*
- **Usage**: `addfeed {name} {url}`
###### Example:
```bash
aggreGATOR addfeed "Lanes Blog" "https://www.wagslane.dev/index.xml"
```
###### Output:
```bash
Added feed: Lanes Blog
```
- **Description**: Adds feed with the given name and url to the feeds database. The feed is automatically followed by the user that creates it.
***
### *feeds*
- **Usage**: `feeds`
###### Example:
```bash
aggreGATOR feeds
```
###### Output:
```bash
Name: Lanes Blog , URL: https://www.wagslane.dev/index.xml , User: MrGator
Name: Hacker News RSS , URL: https://hnrss.org/newest , User: MrFrog
Name: Tech Crunch , URL: https://techcrunch.com/feed/ , User: MrFrog
```
- **Description**: Returns all feeds in database with feed name, url, and user that initially added the feed.
***
### *follow*
- **Usage**: `follow {url}`
###### Example:
```bash
aggreGATOR follow "https://hnrss.org/newest"
```
###### Output:
```bash
MrGator is now following Hacker News RSS.
```
- **Description**: Follows feed with the given url for the current user.
***
### *unfollow*
- **Usage**: `unfollow {url}`
###### Example:
```bash
aggreGATOR unfollow "https://hnrss.org/newest"
```
###### Output:
```bash
Unfollowed Hacker News RSS.
```
- **Description**: Unfollows feed with the given url for the current user.
***
### *following*
- **Usage**: `following`
###### Example:
```bash
aggreGATOR following
```
###### Output:
```bash
Lanes Blog
Tech Crunch
```
- **Description**: Returns all feeds followed by current user.
***
### *browse*
- **Usage**: `browse {number (optional, default=2)}`
###### Example:
```bash
aggreGATOR browse 3
```
###### Output:
```bash
Ask HN: How do you get the most out of Deep Research?
https://news.ycombinator.com/item?id=43304917
2025-03-09 00:24:42 +0000 +0000
=================================
2025 RootedCon: Attacking Bluetooth the easy way
https://www.documentcloud.org/documents/25554812-2025-rootedcon-bluetoothtools/
2025-03-09 00:24:17 +0000 +0000
=================================
A Mac smart quote curiosity
https://leancrew.com/all-this/2025/03/a-mac-smart-quote-curiosity/
2025-03-09 00:23:48 +0000 +0000
=================================
```
- **Description**: Returns the given number of posts from followed feeds, most recent first.
***
### *agg*
- **Usage**: `agg {timeout duration}`
###### Example:
```bash
aggreGATOR agg 60s
```
- **Description**: Constantly fetches posts from all feeds and adds them to the database, checks once every given duration time. Valid time units for duration are s (seconds), m (minutes), h (hours).
***
### Other Commands:

### *reset*
- **Usage**: `reset`
###### Example:
```bash
aggreGATOR reset
```
- **Description**: Resets the database by deleting all data. Use with caution.
