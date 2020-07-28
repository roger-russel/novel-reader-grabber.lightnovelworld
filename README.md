# novel-reader-grabber.lightnovelworld
a simple grabber to lightnovelworld

## Json contract

### Novel list

[
  {
    title: string
    cover: image
    Auth: string
    lastUpdate: datetime
    tags: []string
    complete: bool
    chapters: int
    release: string # frequence
  }
]

### Novel

{
  title: string
  cover: image
  autor: string
  Description: string
  complete: bool
  tags: []string
  chapters: [{
    title: string
    url: string
    unread: bool
    release-date: datetime
    downloaded: bool
  }]
}
