
# Incrowd - Sports News 

Regarding the API implementation, I initially considered making direct calls to the provider feed data for each request, due to would be always the latest, and then if in any case they servers had issues, use mongoDB as our data provider. However, I recognized that this approach may not always be the fastest, due to the latency making httpRequest inside our endpoints,that could result in delays for our customers. As such, I decided to implement a second option where the data will be always consumed directly from our mongo database. This approach would ensure faster load times. By the way, if we would like real-time, for example providing info in real-time about a match for people who don't pay for Champions subscription, we would go for websockets.

To address this, I implemented a cronjob that runs every 10 minutes, which periodically updates the data in our mongo database. This approach balances the need for resource availability with the need to avoid overloading the provider feed API with too many requests. Instead, we want to strike a balance between freshness and efficiency, which is why we decided to use a 10-minute cron timer.

The importance of this approach is illustrated by the example of Leo Messi returning to FC Barcelona. If our cron job is set to run every 30 min or 2 hours, there is a possibility that the news about Messi's return might not appear in our database, which would impact badly in our traffic and customer experience, and we dont want that in incrowd. To avoid this, we need to periodically update our database with fresh data from the provider feed API.

Overall, I believe that the approach I took to the technical test strikes a balance between efficiency, freshness, and customer satisfaction, and will be effective in meeting the needs of business.
 
## Authors

- [@xavimg](https://github.com/xavimg)


## Run Locally

Clone the project

```bash
  git clone https://github.com/xavimg/articles.git
```
Run this Makefile command:

```bash
  make up_build
```

## Test Locally

Run this Makefile command:

```bash
  make test
```

## Screenshots

![diagram](diagram.png)
