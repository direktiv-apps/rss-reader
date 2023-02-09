
# rss-reader 1.0

Fetch RSS feeds

---
- #### Categories: social, network
- #### Image: gcr.io/direktiv/functions/rss-reader 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/rss-reader/issues
- #### URL: https://github.com/direktiv-apps/rss-reader
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About rss-reader

This function fetches RSS feeds and converts them into JSON. It can limit the results and has a simple text search.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: rss-reader
  image: gcr.io/direktiv/functions/rss-reader:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: rss-reader
  type: action
  action:
    function: rss-reader
    input: 
      url: ttps://news.un.org/feed/subscribe/en/news/region/americas/feed/rss.xml
```
   #### Advanced
```yaml
- id: rss-reader
  type: action
  action:
    function: rss-reader
    input: 
      url: https://news.un.org/feed/subscribe/en/news/region/americas/feed/rss.xml
      limit: 1
      search: people
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Example Reponses
    
```json
{
  "Channel": {
    "Items": [
      {
        "Content": null,
        "Description": "Caribbean and UN leaders were on hand for the first regional launch of a plan to ensure that every person on the planet is protected by early warning systems by 2027, held in Bridgetown, Barbados on Tuesday.?",
        "Guid": "https://news.un.org/feed/view/en/story/2023/02/1133247",
        "Link": "https://news.un.org/feed/view/en/story/2023/02/1133247",
        "PubDate": "Tue, 07 Feb 2023 12:00:00 +0000",
        "Source": "UN News - Global perspective Human stories",
        "Title": "Caribbean sees first regional launch of global plan on early warning systems"
      }
    ]
  }
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| limit | integer| `int64` |  | | Limit number of news items | `10` |
| rss | string| `string` |  | | URL of the news feed | `https://www.theonion.com/rss` |
| search | string| `string` |  | | Simple search term for RSS items | `hello` |

 
