
# Incrowd - Sports News

**We want to provide stability in new News by implementing cron-job calling external data feed and transforming data into consistent and desirable format that the app developers can consume, so that, as is often the case, when the external provider has issues we can still provide data to the apps, albeit stale data.**

To address this, I implemented a cronjob that runs every 5-minute, which periodically updates the data in our mongoDB. This approach balances the need for resource availability with the need to avoid overloading the provider feed API with too many requests. Instead, we want to strike a balance between freshness and efficiency, which is why we decided to use a 5-minute cron timer. As a general rule, it is recommended to limit the number of requests per minute to an external API to a reasonable number to avoid overloading the server or causing a denial of service (DoS) attack. Typically, most APIs have rate limiting policies in place to prevent excessive usage, and it is important to respect those limits to avoid getting banned or blocked. Since, I dont know the policies of data feed provider, I can't be more accurate in my decision.

The importance of this approach is illustrated by the hypothetical case of Leo Messi returning to FC Barcelona. If our cron job is set to run every 30 min or 2 hours, there is a possibility that the news about Messi's return might not appear in our database, which would impact badly in our traffic and customer experience, and we dont want that in incrowd. To avoid this, we need to periodically update our database with fresh data from the provider feed API.

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

# REST API

| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/teams/{team}/news`                             | Retrieve all articles.           |
| `GET`    | `/teams/{team}/news/{id}`                        | Retrieve one article.            |


The REST API to the example app is described below.

## Get list of Articles

### Request where t94 is param {team}

curl --location --request GET 'localhost:4007/teams/t94/news'

### Response

{
    "status": "succes",
    "data": [
        {
            "id": "63fc20de5acc731fb99f0290",
            "teamId": "t94",
            "optaMatchID": "g2300117",
            "title": "Bristol City 1-0 Hull City",
            "type": [
                "Match Reports"
            ],
            "teaserText": "Hull City’s winless run extended to four matches following a narrow defeat at improving Bristol City.",
            "content": "<p>Substitute Nahki Wells&rsquo; confidently taken 70<sup>th</sup>-minute penalty, after Cyrus Christie was adjudged to have handled Anis Mehmeti's cross, stretched the Robins&rsquo; unbeaten run to 12 games in all competitions.</p>\n<p>January loan signings Karl Darlow and Malcolm Ebiowei were handed their first starts in black and amber as Liam Rosenior made three changes from the goalless draw against Preston North End.</p>\n<p>Regan Slater also came into the side, while Adama Traor&eacute; was named in a matchday squad for the first time since joining in the summer following a long injury lay-off.</p>\n<p>Ryan Longman dropped to the bench, with Matt Ingram and Dimitrios Pelkas absent through injury.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/first-team-images-2022-23/match-galleries/g34---bristol-city-a-250223/g34-bristol-city-a-feb-2023-03.jpg\" alt=\"G34 Bristol City A Feb 2023 03.JPG\" width=\"1360\" height=\"1278\" loading=\"lazy\"></p>\n<p>Tigers chief Rosenior started his playing career at Ashton Gate, inspiring the Robins to EFL Trophy glory in 2003 with the second goal in a 2-0 victory over Carlisle United.</p>\n<p>The erstwhile full-back was brought to Hull City and handed his England Under-21 debut at the MKM Stadium by Bristol City boss Nigel Pearson, who celebrated two years in charge this week.</p>\n<p>In the first managerial meeting between the two, Pearson&rsquo;s charges held the upper hand in a dominant opening quarter of an hour.</p>\n<p>Debutant Darlow was called into action inside the first 40 seconds to block Mark Sykes&rsquo; low drive with his foot before Sean McLoughlin deflected a George Tanner effort over the bar.</p>\n<p>The visitors were looking shaky and Darlow atoned for a sloppy pass out from the back by getting down low to palm away Sam Bell&rsquo;s driven cross with Alex Scott sliding in.</p>\n<p>The Newcastle loanee then claimed a Scott centre from the byline after a surging 60-yard run from the highly-rated teenager.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/first-team-images-2022-23/match-galleries/g34---bristol-city-a-250223/g34-bristol-city-a-feb-2023-10.jpg\" alt=\"G34 Bristol City A Feb 2023 10.JPG\" width=\"1359\" height=\"1032\" loading=\"lazy\"></p>\n<p>Having weathered the early storm, tenacious play from Slater set up a shooting chance for &Oacute;scar Estupi&ntilde;&aacute;n on 25 minutes but he dragged his left-footed drive wide from 20 yards.</p>\n<p>That was the Tigers&rsquo; only meaningful effort of a first 45 which fizzled out after a lively start from the hosts, with chances remaining at a premium in the second half.</p>\n<p>There were fine blocks at both ends as Christie&rsquo;s shot was charged down, while McLoughlin and Alfie Jones stood in the way of efforts from Sykes and Mehmeti.</p>\n<p>In a match of precious few opportunities, Bristol City were given a glorious chance to take the lead on 70 minutes when Mehmeti&rsquo;s cross struck Christie's hand inside the box.</p>\n<p>Wells stepped up, having converted the Robins&rsquo; first spot-kick in 469 days in their previous match at Sunderland, and emphatically drilled home right footed, sending Darlow the wrong way.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/first-team-images-2022-23/match-galleries/g34---bristol-city-a-250223/g34-bristol-city-a-feb-2023-09.jpg\" alt=\"G34 Bristol City A Feb 2023 09.JPG\" width=\"1360\" height=\"813\" loading=\"lazy\"></p>\n<p>That setback almost brought an instant riposte; first, Longman&rsquo;s back-post strike from an acute angle was blocked by the legs of Max O&rsquo;Leary before Benjamin Tetteh nodded a Lewie Coyle corner onto the post.</p>\n<p>With the match now more open, Darlow spilled a shot from distance before Traor&eacute; was sent on for his Tigers debut with 12 minutes left but there was to be no equaliser.</p>\n<p>&nbsp;</p>\n<p><strong><em>Bristol City (4-2-3-1):</em></strong><em> O&rsquo;Leary; Tanner (Cornick 71&rsquo;), Vyner, Kalas, Pring; James &copy;, Williams; Sykes, Scott (Weimann 80&rsquo;), Mehmeti (Dasilva 80&rsquo;); Bell (Wells 62&rsquo;).</em></p>\n<p><em>Subs not used: Haikin, King, Taylor-Clarke.</em></p>\n<p>&nbsp;</p>\n<p><strong><em>Hull City (4-2-3-1):</em></strong><em> Darlow; Christie, A Jones, McLoughlin, Greaves &copy;; Simons, Docherty (</em><em>Traor&eacute; 78&rsquo;)</em><em>, Ebiowei (</em><em>Tetteh 55&rsquo;)</em><em>, Tufan (Longman 46&rsquo;), </em><em>Slater (</em><em>Coyle 65&rsquo;)</em><em>; Estupi&ntilde;&aacute;n.</em></p>\n<p><em>Subs not used</em><em>: Lo-Tutala, Elder, Figueiredo.</em></p>\n<p>&nbsp;</p>\n<p><em><strong>Attendance:</strong> 20,333</em></p>\n<p style=\"text-align: center;\"><iframe src=\"https://www.youtube.com/embed/56KO6Cv2ch4\" frameborder=\"0\" width=\"560\" height=\"315\" allowfullscreen=\"allowfullscreen\" title=\"YouTube video player\" loading=\"lazy\"></iframe></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/bristol-city-v-hull-city/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/bristol-city-v-hull-city/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/e1bc8196-2fa9-47a1-99ef-6c19d3192773/Medium/g34-bristol-city-a-feb-2023-01.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/48bb5ab7-2c91-4201-97ab-1c70a4a0d766/Medium/g34-bristol-city-a-feb-2023-02.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/2854aef6-abbc-4b5b-b83f-f54630b727b0/Medium/g34-bristol-city-a-feb-2023-04.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/73a8784d-90c7-4aee-aefb-505654073c6a/Medium/g34-bristol-city-a-feb-2023-05.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/4e96afd0-61d8-4a6b-b9ed-a40ea3e7235a/Medium/g34-bristol-city-a-feb-2023-06.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/6ce4090e-db3d-4a85-a82b-a9afbc7b7d34/Medium/g34-bristol-city-a-feb-2023-07.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/785fcb27-fc9b-4f78-b6fe-d459ff1b9afb/Medium/g34-bristol-city-a-feb-2023-08.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/8093ea07-9730-495f-a637-d8ce66af870c/Medium/g34-bristol-city-a-feb-2023-11.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/dfc19055-9164-4536-82c6-1046dd401ee5/Medium/g34-bristol-city-a-feb-2023-12.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/2be72229-fd5c-43c1-8215-b874efbbb2c1/Medium/g34-bristol-city-a-feb-2023-13.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-25T13:30:35Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f028f",
            "teamId": "t94",
            "optaMatchID": "g2300117",
            "title": "Rosenior’s Bristol City Reaction",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior felt a sluggish start set the tone for Hull City’s narrow 1-0 loss at Bristol City.",
            "content": "<p>Substitute Nahki Wells&rsquo; second-half penalty, after Cyrus Christie was adjudged to have handled Anis Mehmeti's cross, settled a match of few chances at Ashton Gate.</p>\n<p>&ldquo;Our performance didn&rsquo;t deserve any luck,&rdquo; said Rosenior.</p>\n<p>&ldquo;We gave them a leg up in the first 10 minutes, played with a lack of energy, which is something I haven&rsquo;t seen.</p>\n<p>&ldquo;That&rsquo;s why you end up losing the game, if you start a game of football like that and you lose so many headers, so many tackles, so many second balls.</p>\n<p>&ldquo;You give the opposition confidence and energy, and we did that. We lost the game in the first half not the second.</p>\n<p>&ldquo;The two centre-halves were magnificent &ndash; that was of the only bright spots, the other was Adama (Traor&eacute;) coming on for the first time in a long time.</p>\n<p>&ldquo;They&rsquo;ve been magnificent but it&lsquo;s not just been down to them &ndash; the team worked really hard before this game, for us to be disciplined, organised, hard to beat.</p>\n<p>&ldquo;Too many box entries, too many times play was breaking down and they were breaking on us.</p>\n<p>&ldquo;We didn&rsquo;t deserve anything from the game.&rdquo;</p>\n<p>Rosenior confirmed that Dimitrios Pelkas has a minor groin injury but could return for West Bromwich Albion, while Matt Ingram is likely to be out for two weeks after tweaking his hamstring.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-bristol-city-reaction/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-bristol-city-reaction/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-25T21:23:37Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0291",
            "teamId": "t94",
            "optaMatchID": "g2300117",
            "title": "Team News: Bristol City (A)",
            "type": [
                "Club News"
            ],
            "teaserText": "January loan signings Karl Darlow and Malcolm Ebiowei have been handed their first starts in Hull City colours for the trip to in-form Bristol City.",
            "content": "<p>Regan Slater also comes in as Liam Rosenior makes three changes from the goalless draw against Preston North End.</p>\n<p>Ryan Longman is named on the bench, with Matt Ingram and Dimitrios Pelkas absent.</p>\n<p>Adama Traor&eacute; is included in a matchday squad for the first time since his summer move having recovered from his long injury lay-off.</p>\n<p>&nbsp;</p>\n<p><em><strong>Hull City:</strong> Darlow; Christie, A Jones, McLoughlin, Greaves &copy;; Simons, Docherty, Ebiowei, Tufan, Slater; Estupi&ntilde;&aacute;n.</em></p>\n<p><em>Substitutes: Lo-Tutala, Coyle, Elder, Figueiredo, Longman, Traor&eacute;, Tetteh.</em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/team-news-bristol-city-a/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/team-news-bristol-city-a/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-25T13:29:49Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0292",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Put your thoughts to the Supporters' Committee",
            "type": [
                "Club News"
            ],
            "teaserText": "Supporters have the chance to put their questions to this season's Supporters' Committee before the next meeting on Wednesday 26th April.",
            "content": "<p>If you have a question for one of the members or have an issue you would like to raise, please contact the relevant committee member below:</p>\n<div>\n<div class=\"adition-container\">\n<div id=\"adition-instance-4558297\" class=\"adition-instance adition-outstream\" data-adition-group=\"outstream\">&nbsp;</div>\n<div class=\"adition-instance adition-outstream\" data-adition-group=\"outstream\"><strong>Supporters' Committee</strong></div>\n</div>\n</div>\n<p>North Stand Representatives: <a href=\"mailto:Northstand@wearehullcity.co.uk\">Northstand@wearehullcity.co.uk</a></p>\n<p>Chris Chilton Stand Representatives:&nbsp;<a href=\"mailto:chrischilternstand@wearehullcity.co.uk\">chrischiltonstand@wearehullcity.co.uk</a></p>\n<p>South Stand Representatives:&nbsp;<a href=\"mailto:southstand@wearehullcity.co.uk\">southstand@wearehullcity.co.uk</a></p>\n<p>West Stand Representatives:&nbsp;<a href=\"mailto:weststandrep@gmail.com\"></a><a href=\"mailto:weststand@wearehullcity.co.uk\">weststand@wearehullcity.co.uk</a></p>\n<p>&nbsp;</p>\n<p><strong>Junior Supporters' Committee</strong></p>\n<p>North Stand Representatives: <a href=\"mailto:Northstand@wearehullcity.co.uk\">juniornorthstand@wearehullcity.co.uk</a></p>\n<p>Chris Chilton Stand Representatives: <a href=\"mailto:juniorchrischiltonstand@wearehullcity.co.uk\">juniorchrischiltonstand@wearehullcity.co.uk</a></p>\n<p>South Stand Representatives: <a href=\"mailto:juniorsouthstand@wearehullcity.co.uk\">juniorsouthstand@wearehullcity.co.uk</a></p>\n<p>West Stand Representatives: <a href=\"mailto:juniorweststand@wearehullcity.co.uk\">juniorweststand@wearehullcity.co.uk</a></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/put-your-thoughts-to-the-supporters-committee/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/put-your-thoughts-to-the-supporters-committee/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-24T15:59:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0295",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "20% off our 2022/23 Away Kit!",
            "type": [
                "Club News"
            ],
            "teaserText": "Starting from Friday 24th February, there is a 20% discount on our 22/23 Away Kit, available online and in-store!",
            "content": "<p>Supporters can get their hands on the iconic white away shirt this weekend, with adult shirts reduced from &pound;49.99 to &pound;39.99 and junior shirts down from &pound;34.99 to &pound;27.99.</p>\n<p>Press the button below to shop today!</p>\n<table class=\"scroll-table\" style=\"width: 100%; border-collapse: collapse; border-style: hidden; margin-left: auto; margin-right: auto; height: 18px;\" border=\"0\">\n<tbody>\n<tr style=\"height: 18px;\">\n<td style=\"width: 42.7889%; background-color: #ffffff; height: 18px;\">&nbsp;</td>\n<td style=\"width: 17.8362%; border-style: solid; border-color: #e77400; background-color: #e77400; color: #ffffff; text-align: center; vertical-align: middle; height: 18px;\"><a href=\"https://www.tigerleisure.com/replicakit/away-kit-2223/?utm_source=Story&amp;utm_medium=Website&amp;utm_campaign=Away+Shirt+Discount\"><strong>SHOP NOW</strong></a></td>\n<td style=\"width: 39.3748%; background-color: #ffffff; height: 18px;\">&nbsp;</td>\n</tr>\n</tbody>\n</table>\n<p>&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/20-off-our-2223-away-shirt/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/20-off-our-2223-away-shirt/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-24T11:59:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0297",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "No pay on the day at Bristol",
            "type": [
                "Ticket News"
            ],
            "teaserText": "Tickets for Hull City’s Sky Bet Championship fixture away at Bristol City will not be made available at Ashton Gate - meaning supporters must have purchased tickets in advance to attend.",
            "content": "<p>The game kicks off at 3pm at Ashton Gate on Saturday 25th February.</p>\n<p>Tickets are available over the phone at 01482 505600 or from Tiger Leisure until 12 noon today (Friday 24th February).</p>\n<p><strong>Ticket Prices:</strong></p>\n<p>&pound;28 - Adults</p>\n<p>&pound;25 - Senior (Over 65), Under 25</p>\n<p>&pound;21 - Under 22</p>\n<p>&pound;15 - Under 19&nbsp;</p>\n<p>&pound;10 - Under 12&nbsp;</p>\n<p>Please do not travel down to Bristol City without a ticket.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/no-pay-on-the-day-at-bristol/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/no-pay-on-the-day-at-bristol/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-24T10:29:56Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0293",
            "teamId": "t94",
            "optaMatchID": "g2300117",
            "title": "Greg Docherty Pre-Bristol City (A) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Greg Docherty spoke to the media ahead of the Tigers' trip to Ashton Gate.",
            "content": "<p>The Scottish midfielder revealed his pride after reaching 300 career appearances in the goalless draw with Preston and shared his thoughts on the challenge in-form Bristol City will pose.</p>\n<p>Watch his pre-match press conference above!</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/greg-docherty-pre-bristol-city-a-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/greg-docherty-pre-bristol-city-a-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-24T12:37:17Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0294",
            "teamId": "t94",
            "optaMatchID": "g2300117",
            "title": "Liam Rosenior Pre-Bristol City (A) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior spoke to the media ahead of Hull City's Championship match away to Bristol City.",
            "content": "<p>The 38-year-old discussed recent results, injury news, his playing days at Ashton Gate and relationship with his former Tigers boss Nigel Pearson.</p>\n<p>Watch his pre-match press conference above!&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/liam-rosenior-pre-bristol-city-a-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/liam-rosenior-pre-bristol-city-a-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-24T12:36:49Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0296",
            "teamId": "t94",
            "optaMatchID": "g2300117",
            "title": "Traoré set to be involved against Bristol City",
            "type": [
                "Match Previews"
            ],
            "teaserText": "Liam Rosenior has revealed that Adama Traoré will be in the squad for this weekend’s trip to Bristol City.",
            "content": "<p style=\"font-weight: 400;\">The 27-year-old midfielder is yet to feature for the Tigers having been ruled out with a hamstring injury since joining from Hatayspor in September.</p>\n<p style=\"font-weight: 400;\">&ldquo;He played a big part on Tuesday in terms of our behind closed doors game (against Manchester City),&rdquo; revealed Rosenior.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;ll definitely be in the squad. He&rsquo;s a very good player. What I need to gauge with Adama is his robustness for the league which I&rsquo;ve had to do with a lot of players since I&rsquo;ve come in, especially the ones who aren&rsquo;t used to the Championship.</p>\n<p style=\"font-weight: 400;\">&ldquo;What I&rsquo;ve seen from him in training and in the two games he&rsquo;s played for the Under-21s and the one game with us against Manchester City is that he&rsquo;s a very good player and he&rsquo;s fit.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s not match fit, but he&rsquo;s fit enough to be in our squad, and he&rsquo;ll definitely travel with us down to Bristol.&rdquo;</p>\n<p style=\"font-weight: 400;\">Allahyar Sayyadmanesh is also back in full training with the rest of the squad, but Rosenior insists he will not risk the 21-year-old forward after he broke down on his last return to the side.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s training fully,&rdquo; explained Rosenior.</p>\n<p style=\"font-weight: 400;\"><img src=\"https://www.wearehullcity.co.uk/contentassets/fed16d54a589432d858ddd6cfeebd779/imagexrhhf.png\" alt=\"\" width=\"4608\" height=\"3072\" loading=\"lazy\"></p>\n<p style=\"font-weight: 400;\">&ldquo;What I will not do is risk him. I will not risk a player&rsquo;s fitness for our short-term benefit. I have a discussion with the medical staff in terms of what we do with Allahyar next week.</p>\n<p style=\"font-weight: 400;\">&ldquo;We have a game against Burnley (Under-21s on Monday) and he could be involved in that. We haven&rsquo;t made a decision yet. He won&rsquo;t be available for Saturday.&rdquo;</p>\n<p style=\"font-weight: 400;\">Rosenior also provided an update on Aaron Connolly, who has been missing from the squad since being withdrawn against Stoke City.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s so frustrated and angry but he loves it here, the players love having him here and the staff love having him here,&rdquo; added Rosenior.</p>\n<p style=\"font-weight: 400;\">&ldquo;Hopefully, he&rsquo;ll be back sooner rather than later.&rdquo;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/traore-set-to-be-involved-against-bristol-city/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/traore-set-to-be-involved-against-bristol-city/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-24T10:50:48Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0298",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Ashbee on playing for ROI Under-17s",
            "type": [
                "Academy"
            ],
            "teaserText": "Hull City Academy prospect Stan Ashbee was delighted to earn his first caps for Republic of Ireland Under-17s earlier this month.",
            "content": "<p>The defender was part of a 22-man squad selected by Colin O&rsquo;Brien to represent the nation as they played a friendly double header against Hungary at the Pinatar Arena in Spain.</p>\n<p>Ashbee featured in both games, coming on as a second-half substitute &ndash; a 1-0 win courtesy of a finish from Shamrock Rovers&rsquo; Naj Razi.</p>\n<p>Ashbee then started the second match, playing the full 90 minutes as he helped keep a clean sheet on his first start for O&rsquo;Brien&rsquo;s side.</p>\n<p>&ldquo;The camp was a good opportunity for me to get my first minutes as an Under-17,&rdquo; said Ashbee. &ldquo;They were two good games with two clean sheets and good performances from all the lads.</p>\n<p>&ldquo;Playing international football is different. It is harder. There are new challenges and there is something that you can always improve on in games so it was a good experience for me.&rdquo;</p>\n<p>Son of former Hull City captain and Hall of Fame member Ian Ashbee, Stan is thankful to his dad for attending all the camps he has previously featured in for Ireland.</p>\n<p>&ldquo;My dad will always try and get to all the camps that he can,&rdquo; he added. &ldquo;I don&rsquo;t think he has missed one yet so he is always there for me after the games to talk about how I have played and what I can improve on which is great.&rdquo;</p>\n<p>Ashbee was also quick to praise Under-18s assistant manager Billy Clarke for his role in advising the defender when attending the camps, with Clarke featuring for Republic of Ireland at Under-17, Under-19 and Under-21 levels in his playing career.</p>\n<p>&ldquo;Billy is good. He knows when I go on the camps, he will always come to me for a chat beforehand. We talk about how I am doing and where we are going,&rdquo; explained Ashbee.</p>\n<p>&ldquo;Before my first camp, he gave me a talking to before saying that everything would be alright and to just enjoy the experience.&rdquo;</p>\n<p>The double header against Hungary was the final friendly fixtures for O&rsquo;Brien&rsquo;s side as they look forward to playing in their Euro qualifiers next month in Cyprus.</p>\n<p>Playing in Group 6 along with Italy, Ukraine and hosts Cyrpus, Ashbee is hoping to be involved in the upcoming squad.</p>\n<p>&ldquo;It will mean everything to me to be selected to go play in the Euros. I just want to play and help the lads and get as far as we can in the competition.&rdquo;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/ashbee-on-playing-for-roi-under-17s/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/ashbee-on-playing-for-roi-under-17s/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-22T17:58:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f0299",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Watch Hull City from the Best Seat in the House!",
            "type": [
                "Community"
            ],
            "teaserText": "As part of the Tigers Trust’s upcoming matchday takeover, the charity is offering you the chance to watch Hull City’s game against West Bromwich Albion from the Best Seat in the House.",
            "content": "<p>Courtesy of MKM, this great prize includes an amazing perspective of the match from the West Stand Upper, as well as a huge basket of food to get stuck into before kick-off.</p>\n<p>This fixture, taking place on Friday 3<sup>rd</sup> March, will be centred around the theme of &lsquo;Inspired Communities&rsquo; and showcase the Trust&rsquo;s dedicated work around inclusivity.</p>\n<p>Entry to the Best Seat in the House competition is free (text messages cost your standard network message charge), and entrants will be asked to make a voluntary donation to the Tigers Trust. &nbsp;</p>\n<p>Donations from this competition will go towards supporting more children, young people and adults in the community through the Tigers Trust&rsquo;s Health and Wellbeing Fund.</p>\n<p>To enter this competition, please text <strong>TRUST</strong> to 70215.</p>\n<p>This competition will close on Sunday 26<sup>th</sup> February, with the winner contacted the following day.</p>\n<p>To find out more about the Best Seat in the House competition and to view the terms and conditions, please <a href=\"https://www.tigerstrust.co.uk/news/competition-best-seat-in-the-house/\">click here.</a></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/watch-hull-city-from-the-best-seat-in-the-house/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/watch-hull-city-from-the-best-seat-in-the-house/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-22T10:44:54Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f029a",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Pugh delivers academy workshop",
            "type": [
                "Academy"
            ],
            "teaserText": "Hull City were delighted to welcome back former midfielder Marc Pugh to speak to the Under-21s about healthy eating and nutrition.",
            "content": "<p>Part of the life skills programme from the LFE (League Football Education), Pugh delivered a workshop at Bishop Burton College for Conor Sellars&rsquo; side to help them in their future careers.</p>\n<p>&ldquo;I go around and deliver workshops for the 72 EFL clubs and I am an online health coach and I absolutely love it,&rdquo; said Pugh.</p>\n<p>&ldquo;It is nice to give something back and educate the young lads on the importance of nutrition. It is a really important part for me adding value; it is like scoring a goal.</p>\n<p>&ldquo;Since retiring from football, having my why in life is really important and to see the lads asking questions is absolutely brilliant.</p>\n<p>&ldquo;I think nutrition is an important part of life, not just for an athlete but for everyone. The healthier we are, the better we feel on a day-to-day basis.&rdquo;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/223884d9c510479b8a2145e7d762ca23/marc-pugh-workshop-1-min.jpg\" alt=\"\" width=\"5909\" height=\"3939\" loading=\"lazy\"></p>\n<p>The 35-year-old scored three goals in 14 appearances during a short-term loan spell with the club in 2019.</p>\n<p>Despite his time at the Tigers being short, the retired midfielder enjoyed returning to the club and educate the future generation of Hull City.</p>\n<p>&ldquo;I love Hull; I had some lovely times here with a great set of fans,&rdquo; added Pugh.</p>\n<p>&ldquo;I loved playing for the club even though it was for a short spell under Nigel Adkins. He made me feel so welcome and it was a really nice environment to play in.</p>\n<p>&ldquo;I made a lot of special friends and memories here so it was really nice to come back.&rdquo;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/223884d9c510479b8a2145e7d762ca23/marc-pugh-v-millwall-h-2018.jpg\" alt=\"\" width=\"3000\" height=\"2064\" loading=\"lazy\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/pugh-delivers-academy-workshop/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/pugh-delivers-academy-workshop/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-21T17:10:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b0",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Members Event | David Meyler Q&A and Book Signing",
            "type": [
                "Ticket News"
            ],
            "teaserText": "Hull City will host an exclusive members Q&A & book signing event at the MKM Stadium on Tuesday 28th March, with ex-Tiger David Meyler.",
            "content": "<p>Alex Burgess will host the evening, and attendees will have the opportunity to put their questions to the former Hull City midfielder.</p>\n<p>Meyler won two promotions during his stint with the Tigers, as well as starting in the club's appearance in the 2014 FA Cup Final.&nbsp;</p>\n<p>The evening will take place in the Kingston Suite with doors opening from 6pm for a 7pm start. Tickets will be <strong>FREE</strong>&nbsp;and limited to two per Member.</p>\n<p>The bar in the Kingston Suite will be open for the evening to serve drinks, meals and snacks.</p>\n<p>Tickets will become available at 9.30am on Monday 27th February <a href=\"https://wearehullcity.talent-sport.co.uk/PagesPublic/ProductBrowse/productEvent.aspx?ProductSubtype=MBEV\">online here</a>&nbsp;or over the phone at 01482 505600.</p>\n<p>Tickets are expected to sell out fast, so be ready on the phone or laptop for Monday 27th February!</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/members-event--david-meyler-qa-and-book-signing/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/members-event--david-meyler-qa-and-book-signing/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T14:00:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f029b",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Loan Watch: Trio make debuts",
            "type": [
                "Club News"
            ],
            "teaserText": "Hull City have 14 players out on loan. Here we take a look at how our Tigers are getting on in loan watch…",
            "content": "<p>&nbsp;</p>\n<p><strong>Louie Chorlton &amp; Josh Hinds at Bradford Park Avenue: </strong></p>\n<p>We start with our most recent departures as on-loan duo Louie Chorlton and Josh Hinds made their debuts for Bradford Park Avenue.</p>\n<p>Despite Will Longbottom netting the opener, the Green Army suffered a 4-1 away defeat against promotion-chasing Brackley Town. Chorlton played the full game while Hinds featured as a second-half substitute, coming on in the 63<sup>rd</sup> minute.</p>\n<p><strong>McCauley Snelgrove at Spennymoor Town: </strong></p>\n<p>Another recent loan departure was McCauley Snelgrove, who made his Spennymoor Town debut as a second-half substitute in a 2-2 draw against Scarborough Athletic.</p>\n<p>Introduced in the 69<sup>th</sup> minute, the winger played in front of 1,472 fans who attended the four-goal thriller. Spennymoor sit 12<sup>th</sup> in the National League North table.</p>\n<p><strong>Tom Nixon &amp; Billy Chadwick at Boston United: </strong></p>\n<p>Staying in National League North, Tom Nixon and Billy Chadwick both started for Boston United as they suffered a 1-0 defeat against Chester, conceding a 95<sup>th-</sup>minute goal from Declan Weeks.</p>\n<p>Nixon, who was making his seventh appearance for the club, played the full game while Chadwick was replaced after 83 minutes.</p>\n<p><strong>Brandon Fleming &amp; Tyler Smith at Oxford United: </strong></p>\n<p>Our final on-loan duo in Brandon Fleming and Tyler Smith both started as out-of-form Oxford United suffered a 1-0 defeat against Cambridge United. Lloyd Jones&rsquo; eighth-minute finish was enough for struggling Cambridge to secure a much-needed three points in League One.</p>\n<p>Fleming played the full game on his eighth Yellows appearance, while Smith, in his fifth outing for the U&rsquo;s, was replaced at half-time.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/22ddedb911fa4612b7c9945407bf679c/jake-leake-v-cambridge.jpg\" alt=\"\" width=\"3399\" height=\"2262\" loading=\"lazy\"></p>\n<p><strong>Jake Leake at Scunthorpe United: </strong></p>\n<p>Defender Jake Leake helped Scunthorpe United earn a vital three points in their fight against relegation from the National League, defeating Dagenham &amp; Redbridge 3-2 at Glanford Park.</p>\n<p>Leake made his home debut for the Iron as goals from Ben Richards-Everton, Cameron Wilson and Liam Feeney saw Jimmy Dean&rsquo;s side reign victorious.</p>\n<p><strong>Jevon Mills at Solihull Moors: </strong></p>\n<p>Staying in the National League, Jevon Mills made his third appearance for Solihull Moors in their goalless draw against fourth-placed Chesterfield.</p>\n<p>Playing the full 90 minutes in his three appearances for the Moors, the defender will be looking to earn his fourth consecutive start tomorrow night as Neal Ardley&rsquo;s side travel to FC Halifax Town.</p>\n<p><strong>Ben Voase at Beverley Town: </strong></p>\n<p>Second-year scholar Ben Voase made his ninth appearance for Beverley Town as they travelled to The Welfare Ground to face Nostell Miners Welfare.</p>\n<p>Despite saving a penalty in the seventh minute from Joe Wood, Voase could not help prevent the Beavers from suffering a 3-0 defeat in the Northern Counties East League Division One.</p>\n<p><strong>Harvey Cartwright at Wycombe Wanderers: </strong></p>\n<p>Fellow goalkeeper Harvey Cartwright was an unused substitute for Wycombe Wanderers as they defeated Bolton Wanderers 1-0.</p>\n<p>Lewis Wing&rsquo;s finish eight minutes into first-half stoppage-time made it five consecutive wins for the League One play-off hopefuls. Cartwright has been an unused substitute in four games for the Chairboys.</p>\n<p><strong>Andy Smith at Grimsby Town: </strong></p>\n<p>In the league below, Andy Smith was an unused substitute in Grimsby&rsquo;s most recent fixture, a 2-1 win over Northampton Town.</p>\n<p>Anthony Glennon&rsquo;s stoppage-time finish saw the Mariners secure the three points in League Two, with ex-Tiger Josh Emmanuel registering both assists for Grimsby in the fixture.</p>\n<p><strong>Yuriel Celi at Club Universitario de Deportes: </strong></p>\n<p>Finally in Peru, Yuriel Celi wasn&rsquo;t named in the team for Universitario as they suffered a 1-0 defeat against Club Alianza Lima.</p>\n<p>Despite being named in the initial squad heading into the fixture, Celi was not featured either in the starting XI or substitutes bench.</p>\n<p>*<strong><em>Doğukan Sinik </em></strong><em>has not featured for Antalyaspor due to the events in T&uuml;rkiye. The S&uuml;per Lig season is scheduled to return on the 3<sup>rd</sup> of March. </em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/loan-watch-trio-make-debuts/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/loan-watch-trio-make-debuts/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-20T17:00:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02aa",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Matchday takeover to raise crucial funds for the Tigers Trust",
            "type": [
                "Community"
            ],
            "teaserText": "The Tigers Trust, the affiliated charity of Hull City, is set to host its second matchday takeover in as many years when the Tigers host West Bromwich Albion at the MKM Stadium on Friday 3rd March.",
            "content": "<p>This fixture, which is to be televised on Sky Sports, will highlight the theme of &lsquo;Inspired Communities&rsquo; and showcase the Trust&rsquo;s dedicated work around inclusivity.&nbsp;</p>\n<p>Inclusion is at the heart of everything that the Tigers Trust delivers, and staff are committed to making sure that activities can be accessed and enjoyed by all.</p>\n<p>The 2021/22 season saw just over 43,000 children, young people and adults engage with the Trust.</p>\n<p>This includes 16,775 unique individuals through the Trust&rsquo;s various programmes in schools and communities, and another 26,272 who accessed a range of activities at the Tigers Trust Arena.</p>\n<p>Further details of the Tigers Trust&rsquo;s plans for the day will be released across the charity&rsquo;s website and social media channels.</p>\n<p>If you would like to donate &pound;3 to the Tigers Trust, please text TIGERSTRUST to 70580. Donations to the Tigers Trust can also be made by <a href=\"https://wonderful.org/pay?ref=1092287\">clicking here.</a></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/matchday-takeover-to-raise-crucial-funds-for-the-tigers-trust/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/matchday-takeover-to-raise-crucial-funds-for-the-tigers-trust/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/fd934fcc-3d01-45a1-8f4d-8326775f9d3b/Medium/tt-matchday-takeover-feb-2023.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/ef423dc8-6ebe-4550-87b2-01b2c5920b03/Medium/tt-matchday-takeover-greaves-feb-2023.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-15T13:56:11Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f029d",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "Hull City 0-0 Preston North End",
            "type": [
                "Match Reports"
            ],
            "teaserText": "Hull City made it five league games unbeaten at the MKM Stadium as the Tigers played out a goalless draw against Preston North End.",
            "content": "<p>Preston goalkeeper Freddie Woodman produced a string of impressive saves to deny the Tigers as the points were shared in the EFL Championship.</p>\n<p>Head Coach Liam Rosenior made four changes to his starting XI to the one that faced Norwich City last time out. Dimitrios Pelkas, Ozan Tufan, Xavier Simons and Greg Docherty all came into the side for Ryan Woods, Ryan Longman, Benjamin Tetteh and Jean Micha&euml;l Seri. Woods later was injured in the warm-up and replaced on the substitutes bench by Tobias Figueiredo.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/b59b62e5d5f74b89b797553788b79693/minutes-silence-v-pne-h-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></p>\n<p>The Tigers crafted the first opportunities of the game as Tufan fired high and wide with his first-time shot, after a good initial run by Pelkas, before the Turkish midfielder tried his luck again, with the result being the same outcome.</p>\n<p>Pelkas almost managed to find &Oacute;scar Estupi&ntilde;&aacute;n after the Greek midfielder dispossessed Jordan Storey and chipped the onrushing Woodman, but Estupi&ntilde;&aacute;n couldn&rsquo;t get on the end of Pelkas&rsquo; delivery.</p>\n<p>At the other end, Liam Delap drove into the Tigers&rsquo; penalty area and rifled his powerful right-footed strike just wide, hitting the side netting.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/b59b62e5d5f74b89b797553788b79693/ozan-tufan-v-pne-h-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1667\" loading=\"lazy\"></p>\n<p>The Tigers felt they should have had a penalty as Tufan was played in on goal and went down under the challenge of Woodman, but the referee adjudged the midfielder to have dived. Moments later, Cyrus Christie drilled his low effort narrowly wide past the right-hand post.</p>\n<p>With five minutes of the first 45 to play, Longman threaded the ball through for Estupi&ntilde;&aacute;n, but the Colombian forward saw his attempt pushed away by Woodman. From the resulting corner by Pelkas, Docherty headed wide.</p>\n<p>The Scottish midfielder almost found an opener on the stroke of half-time as his charging run caused panic in the Preston penalty area, but Woodman was on hand to save following a deflection off substitute Liam Lindsay.</p>\n<p>Two minutes into the second-half, Woodman produced a brilliant reaction save with his feet to deny Alfie Jones from Pelkas&rsquo; corner before the ball worked its way back out for Pelkas, who&rsquo;s out-of-the-foot strike was held by the Preston goalkeeper.</p>\n<p>The visitors made a double substitution at half-time, with one of them almost finding the opener on the 54<sup>th</sup> minute. Tom Cannon drilled a low cross for Troy Parrott who in-turn directed his effort goalwards, but the on-loan Spurs&rsquo; forward fired narrowly wide.</p>\n<p>Preston started to build in momentum after Parrott&rsquo;s chance, with Greg Cunningham going close with a diving header before Dan Potts saw his effort deflect into the side netting off Jacob Greaves.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/b59b62e5d5f74b89b797553788b79693/malcolm-ebiowei-v-pne-h-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></p>\n<p>It was end-to-end action as Greaves&rsquo; floating delivery saw a back peddling Woodman flick the ball behind for a corner. The Preston shot-stopper was on hand again minutes later to punch away Tufan&rsquo;s attempt at the near post, after a good run from the Turkish midfielder.</p>\n<p>With quarter-of-an-hour to go, Slater blocked well to deny Parrott before Preston captain Alan Browne fired over from distance.</p>\n<p>Malcolm Ebiowei then picked out the run of Slater, but the substitutes&rsquo; header was blocked by Cunningham before Simons&rsquo; ambitious strike deflected behind off a Preston man in injury-time.</p>\n<p>&nbsp;</p>\n<p><strong><em>Hull City (4-3-3): </em></strong><em>Ingram; Christie, A. Jones, McLoughlin, Greaves &copy;; Tufan, Docherty, Simons; Longman (Ebiowei 65&rsquo;), Estupi&ntilde;&aacute;n (Tetteh 75&rsquo;), Pelkas (Slater 65&rsquo;). </em></p>\n<p><strong><em>Subs Not Used: </em></strong><em>Darlow, Coyle, Elder, Figueiredo.</em></p>\n<p><em>&nbsp;</em></p>\n<p><strong><em>Preston North End (3-4-1-2): </em></strong><em>Woodman; Storey (Diaby 46&rsquo;), Cunningham, Hughes (Lindsay 36&rsquo;); Potts, Browne &copy;, Ledson, Fern&aacute;ndez; Woodburn (Johnson 74&rsquo;); Delap (Parrott 46&rsquo;), Cannon. </em></p>\n<p><strong><em>Subs Not Used: </em></strong><em>Cornell, Brady, Onomah. &nbsp;</em></p>\n<p><em>&nbsp;</em></p>\n<p><strong><em>Attendance: </em></strong><em>17,776</em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/hull-city-0-0-preston-north-end/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/hull-city-0-0-preston-north-end/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/0c8107e5-8466-40cd-99b0-65318cbcab27/Medium/g33-pne-h-001.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/ac95745a-3545-47a2-a94c-2879d79bc1d7/Medium/g33-pne-h-002.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/36e79b15-e7ec-4c42-a3f3-5bec6eced65d/Medium/g33-pne-h-003.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/fa564473-262c-4ba1-8b82-664c2dd5a4e4/Medium/g33-pne-h-004.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/98e32c3a-ca05-4a1e-b4fa-109f4f59280c/Medium/g33-pne-h-005.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/04ffbced-ffd8-4d6e-b7bf-36d84e9ef4b1/Medium/g33-pne-h-006.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/cd58e107-2b0d-4f68-b0e7-fda3d9e45f5d/Medium/g33-pne-h-007.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/dc117669-c1f6-4c0e-a4fc-dcf5e0ab98cf/Medium/g33-pne-h-008.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/ec7cf627-0f59-4dd0-ba04-62e76c458ca9/Medium/g33-pne-h-009.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/168fc212-0a82-43ae-b7bd-0207c887cb0d/Medium/g33-pne-h-010.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/56be2733-324a-417a-8ae2-3790e0a4219e/Medium/g33-pne-h-011.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/b6adb23e-8c07-4d90-b123-8eebd81c1364/Medium/g33-pne-h-012.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-18T17:04:38Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02ad",
            "teamId": "t94",
            "optaMatchID": "g2300095",
            "title": "Norwich City 3-1 Hull City",
            "type": [
                "Match Reports"
            ],
            "teaserText": "Jacob Greaves’ fourth goal of the season could not prevent Hull City from falling to defeat at Norwich City.",
            "content": "<p>Despite taking an early lead, goals from Kieran Dowell, Gabriel Sara and Josh Sargent condemned the Tigers to just a second reverse in 11 games.</p>\n<p>Having named the same starting XI for the previous three matches, Liam Rosenior freshened up his side by making four changes for the long Valentine&rsquo;s Day trip to Carrow Road.</p>\n<p>Captain for the night Greaves replaced Callum Elder at left-back, with Ryan Woods and Ryan Longman slotting into central and wide midfield roles in place of Greg Docherty and Ozan Tufan.</p>\n<p>Benjamin Tetteh, back available after serving a three-match suspension for his red card at Sheffield United, was deployed up front as Aaron Connolly missed out through a foot injury.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/first-team-images-2022-23/match-galleries/g32---norwich-a-140223/g32-norwich-a-feb-2023-02.jpg\" alt=\"G32 Norwich A Feb 2023 02.jpg\" width=\"1359\" height=\"826\" loading=\"lazy\"></p>\n<p>A misplaced pass from Jean Micha&euml;l Seri presented Norwich with an early sighter as Sargent&rsquo;s reverse ball freed Onel Hernandez but his off-balance effort from a tight angle rippled the side-netting.</p>\n<p>The Canaries had failed to win in seven at home, losing the last three without scoring, and would have been fearing the worst when the Tigers took a 14<sup>th</sup>-minute lead.</p>\n<p>After Woods&rsquo; corner was not fully cleared, Regan Slater&rsquo;s well-flighted cross back into the danger area was headed up but not away by skipper Grant Hanley, with goalkeeper Angus Gunn flapping.</p>\n<p>Tetteh rose highest to nod down for Greaves, who took a touch before swivelling and firing left footed into the roof of the net.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/first-team-images-2022-23/match-galleries/g32---norwich-a-140223/g32-norwich-a-feb-2023-05.jpg\" alt=\"G32 Norwich A Feb 2023 05.jpg\" width=\"1360\" height=\"1052\" loading=\"lazy\"></p>\n<p>City were eyeing a first double over Norwich in 52 years, having won the reverse fixture 2-1 in August when both teams had different managers, but the lead only lasted four minutes.</p>\n<p>Hernandez outwitted Cyrus Christie to reach the byline and pulled the ball back for Dowell, whose first-time strike deflected off Greaves and looped past a luckless Matt Ingram.</p>\n<p>The Norfolk outfit, needing a win to reignite their play-off ambitions, dominated the rest of the half as the Tigers struggled to keep hold of the ball.</p>\n<p>Dimitris Giannoulis fashioned two chances with crosses from the left; Christie diverted one delivery into the path of Dowell, who scuffed wide, and Sargent headed another too close to Ingram.</p>\n<p>In between, Max Aarons supplied the ammunition from the opposite side, accurately standing up a cross which Adam Idah, outmuscling Sean McLoughlin, headed over the top.</p>\n<p>There was little change after the break as Norwich continued to probe, Hernandez poking the ball through for Idah, whose clipped finish across Ingram curled beyond the far post.</p>\n<p>The pressure told on 58 minutes, Hernandez skilfully holding off his marker before teeing up Sara to strike first time from 20 yards, sweeping a well-placed left-footed shot into the bottom corner.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/first-team-images-2022-23/match-galleries/g32---norwich-a-140223/g32-norwich-a-feb-2023-11.jpg\" alt=\"G32 Norwich A Feb 2023 11.JPG\" width=\"1359\" height=\"908\" loading=\"lazy\"></p>\n<p>City were fortunate not to concede a quick-fire third when Hanley&rsquo;s stooping close-range header from a corner flicked on by Andrew Omobamidele was ruled offside.</p>\n<p>Giannoulis breezed past Christie but his placed finish was gathered by Ingram as Norwich streamed forward with menace.</p>\n<p>A golden chance went begging on 65 minutes as Idah shrugged off McLoughlin and slipped Dowell in behind Greaves but he slotted wide with just Ingram to beat.</p>\n<p>Despite making all five substitutions, the Tigers remained second best and Norwich wrapped up the points with a minute left.</p>\n<p>The excellent Sara exchanged passes with Idah before finding the feet of Sargent, who clinically drilled an unerring drive past Ingram.</p>\n<p>&nbsp;</p>\n<p><em><strong>Norwich City (4-2-3-1): </strong>Gunn; Aarons (Sorensen 91&rsquo;), Omobamidele, Hanley &copy;, Giannoulis (Gibson 91&rsquo;); Sara, McLean; Dowell (Gibbs 72&rsquo;), Sargent, Hernandez (Tzolis 86&rsquo;); Idah.</em></p>\n<p><em>Subs not used: Krul, Pukkis, Nunez.</em></p>\n<p>&nbsp;</p>\n<p><em><strong>Hull City (4-4-2):</strong> Ingram; Christie (Coyle 79&rsquo;), A Jones, McLoughlin, Greaves &copy;; Slater (Tufan 68&rsquo;), Woods (Ebiowei 68&rsquo;), Seri (Docherty 59&rsquo;), Longman; Tetteh, Estupi&ntilde;&aacute;n (Pelkas 59&rsquo;).</em></p>\n<p><em>Subs not used: Darlow, Elder.</em></p>\n<p>&nbsp;</p>\n<p><em><strong>Attendance:</strong> 25,767</em></p>\n<p style=\"text-align: center;\"><iframe src=\"https://www.youtube.com/embed/thxt1fOv7N0\" frameborder=\"0\" width=\"560\" height=\"315\" allowfullscreen=\"allowfullscreen\" title=\"YouTube video player\" loading=\"lazy\"></iframe></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/norwich-city-v-hull-city/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/norwich-city-v-hull-city/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/23adddac-df04-422d-992b-529a3bedecba/Medium/g32-norwich-a-feb-2023-03.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/026c2535-a10c-4f3a-bac9-f69e5747bf38/Medium/sean-mcloughlin-norwich-a-feb-2023.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/48d5f6ef-0c84-443e-8cd8-3c85d9adda0c/Medium/g32-norwich-a-feb-2023-04.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/e5ef4880-06d9-4b02-aa26-e08c2afeb0cb/Medium/g32-norwich-a-feb-2023-06.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/3dfa7dc0-073e-4d93-92e8-72bc75871a79/Medium/g32-norwich-a-feb-2023-07.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/04271bde-8084-477c-a65e-559d32145cae/Medium/g32-norwich-a-feb-2023-08.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/81e133cc-48a9-4b4b-8ef6-a1788181cb80/Medium/g32-norwich-a-feb-2023-09.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/3312ede5-3b53-457b-a9f8-0d67adbe4a14/Medium/g32-norwich-a-feb-2023-10.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/b9786818-2236-4d6b-9eb6-eb4d8fab678a/Medium/g32-norwich-a-feb-2023-12.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T21:30:24Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f029c",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "Rosenior’s Preston Reaction",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior was happy with his side’s build-up play but felt the “final moment” was lacking in the goalless draw with Preston North End.",
            "content": "<p>The Tigers registered 15 attempts at goal but were unable to breach the Lilywhites&rsquo; defence, with the stalemate extending City&rsquo;s unbeaten home run to six matches.</p>\n<p>&ldquo;First half, we were so much in the ascendancy,&rdquo; said Rosenior, whose side have now kept three consecutive clean sheets at the MKM Stadium.</p>\n<p>&ldquo;Our football up to the final third was outstanding. But for some very good saves from their goalkeeper and we missed the final pass maybe four or five times, we could have come in 3-0 up.</p>\n<p>&ldquo;That&rsquo;s the frustrating thing because what I asked from the players was a reaction to Tuesday&rsquo;s performance (at Norwich City). They gave me that.</p>\n<p>&ldquo;Unfortunately for us, we didn&rsquo;t have that final moment, which is really frustrating because 80, 90% of our play was very good.&rdquo;</p>\n<p>Jean Micha&euml;l Seri and Aaron Connolly were both missing from the squad as Xavier Simons came in for his full league debut.</p>\n<p>The boss explained: &ldquo;Aaron, from the knock against Stoke when the ball hit him on the toe, looks like he&rsquo;s going to be out for a month.</p>\n<p>&ldquo;&rsquo;Mika&rsquo; (Seri) has got a groin issue, which we didn&rsquo;t know until late Thursday, early Friday morning. That&rsquo;s a frustrating one for us because Mika&rsquo;s an outstanding player. I expect him to be out until the March international break, the same as Aaron.</p>\n<p>&ldquo;What I do have to say is look at young Xavi (Xavier Simons) who came in and was absolutely magnificent. Xavi's training level since I've been in, his learning, his improvement, has been massive.</p>\n<p>&ldquo;He played the game like an experienced pro. For a midfield player to play against Preston like that at 19 years old was absolutely outstanding but I expect that from him because I think he's going to be a top player.&rdquo;</p>\n<p>The boss also provided a fitness update on Ryan Woods, who was originally named on the bench before being replaced by Tobias Figueiredo, revealing: &ldquo;Woodsy felt his groin kicking a ball in the warm-up. We&rsquo;ll have to scan him and get the results on that.&rdquo;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-preston-reaction/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-preston-reaction/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-18T18:35:50Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f029e",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "Team News: Preston North End (H)",
            "type": [
                "Club News"
            ],
            "teaserText": "Liam Rosenior has made four changes to his starting XI as the Tigers welcome Preston North End to the MKM Stadium.",
            "content": "<p>Dimitrios Pelkas, Ozan Tufan, Xavier Simons and Greg Docherty all come into the side for Ryan Woods, Regan Slater, Benjamin Tetteh and Jean Micha&euml;l Seri.</p>\n<p>Jacob Greaves captains the side while Matt Ingram starts in goal.</p>\n<p><strong><em>Hull City (4-3-3): </em></strong><em>Ingram; Christie, A. Jones, McLoughlin, Greaves &copy;; Tufan, Docherty, Simons; Longman; Estupi&ntilde;&aacute;n, Pelkas. </em></p>\n<p><strong><em>Substitutes: </em></strong><em>Darlow, Coyle, Elder, Ebiowei, Woods, Slater, Tetteh. </em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/team-news-preston-north-end-h/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/team-news-preston-north-end-h/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-18T14:01:23Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a0",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Chorlton and Hinds join Bradford PA on loan",
            "type": [
                "Academy"
            ],
            "teaserText": "Under-21s duo Louie Chorlton and Josh Hinds have joined National League North side Bradford Park Avenue on loan until 18th March.",
            "content": "<p>Former Bradford City defender Chorlton, 20, signed a contract extension in January until the end of the season, having initially arrived last summer on a short-term deal following a successful trial.</p>\n<p>Forward Hinds, 19, made two senior appearances for the Tigers last season and has previously enjoyed loan spells with non-league sides Spalding United, Gainsborough Trinity and Boston United.</p>\n<p>Mark Bower&rsquo;s Green Army are 19<sup>th</sup> in the sixth tier, two points above the relegation zone with 15 matches remaining, and travel to promotion-chasing Brackley Town tomorrow.</p>\n<p>We wish both players the best of luck for their time at the Horsfall Stadium.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/siteassets/academy/2022-23-academy/josh-hinds-u21s-forest-jan-22-23.jpg\" alt=\"Josh Hinds U21s Forest Jan 22-23.jpg\" width=\"1359\" height=\"906\" loading=\"lazy\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/chorlton-and-hinds-join-bradford-pa-on-loan/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/chorlton-and-hinds-join-bradford-pa-on-loan/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T18:29:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f029f",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Ticket Details: Coventry City (a)",
            "type": [
                "Ticket News"
            ],
            "teaserText": "Tickets for our trip to Coventry City in the Sky Bet Championship on Saturday 11th March will go on sale to Members on Tuesday 21st February at 9.30am",
            "content": "<p>The game kicks off at 3pm at the Coventry Building Society Arena</p>\n<p>The Club have been given an initial allocation of 2,500.</p>\n<div>\n<div class=\"adition-container\">\n<div id=\"adition-instance-4558297\" class=\"adition-instance adition-outstream\" data-adition-excl=\"children\" data-adition-group=\"outstream\">During the Members' priority window, tickets are restricted to one ticket per Member.&nbsp;</div>\n<div class=\"adition-instance adition-outstream\" data-adition-excl=\"children\" data-adition-group=\"outstream\">&nbsp;</div>\n</div>\n</div>\n<p>Tickets will go on general sale at 9.30am on Friday 24th February</p>\n<p>&nbsp;</p>\n<p><strong>Ticket Prices:</strong></p>\n<p>&pound;25 - Adults</p>\n<p>&pound;20 - Senior (60+), Young Adults (18-21) and NUS/Armed Forces</p>\n<p>&pound;15 - Under 18</p>\n<p>&nbsp;</p>\n<p>Tickets will be available&nbsp;<a href=\"https://wearehullcity.talent-sport.co.uk/PagesPublic/ProductBrowse/productHome.aspx?ProductSubtype=AWAY\">online</a>, over the phone at 01482 505600, and from Tiger Leisure until 12 noon (9am online) on Friday 10th March.</p>\n<p>The last day to order your tickets for the post is Wednesday 8th March at 2pm.</p>\n<p>Tickets for collection must be collected from Tiger Leisure by no later than 12 noon on Friday 10th March. If you cannot collect from Tiger Leisure, please contact the Sales Centre before 12 noon on Friday 10th March on 01482 505600 to make alternative arrangements.</p>\n<p>Ambulant and wheelchair-disabled supporters are entitled to a free ticket for a personal assistant if required and are charged the price as seen above.</p>\n<p>Disabled tickets (ambulant and wheelchair) are only available to purchase over the phone at 01482 505600 or from Tiger Leisure at the MKM Stadium.</p>\n<p>Coventry City would like to make supporters aware that they maybe held back after the final whistle due to crowd control.</p>\n<p><strong>Tiger Travel:</strong></p>\n<p>Coach travel for the trip down to the West Midlands is<strong>&nbsp;FREE</strong>,&nbsp;<a href=\"https://www.wearehullcity.co.uk/news/2023/january/free-away-travel-for-remainder-of-202223-season/\">courtesy of our Chairman&nbsp;<strong>Acun Ilıcalı</strong></a>, until&nbsp;<strong>Friday 3rd March 2023 at 4pm.</strong>&nbsp;After this date, the below prices would apply.</p>\n<p>Members: &pound;22.50</p>\n<p>Non-Members: &pound;30</p>\n<p>Coaches will depart from the&nbsp;<strong>Spring Bank End, Walton Street</strong> at 10.30am on the day of the game.</p>\n<p>You can get Tiger Travel&nbsp;<a href=\"https://wearehullcity.talent-sport.co.uk/PagesPublic/ProductBrowse/ProductTravel.aspx\">online</a>, over the phone at 01482 505600, and in person at the Tiger Leisure store.</p>\n<p>To claim your free Tiger Travel online, you must have tickets for Coventry away in your basket or have purchased them on your ticketing account. The full price for Tiger Travel will show until you arrive at the basket page.</p>\n<p>Please note, once coach travel is booked, a supporter can then not change their designated coach (e.g Coach 1 to Coach 6). So we ask supporters to coordinate with their friends before purchasing if they wish to travel together on a coach.&nbsp;&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/ticket-details-coventry-city-a/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/ticket-details-coventry-city-a/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T19:30:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a1",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Estupiñán wins January Goal of the Month!",
            "type": [
                "Club News"
            ],
            "teaserText": "Óscar Estupiñán's strike against Huddersfield Town has seen him win Hull City Goal of the Month for January, presented by Uber Eats.",
            "content": "<p>The Colombian equalised for the Tigers after 98 minutes at the MKM Stadium&nbsp;as he remained alert to latch onto a header from Callum Elder in the last few seconds of the game before remaining composed to find the far corner.&nbsp;</p>\n<p>The 26-year-old's strike against the Terriers was good enough to see off competition from Jacob Greaves, Aaron Connolly and Under-18s forward Ralph Nkomba.</p>\n<p>Congratulations, &Oacute;scar!</p>\n<p><em>For more information about Uber Eats, visit&nbsp;<a href=\"https://www.ubereats.com/gb\">www.ubereats.com</a></em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/estupinan-wins-january-goal-of-the-month/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/estupinan-wins-january-goal-of-the-month/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T16:30:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a3",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "The Tiger - Issue 18",
            "type": [
                "Club News"
            ],
            "teaserText": "The next edition of 'The Tiger', our official matchday programme, is on sale ahead of our EFL Championship fixture against Preston North End.",
            "content": "<p>Notes from Chairman Acun Ilıcalı and Head Coach Liam Rosenior are included in this weekend's 64-page issue.&nbsp;</p>\n<p>We catch-up with Malcolm Ebiowei who made his debut for the Tigers just over a week ago against Stoke City, as well as Adama Traor&eacute; who continues his recovery from injury by featuring for our Under-21s.&nbsp;</p>\n<p>A Hall of Fame interview with Andy Dawson also features in this weekend's issue, as well as an interview with Under-18s defender Keegan Green on his recovery from injury.&nbsp;</p>\n<p>Make sure you get hold of your copy of 'The Tiger' at the MKM on Saturday, priced &pound;3, or buy online by clicking <a href=\"https://www.ignitionsportsmedia.com/collections/hull-city/products/hull-city-vs-preston-north-end-1\">here</a>.&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/the-tiger---issue-18/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/the-tiger---issue-18/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T14:20:33Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a2",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Join us early for Preston!",
            "type": [
                "Club News"
            ],
            "teaserText": "We host Preston North End at the MKM Stadium on Saturday afternoon, and we have plenty to keep those arriving early entertained ahead of kick-off!",
            "content": "<ul>\n<li>Hull City Sesh with Sam Connolly</li>\n<li>Opportunities to donate to the DEC Turkey-Syria Earthquake Appeal</li>\n<li>20% off Training Wear at Tiger Leisure</li>\n<li>Face Painters and Balloon Artists&nbsp;</li>\n<li>Food Bank Drop-Off at Tiger Leisure &amp; Tigers Trust Arena</li>\n<li>Table Football in the East Stand concourse</li>\n<li>Panna Cage outside West Stand Reception</li>\n<li>Fast Feet football drill in the South Stand Concourse</li>\n<li>FIFA Stations in the East, South and West Concourses&nbsp;</li>\n<li>Free Fruit</li>\n<li>Street Food - Including the Katsu Chicken Box (voted in by supporters!)</li>\n<li>Turnstiles Open at 1.30pm</li>\n</ul>\n<p>Sam Connolly, a singer-songwriter, makes his debut on the Hull City Sesh Stage in the Chris Chilton Stand concourse to play a live set from approximately 1.50pm to 2.30pm. To learn more about the artist, <a href=\"https://www.theseshhull.co.uk/artist/sam-connolly/\">tap here.</a></p>\n<p><a href=\"https://www.wearehullcity.co.uk/news/2022/october/coyle-pays-visit-to-hull-foodbank-with-your-donations/\">Supporters are encouraged to donate food and other essential items to provide Hull Foodbank with vital support, with drop-off points in place at Tiger Leisure and the Tigers Trust Arena.</a></p>\n<p>Fans are to be aware that the Club will hold a minute&rsquo;s silence before kick-off at tomorrow's fixture as a mark of respect to those who have tragically lost their lives due to the T&uuml;rkiye and Syria Earthquake. Donations can be made to the <a href=\"https://donate.redcross.org.uk/appeal/turkey-syria-earthquake-appeal\">DEC Turkey-Syria Earthquake Appeal here</a> and all profit from the <a href=\"https://www.tigerleisure.com/giftssouvenirs/souvenirs/flags/3580_small-turkish-flag-ac003.html?utm_source=Pre+Match&amp;utm_medium=Email&amp;utm_campaign=Turkey+Flag+Donation&amp;dm_t=0,0,0,0,0\">T&uuml;rkiye handhelds flags sold at Tiger Leisure</a> will go to the appeal.</p>\n<p>The Dugout Bar in the South Stand is open exclusively to Members for food and beverages from 12.30pm until 2.45pm and reopens after the game.&nbsp;</p>\n<p>If you still haven't got your ticket and want to see the Tigers take on Preston, then you can still purchase tickets <a href=\"https://wearehullcity.talent-sport.co.uk/PagesPublic/ProductBrowse/VisualSeatSelectionBoost.aspx?stadium=HT&amp;product=222317&amp;campaign=&amp;type=H&amp;productsubtype=HOME&amp;productIsHomeAsAway=N&amp;prodDesc=Hull+City+vs+Preston+North+End&amp;utm_source=Join+Us+Early&amp;utm_medium=Web&amp;utm_campaign=Preston+%28h%29+18.2.23\">online</a>, in-person from Tiger Leisure at the MKM Stadium, or by calling our Sales Centre on 01482 505600.&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/join-us-early-for-preston/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/join-us-early-for-preston/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T14:30:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a4",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Traoré: 'It felt good to be back'",
            "type": [
                "Interviews"
            ],
            "teaserText": "Making his first appearances in a Hull City shirt since his move from Hatayspor, Adama Traoré is hoping to make an impact in the first-team.",
            "content": "<p>Recovering from a long-term injury, the Mali international made his comeback when playing for the Under-21s in their fixtures against Barnsley in the Professional Development League.</p>\n<p>Playing 45 minutes in the first fixture before playing 60 minutes in the second game, the 27-year-old claimed an assist and scored his first goal in a Hull shirt in the most recent of fixtures, a 3-0 victory over the Tykes.</p>\n<p>&ldquo;It feels good to be back out on the pitch,&rdquo; said Traor&eacute;. &ldquo;It was a slow, but healthy recovery and I feel so happy to be able to play again.</p>\n<p>\"It was great to score and assist for the Under-21s (against Barnsley), but I cannot wait to help and score for the first team.&rdquo;</p>\n<p>Making the switch to the Tigers in September 2022, the attacking midfielder has enjoyed his time in Hull so far despite the injury.</p>\n<p>Speaking on his time at the club so far and the relationship he has with Head Coach Liam Rosenior, the 27-year-old first spoke about his targets for the remainder of the 2022/23 season as he steps up his recovery.</p>\n<p>&ldquo;I would just like to help the team push up the table,&rdquo; explained Traor&eacute;. &ldquo;I have been happy here at the football club and living in the city.</p>\n<p>\"The manager has been very good with me. He has given me good advice about my recovery and my importance to the team for when I am fully back.&rdquo;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/5a58f89fadd54fc19d5b95073fa6b3a1/image6lza.png\" alt=\"\" width=\"2700\" height=\"1800\" loading=\"lazy\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/traore-it-felt-good-to-be-back/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/traore-it-felt-good-to-be-back/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T12:30:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a5",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Christie wins January Player of the Month!",
            "type": [
                "Club News"
            ],
            "teaserText": "Cyrus Christie has been named Hull City Player of the Month for January.",
            "content": "<p>The 30-year-old defender was named in the starting line-up for all four of our Championship games against Wigan Athletic, Huddersfield Town, Sheffield United and Queens Park Rangers.&nbsp;</p>\n<p>Although the rampaging the right-back did not pick up any goals or assists in January, Christie still played a key role in the wins against Wigan and QPR and further established himself as one of the standout performers under the management of Liam Rosenior.&nbsp;</p>\n<p>Christie has now been named Player of the Month for October and January and has been nominated to win the award for four consecutive months. Having already scored our winning goal against Cardiff in February, the right-back will fancy his chances of being nominated again with two games left this month.&nbsp;</p>\n<p>Congratulations, Cyrus!&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/christie-wins-january-player-of-the-month/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/christie-wins-january-player-of-the-month/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T09:29:53Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a6",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "Dimitrios Pelkas Pre-Preston (H) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Dimitrios Pelkas spoke to the media ahead of Hull City's Championship match at home to Preston North End.",
            "content": "<p>The midfielder was asked about his return from injury against Norwich, how he is finding life under Liam Rosenior and the importance of returning to winning ways against Preston.</p>\n<p>Watch his pre-match press conference above!</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/dimitrios-pelkas-pre-preston-h-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/dimitrios-pelkas-pre-preston-h-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T09:12:56Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a7",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "Liam Rosenior Pre-Preston (H) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior spoke to the media ahead of Hull City's Championship match at home to Preston North End.",
            "content": "<p>The 38-year-old reflected on a dissapointing result in midweek against Norwich City, discussed the reaction shown by the players in training and previewed this weekend's game against the Lilywhites.&nbsp;</p>\n<p>Watch his pre-match press conference above!&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/liam-rosenior-pre-preston-h-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/liam-rosenior-pre-preston-h-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-17T08:53:10Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b7",
            "teamId": "t94",
            "optaMatchID": "g2300088",
            "title": "Stoke City 0-0 Hull City",
            "type": [
                "Match Reports"
            ],
            "teaserText": "Hull City made it three games unbeaten in the Sky Bet Championship after playing out a goalless draw against Stoke City.",
            "content": "<p>Both goalkeepers produced great saves throughout the fixture as the points were shared at the bet365 Stadium.</p>\n<p>Head Coach Liam Rosenior named an unchanged starting XI for a third consecutive fixture, with the only changes to the squad being on the bench as January signing Malcolm Ebiowei was named in the squad for the first time, while there was also a welcome return from injury for Dimitrios Pelkas.</p>\n<p>Prior to the fixture, both teams observed a minute&rsquo;s silence to pay respect to those affected by the recent tragedy in T&uuml;rkiye.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/81fa740230a04891942113fc46842ca9/hull-city-turkey-flag-v-stoke-city-a-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></p>\n<p>Ki-Jana Hoever had the first opportunity of the fixture, but his low shot was blocked well in the penalty area by City captain Jean Micha&euml;l Seri before Will Smallbone curled his left-footed strike around Matt Ingram&rsquo;s right-hand post.</p>\n<p>At the other end, &Oacute;scar Estupi&ntilde;&aacute;n forced the first save of the afternoon as his 20-yard effort was parried away by Matija &Scaron;arkić. From the resulting corner, taken by Callum Elder, the Montenegro international was on hand to deny Greg Docherty at the back post.</p>\n<p>Both goalkeepers were being kept busy as from a quick free-kick, Ingram was sharp down to his right-hand side to parry Tyrese Campbell&rsquo;s effort, before gathering the rebound from Josh Laurent, while Regan Slater stung the palms of &Scaron;arkić at his near post.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/81fa740230a04891942113fc46842ca9/matt-ingram-save-v-stoke-a-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1667\" loading=\"lazy\"></p>\n<p>Morgan Fox then blocked Ozan Tufan&rsquo;s strike before Sean McLoughlin fired over first-time from the resulting corner routine by Elder and Tufan. The Turkish midfielder tried his luck from range, hitting his free-kick from a tight angle towards the front post, but &Scaron;arkić punched the effort away.</p>\n<p>The Tigers continued to search for an opener in the first-half as Tufan&rsquo;s long ball forward was brought down brilliantly by Aaron Connolly. The Irish striker created a shooting opportunity on the edge of the &lsquo;D&rsquo;, but the curling shot went narrowly wide past the frame of the goal.</p>\n<p>Chances were a premium in the opening 25-minutes of the second-half, with Tufan&rsquo;s through ball for Connolly being pushed away by the onrushing &Scaron;arkić before Slater gathered the rebound on the left wing. The midfielder then cut the ball back for Connolly, but the forward couldn&rsquo;t keep his effort down as his first-time effort flew over the bar.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/81fa740230a04891942113fc46842ca9/aaron-connolly-v-stoke-city-a-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></p>\n<p>Both teams had decent crosses that were dealt with by the respected defences before Jacob Brown saw his volley deflect off Alfie Jones into the arms of Ingram.</p>\n<p>With six minutes left to play, the Tigers had a counter attack following a good pass by Ryan Woods. Fellow substitute Ryan Longman drove the ball forward before cutting inside on his right foot, but the winger could only drag his effort wide of the goal.</p>\n<p>Stoke substitute Emre Tezgel then delivered a low cross for Smallbone to flick on for Nick Powell, but the former Tigers&rsquo; flicked attempt was easily held by Ingram.</p>\n<p><strong>&nbsp;</strong></p>\n<p><strong><em>Stoke City (3-5-2): </em></strong><em>&Scaron;arkić; Wilmot, Jagielka &copy;, Fox; Hoever (Tezgel 75&rsquo;), Smallbone (Baker 90&rsquo;), Thompson, Laurent (Celina 90&rsquo;), Sterling; Campbell (Powell 60&rsquo;), Gayle (Brown 60&rsquo;). </em></p>\n<p><strong><em>Subs Not Used: </em></strong><em>Bonham, Okagbue. </em></p>\n<p><strong><em>&nbsp;</em></strong></p>\n<p><strong><em>Hull City (4-4-2): </em></strong><em>Ingram; Christie, A. Jones, McLoughlin, Elder (Greaves 63&rsquo;); Docherty, Tufan (Ebiowei 63&rsquo;), Seri &copy; (Woods 75&rsquo;), Slater (Coyle 90&rsquo;); Connolly (Longman 63&rsquo;), Estupi&ntilde;&aacute;n.</em></p>\n<p><strong><em>Subs Not Used: </em></strong><em>Darlow, Pelkas. </em></p>\n<p><em>&nbsp;</em></p>\n<p><strong><em>Attendance: </em></strong><em>23,970 (2,520 Hull City Supporters)</em></p>\n<p><em><img src=\"https://www.wearehullcity.co.uk/contentassets/81fa740230a04891942113fc46842ca9/hull-city-away-following-v-stoke-city-a-feb-23.jpg\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/stoke-city-0-0-hull-city/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/stoke-city-0-0-hull-city/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/5bd3e3fe-5f3e-4c36-b681-d3a390d1912a/Medium/g32-stoke-city-a-001.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/0abb5ed0-9a5e-4a45-a651-e335e1b9e743/Medium/g32-stoke-city-a-002.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/3fe86167-9c9d-45db-b402-0ce25926495c/Medium/g32-stoke-city-a-003.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/5fd9607c-edfa-412f-ad83-05f10b4f957f/Medium/g32-stoke-city-a-004.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/0c1c87eb-64f3-4721-90dc-45bf9257c519/Medium/g32-stoke-city-a-005.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/5bd5157b-808c-4b72-9bcf-fa9841235e49/Medium/g32-stoke-city-a-006.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/e01878ba-349e-4122-8ffd-ba86ef2cce7c/Medium/g32-stoke-city-a-007.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/81f31e95-1e04-4a2b-8c2d-6124a8dacee7/Medium/g32-stoke-city-a-008.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/36af4ac1-4b13-4615-a760-e2f2c81b1d40/Medium/g32-stoke-city-a-009.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/41590809-5d09-4351-ab86-a02152a91821/Medium/g32-stoke-city-a-010.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/5ecd1094-a0fc-40f5-9710-7a0a93b5dfad/Medium/g32-stoke-city-a-011.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/0c508229-280d-4ea3-b40f-de6dfa00c024/Medium/g32-stoke-city-a-012.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-11T16:56:13Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a8",
            "teamId": "t94",
            "optaMatchID": "g2300108",
            "title": "Connolly 'touch and go' for Preston",
            "type": [
                "Match Previews"
            ],
            "teaserText": "Liam Rosenior says he won’t know whether Aaron Connolly will feature against Preston North End until receiving a second opinion on the forward’s scan.",
            "content": "<p style=\"font-weight: 400;\">The 23-year-old was withdrawn with a foot injury after 64 minutes in the goalless draw against Stoke City and missed Tuesday&rsquo;s defeat against Norwich City at Carrow Road.</p>\n<p style=\"font-weight: 400;\">&ldquo;Aaron has had a scan,&rdquo; revealed Rosenior.</p>\n<p style=\"font-weight: 400;\">&ldquo;We&rsquo;re just waiting for a second opinion - the first opinion was that it could be something bigger than what it is. We&rsquo;re getting different reports that it may not be that.</p>\n<p style=\"font-weight: 400;\">&ldquo;I don&rsquo;t want to say something when it could be something else. At the moment, we&rsquo;re just waiting for the second opinion from the scan.</p>\n<p style=\"font-weight: 400;\">&ldquo;It was so innocuous what happened (against Stoke). When you aren&rsquo;t braced for impact, that&rsquo;s when you do damage. It&rsquo;s a really unfortunate one.</p>\n<p style=\"font-weight: 400;\">&ldquo;Aaron has been superb since he&rsquo;s come in. He&rsquo;s given us a different dimension. I think it was a dimension that we missed on Tuesday in terms of the way Norwich pressed and where the space was. The sooner he gets back, the better.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s not in a boot. He&rsquo;s walking around and is fine. It&rsquo;s not a serious one but we just need to see if he&rsquo;s available for Saturday. He&rsquo;s touch and go. Based off the scan, he could be training tomorrow, or he could not be.&rdquo;</p>\n<p style=\"font-weight: 400;\">Rosenior also confirmed that Cyrus Christie is available for selection ahead of this weekend&rsquo;s game after training with no problems on Thursday. &nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/connolly-touch-and-go-for-preston/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/connolly-touch-and-go-for-preston/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-16T16:18:32Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02ab",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Ticket Details: Burnley (h)",
            "type": [
                "Ticket News"
            ],
            "teaserText": "Tickets for Hull City's home Championship match against the current league leaders Burnley are now on sale.",
            "content": "<p>The Clarets, managed by <span class=\"yKMVIe\" role=\"heading\" aria-level=\"1\">Vincent Kompany</span>, will visit the MKM Stadium on Wednesday 15th March, with the game kicking off at 7.45pm.</p>\n<p>Ticket Prices:</p>\n<p><strong>Adult</strong></p>\n<p>Zone 1 &pound;20 | Zone 2 &pound;24 | Zone 3 &pound;28&nbsp;</p>\n<p><strong>Senior (65+)</strong></p>\n<p>Zone 1 &pound;13.50 | Zone 2 &pound;16 | Zone 3 &pound;18</p>\n<p><strong>Young Adults (16-22 Years)</strong></p>\n<p>Zone 1 &pound;10 | Zone 2 &pound;12 | Zone 3 &pound;14</p>\n<p><strong>Juniors (11-15 Years)</strong></p>\n<p>&pound;7 in all areas</p>\n<p><strong>Juniors (2-10 Years)</strong></p>\n<p>&pound;3 in all areas</p>\n<p>&nbsp;</p>\n<p>Tickets can be purchased&nbsp;<a href=\"https://wearehullcity.talent-sport.co.uk/PagesPublic/ProductBrowse/productHome.aspx?ProductSubType=HOME\">online</a>, in person from Tiger Leisure, or over the phone at 01482 505600.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/ticket-details-burnley-h/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/ticket-details-burnley-h/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-15T09:29:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02a9",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Vote for your January Goal of the Month!",
            "type": [
                "Club News"
            ],
            "teaserText": "Supporters can now vote for their Hull City Goal of the Month for January, presented by Uber Eats.",
            "content": "<p>We have selected four goals from four different players to make up the shortlist - three from the first team and one from the Under-18s.&nbsp;</p>\n<p>Our first contender is Jacob Greaves after he netted our first goal of the calendar year in the 4-1 win against Wigan Athletic - a composed first-time finish into the top corner with his trusty left foot.&nbsp;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/e38fea1a9e52410da0fb0c359d613998/fil-wigan-hull-090.jpg\" alt=\"\" width=\"2500\" height=\"1667\" loading=\"lazy\"></p>\n<p>&Oacute;scar Estupi&ntilde;&aacute;n netted twice for the Tigers in January, with his second goal of month in the 1-1 draw against Huddersfield Town making the shortlist. The 26-year-old was alert as Callum Elder headed on a lofted ball in the last few seconds of the game and remained calm to find the far corner.&nbsp;</p>\n<p>Estupi&ntilde;&aacute;n was not the only player to net twice for the Tigers in January, with Aaron Connolly also scoring his first goals for the club since joining on loan until the end of the season from Brighton &amp; Hove Albion. A good run and pass by Cyrus Christie down the right-hand side allowed the 23-year-old to see his deflected shot hit the back of the net to give us the lead in the 3-0 win against Queens Park Rangers.&nbsp;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/e38fea1a9e52410da0fb0c359d613998/fil-hull-qpr-0085.jpg\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></p>\n<p>Our final contender is Ralph Nkomba's finish for the Under-18s after their 4-1 home victory over Birmingham City. After winning possession high up the pitch, the young Tigers countered through Jake Brown, who squared the ball for Nkomba to confidently finish his one-on-one past the onrushing goalkeeper.&nbsp;</p>\n<p>Vote for your January Uber Eats Goal of the Month by clicking&nbsp;<a href=\"https://www.surveymonkey.co.uk/r/GoaloftheMonthJan\">here</a>.</p>\n<p><iframe src=\"https://www.youtube.com/embed/SLHYTUDU_lw\" frameborder=\"0\" width=\"560\" height=\"315\" allowfullscreen=\"allowfullscreen\" title=\"YouTube video player\" loading=\"lazy\"></iframe></p>\n<p><em>For more information about Uber Eats, visit&nbsp;<a href=\"http://www.ubereats.com/\">www.ubereats.com</a></em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/vote-for-your-january-goal-of-the-month/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/vote-for-your-january-goal-of-the-month/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-15T16:20:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02af",
            "teamId": "t94",
            "optaMatchID": "g2322023",
            "title": "Simms: ‘We deserved the win’",
            "type": [
                "Academy"
            ],
            "teaserText": "Jim Simms says the team deserved all three points in the Professional Development League as the Under-21s reigned 3-0 victors over Barnsley.",
            "content": "<p>Simms netted the Tigers&rsquo; third on the 33<sup>rd</sup> minute to secure Conor Sellars&rsquo; men the win at the Dransfield Stadium.</p>\n<p>&ldquo;We owed Barnsley after losing 2-1 against them last week and we made it right,&rdquo; said Simms.</p>\n<p>&ldquo;We fully deserved the win. From minute one, we knew we had to be on it by fighting for everything and we knew once we did, we would get our just rewards and earn the right to play.</p>\n<p>&ldquo;I was delighted to get back on the scoresheet. I have had a bit of a drought recently and I say to myself every game, &lsquo;goal or assist&rsquo;.</p>\n<p>&ldquo;I just want to keep getting my stats up and see how far I can go throughout the season, because I think I am in double figures now for both goals and assists.</p>\n<p>&ldquo;It was important to get the clean sheet and fair play to all the lads and &lsquo;Robbo&rsquo; (David Robson) who pulled off some great saves. The defenders and midfielders did their job and we did our job as forwards so it was a full team performance.&rdquo;</p>\n<p>There were milestones for two players in the fixture against the Tykes, with Adama Traor&eacute; netting his first goal in Hull City colours and Harry Vaughan making his official Under-21s debut since his deadline day move from Oldham Athletic.</p>\n<p>Speaking on the duo, Simms said first about Traor&eacute;: &ldquo;You can tell he is getting back into the flow of things after being out for so long and you can tell he has a bit of magic in him.</p>\n<p>&ldquo;You know he can change a game within an instant, as showed by his goal. He has contributed and we&rsquo;re all grateful for him to come down and play for us and do his bit.</p>\n<p>&ldquo;It feels surreal that Harry is here after our time together at Oldham Athletic and hopefully he can kick on now himself and get his stats going and help the team out.&rdquo;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/148be61eec0f4c7aa1a7b1982b695477/imagejtn6h.png\" alt=\"\" width=\"2700\" height=\"1800\" loading=\"lazy\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/simms-we-deserved-the-win/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/simms-we-deserved-the-win/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T14:35:21Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b1",
            "teamId": "t94",
            "optaMatchID": "g2322023",
            "title": "Hull City U21s 3-0 Barnsley U21s",
            "type": [
                "Academy"
            ],
            "teaserText": "Hull City Under-21s returned to winning ways as they defeated Barnsley 3-0 at the Dransfield Stadium.",
            "content": "<p>First-half finishes from Harry Wood, Adama Traor&eacute; and Jim Simms saw the Tigers move into the Professional Development League play-offs after earning all three points.</p>\n<p>Conor Sellars made three changes to the side that were defeated narrowly 2-1 by the Tykes in the reverse of this fixture last time out. Harry Vaughan came in to make his Under-21s debut in place of Jack Leckie, while Louie Chorlton and Alfie Taylor, replacing on-loan duo Tom Nixon and Jevon Mills respectively.</p>\n<p>Prior to the fixture, both teams observed a minute&rsquo;s silence to pay respect to those affected by the tragedy in T&uuml;rkiye.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/0852c6ba2eb348219dc5b791a765c2e4/minutes-silence-2-min.jpg\" alt=\"\" width=\"2700\" height=\"1800\" loading=\"lazy\"></p>\n<p>It took the Tigers just four minutes to open the scoring as Traor&eacute; teed up Wood to drill a first-time shot from 25-yards into the bottom corner.</p>\n<p>The provider turned scorer just six minutes later. The ball worked its way to the edge of the area and after shifting the ball onto his left foot, the Mali international placed his effort into the far corner to double the Tigers&rsquo; lead.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/0852c6ba2eb348219dc5b791a765c2e4/adama-traore-9-min.jpg\" alt=\"\" width=\"2700\" height=\"1800\" loading=\"lazy\"></p>\n<p>Will Jarvis searched for a third as his near-post attempt was held by Barnsley goalkeeper Adam Hayton, while Taylor blocked well to deny Charlie Williams at the other end.</p>\n<p>David Robson was then called into action to first deny Clarke Odour before acrobatically pushing over Harrison Nejman&rsquo;s dipping drive.</p>\n<p>On the 33<sup>rd</sup> minute, the Tigers had their third. A quick initial counter attack saw Jarvis play in Callum Jones, who played an out-of-the-foot cross towards Simms, who guided his header into the far corner.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/0852c6ba2eb348219dc5b791a765c2e4/jim-simms-goal-6-min.jpg\" alt=\"\" width=\"2700\" height=\"1800\" loading=\"lazy\"></p>\n<p>The visitors continued their search to get on the scoresheet as Robson saved against Nejman and Luke Thomas before Danny Benson blasted over from a corner delivery.</p>\n<p>A minute before half-time, Odour hit the crossbar from six-yards after a whipped delivery deflected into his path.</p>\n<p>Chances were a premium in the second half as Hayton held Jarvis&rsquo; drilled strike before the winger set up Jones to shoot, but the midfielder&rsquo;s first-time attempt deflected wide of the right-hand post.</p>\n<p>At the other end, Robson denied Odour before substitute Fareed Salifu somehow missed from close-range after crafting his chance inside the Tigers&rsquo; penalty area.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/0852c6ba2eb348219dc5b791a765c2e4/harry-vaughan-21-min.jpg\" alt=\"\" width=\"2700\" height=\"1800\" loading=\"lazy\"></p>\n<p>&nbsp;</p>\n<p><strong><em>Hull City (4-3-3): </em></strong><em>Robson; Chorlton, Jacob &copy;, Taylor, Trialist; Traor&eacute; (O. Green 60&rsquo;), Wood, C. Jones; Vaughan, Simms, Jarvis (Leckie 87&rsquo;). </em></p>\n<p><strong><em>Subs Not Used: </em></strong><em>Fisk, Snelgrove, Hall. </em></p>\n<p><em>&nbsp;</em></p>\n<p><strong><em>Barnsley (3-5-2): </em></strong><em>Hayton; Pache (Salifu 46&rsquo;), Helliwell, Benson; Otseh-Taiwo, Bramble (Hartley 87&rsquo;), Nejman &copy;, Thomas, Vaz; Williams (Butterfill 81&rsquo;), Odour. </em></p>\n<p><strong><em>Subs Not Used: </em></strong><em>Ravenhill, Hickingbottom. </em></p>\n<p><em>&nbsp;</em></p>\n<p><strong><em>Attendance: </em></strong><em>158</em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/hull-city-u21s-3-0-barnsley-u21s/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/hull-city-u21s-3-0-barnsley-u21s/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/e5f4f3d9-1323-4276-b8ba-08f01215da4a/Medium/hull-city-u21s-v-barnsley-h-feb-23-1-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/c13e54ed-0b16-4e19-9d31-b4d8401a88ff/Medium/hull-city-u21s-v-barnsley-h-feb-23-10-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/8db29578-709a-43d3-b585-941c99447991/Medium/hull-city-u21s-v-barnsley-h-feb-23-11-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/64e15c2e-5ff1-4a16-bb3a-e0362fa6e7a2/Medium/hull-city-u21s-v-barnsley-h-feb-23-2-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/2e89fea0-4d54-4b27-a5a8-388933e759ae/Medium/hull-city-u21s-v-barnsley-h-feb-23-3-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/d8b64971-f730-4cd0-be84-52f1ead66e18/Medium/hull-city-u21s-v-barnsley-h-feb-23-4-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/4a7c2035-b39c-41b3-93df-4ce5101e9a2d/Medium/hull-city-u21s-v-barnsley-h-feb-23-5-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/782b5c58-e5d9-4653-878d-b7f61c71ce19/Medium/hull-city-u21s-v-barnsley-h-feb-23-6-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/cd54d7c7-d265-4d07-94b4-b16ce79d01f2/Medium/hull-city-u21s-v-barnsley-h-feb-23-7-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/997bed94-0419-4ca9-9ac0-5cdb05dfca7f/Medium/hull-city-u21s-v-barnsley-h-feb-23-8-min.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/4b986922-dcf3-44ec-8a7b-28892e928f87/Medium/hull-city-u21s-v-barnsley-h-feb-23-9-min.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T12:00:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b2",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Snelgrove joins Spennymoor on loan",
            "type": [
                "Academy"
            ],
            "teaserText": "Under-21s forward McCauley Snelgrove has joined National League North side Spennymoor Town on a one-month loan deal.",
            "content": "<p>Born in Hull, Snelgrove first joined the club at Under-9s level and progressed through the academy system to pen his first professional contract with his boyhood team in March 2021.</p>\n<p>Last campaign, the second-year professional scored four goals in 11 appearances over two spells for Cleethorpes Town before joining NIFL Premiership outfit Crusaders FC at the beginning of the 2022/23 season.</p>\n<p>Netting twice and registering two assists, Snelgrove made 10 appearances for Crusaders before being recalled by City in January.</p>\n<p>More recently, he has featured for the Under-21s in their Premier League Cup Group A fixture against Nottingham Forest last month.</p>\n<p>Spennymoor, who are under the interim management of Jason Ainsley, are currently sat 12<sup>th</sup> in the National League North table on 41 points, six points from the play-offs.</p>\n<p>Snelgrove will be looking to make his debut on Saturday when they host Scarborough Athletic at The Brewery Field, kick-off: 3pm.</p>\n<p>We would like to wish McCauley all the best in his loan move to Spennymoor.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/snelgrove-joins-spennymoor-on-loan/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/snelgrove-joins-spennymoor-on-loan/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T10:28:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02ac",
            "teamId": "t94",
            "optaMatchID": "g2300095",
            "title": "Rosenior's Norwich Reaction",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior admitted his team were below par in the 3-1 defeat away to Norwich City.",
            "content": "<p>Finishes from Kieran Dowell, Gabriel Sara and Josh Sargent inflicted a second reverse in 11 games after Jacob Greaves&rsquo; fourth goal of the season had given the Tigers an early advantage.</p>\n<p>&ldquo;I&rsquo;m always unhappy when we lose,&rdquo; said Rosenioir.</p>\n<p>&ldquo;I wasn&rsquo;t celebrating the first goal because I felt we started the game shakily. They started the game on the front foot and really pressed us hard, which we expected.</p>\n<p>&ldquo;They&rsquo;ve got a very good manager and very good players at this level. I think they&rsquo;ll be the first to say they expect to be higher up the table than they are.&nbsp;</p>\n<p>&ldquo;That&rsquo;s no excuse for us. We lost too many battles, in terms of holding the ball up and up front, winning second balls &ndash; every single second ball seemed to land to them.</p>\n<p>&ldquo;They&rsquo;ve got a lot of quality and they scored some quality goals. We didn&rsquo;t really lay a glove on them apart from the first goal.</p>\n<p>&ldquo;We didn&rsquo;t play at the right level. They were the much better team. Our performance, probably for the first time since I&rsquo;ve been here, wasn&rsquo;t to the level I expected.</p>\n<p>&ldquo;What we&rsquo;ve got to do is react and make sure we go again on Saturday.&rdquo;</p>\n<p>Aaron Connolly missed the long trip to Carrow Road due to a foot injury, with the boss admitting he is also a doubt for the visit of Preston North End.</p>\n<p>&ldquo;He tried to train yesterday (Monday) and it wasn&rsquo;t right so we got him off for a scan,&rdquo; revealed Rosenior.</p>\n<p>&ldquo;He&rsquo;s had his scan and we&rsquo;re waiting for the results but he looks a doubt for Saturday.</p>\n<p>&ldquo;It&rsquo;s a combination of his toe and his ankle. I didn&rsquo;t want to bring him off against Stoke but I felt he was hobbling.</p>\n<p>&ldquo;Hopefully, he can get better soon rather than later because he&rsquo;s been a real bright spot for us.&rdquo;</p>\n<p>Rosenior said Cyrus Christie was withdrawn in the second half as a precaution due to a dead leg.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-norwich-reaction/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-norwich-reaction/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T23:33:52Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02ae",
            "teamId": "t94",
            "optaMatchID": "g2300095",
            "title": "Team News: Norwich City (A)",
            "type": [
                "Club News"
            ],
            "teaserText": "Liam Rosenior has freshened up his side by making four changes for tonight’s trip to Norwich City.",
            "content": "<p>Jacob Greaves replaces Callum Elder at left-back and takes the captain&rsquo;s armband, with Ryan Woods and Ryan Longman brought into midfield in place of Greg Docherty and Ozan Tufan.</p>\n<p>Benjamin Tetteh, back available after serving a three-match suspension for his red card at Sheffield United, is deployed up front as Aaron Connolly misses out through injury.</p>\n<p>&nbsp;</p>\n<p><strong>Hull City (4-4-2):</strong> Ingram; Christie, A Jones, McLoughlin, Greaves &copy;; Longman, Woods, Seri, Slater; Tetteh, <em>Estupi&ntilde;&aacute;n.</em></p>\n<p>Substitutes: Darlow, Coyle, Elder, Tufan, Docherty, Ebiowei, Pelkas.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/team-news-norwich-city-a/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/team-news-norwich-city-a/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-14T18:40:24Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b5",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Loan Watch: Mills makes Solihull debut",
            "type": [
                "Club News"
            ],
            "teaserText": "Hull City have 11 players out on loan. Here we take a look at how our Tigers are getting on in loan watch…",
            "content": "<p>&nbsp;</p>\n<p><strong>Jevon Mills at Solihull Moors:</strong></p>\n<p>We start with defender Jevon Mills who made his debut for Solihull Moors, playing the full game as Neal Ardley&rsquo;s men suffered a 2-1 defeat against Dagenham &amp; Redbridge. Despite Bartosz Cybulski scoring for Solihull, a brace from Inih Effiong saw Dagenham reign victorious.</p>\n<p>Mills, who was booked in the fixture, will be looking to feature tomorrow night as the Moors travel to face Boreham Wood before they welcome Chesterfield to Damson Park on Saturday.</p>\n<p><strong>Yuriel Celi at Club Universitario de Deportes: </strong></p>\n<p>Another debutant as in Peru, Yuriel Celi made his first appearance for Club Universitario de Deportes in a 1-0 defeat against Comercio in the Peruvian Primera Divisi&oacute;n Schedule Group Apertura.</p>\n<p>The midfielder replaced Rodrigo Ure&ntilde;a in the 81<sup>st</sup> minute as Miguel Carranza&rsquo;s finish in the 79<sup>th</sup> minute saw Comercio secure victory.</p>\n<p><strong>Brandon Fleming &amp; Tyler Smith at Oxford United: </strong></p>\n<p>In League One, Brandon Fleming and Tyler Smith both started for Oxford United as they drew 1-1 away at Milton Keynes Dons. Lewis Bate&rsquo;s equaliser cancelled out Sullay Kaikai&rsquo;s first-half finish to secure Oxford a point in League One.</p>\n<p>Fleming, who received a yellow card, played the full game while Smith played 61 minutes before being replaced by Tyler Goodrham. Oxford host league leaders Plymouth Argyle tomorrow night at the Kassam Stadium.</p>\n<p><strong>Tom Nixon &amp; Billy Chadwick at Boston United: </strong></p>\n<p>Fellow on-loan duo Tom Nixon and Billy Chadwick helped Boston United secure a vital three points in their fight against relegation from the National League North, defeating Darlington 3-1 at the Boston Community Stadium.</p>\n<p>Nixon, who was making his second Boston debut, played the full game while Chadwick played 70 minutes before being replaced by Lirak Hasani. The result saw Ian Culverhouse&rsquo;s men climb out of the relegation places and are currently sat 20<sup>th</sup> in the table.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/1211345a306e40c68d37817587893ae1/ben-voase-beverley-paul-adamson.jpg\" alt=\"\" width=\"1109\" height=\"740\" loading=\"lazy\"></p>\n<p><strong>Ben Voase at Beverley Town: </strong></p>\n<p>Second-year scholar Ben Voase made his seventh league appearance for Beverley Town as they hosted Northern Counties East Football League Division One leaders Campion.</p>\n<p>It was a thrilling encounter that finished in a 3-3 draw. Scott Phillips netted a brace for Beverley alongside a stunning acrobatic finish from Oliver Baldwin, but a 90<sup>th</sup>-minute goal from Campion&rsquo;s Nicky Boshell meant the points were shared at Norwood Park.</p>\n<p><em>(Ben Voase Beverley Town Image: Paul Adamson)</em></p>\n<p><strong>Andy Smith at Grimsby Town:</strong></p>\n<p>Banbury-born defender Andy Smith played the full 90 minutes for Grimsby Town in a 3-0 fourth round FA Cup replay victory over Championship outfit Luton Town.</p>\n<p>After setting up a mouth-watering fifth round tie against Premier League side Southampton, Smith was an unused substitute in the Mariners most recent fixture, a 1-0 home defeat against Colchester United.</p>\n<p><strong>Harvey Cartwright at Wycombe Wanderers:</strong></p>\n<p>Goalkeeper Harvey Cartwright was an unused substitute as Wycombe Wanderers defeated fellow League One play-off rivals Derby County 3-2 at Adams Park, ending the Rams&rsquo; 1-match unbeaten league run.</p>\n<p>The 20-year-old has been an unused substitute in Wycombe&rsquo;s last two league fixtures and will be hopeful to make his debut tomorrow night as the Chairboys travel to face Accrington Stanley.</p>\n<p><strong>Jake Leake at Scunthorpe United:</strong></p>\n<p>After making his Scunthorpe debut in their most recent fixture against Barnet, Jake Leake will be looking to make his second appearance for the Iron as they host Dagenham &amp; Redbridge on Saturday.</p>\n<p><strong>Doğukan Sinik at Antalyaspor: </strong></p>\n<p>In T&uuml;rkiye, S&uuml;per Lig fixtures were postponed following the tragic events in the country.</p>\n<p>Doğukan Sinik and Antalyaspor were scheduled to face Kasimpasa last Saturday, with the league suspended until the 3rd of March.&nbsp;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/1211345a306e40c68d37817587893ae1/image8jg5.png\" alt=\"\" width=\"2500\" height=\"1664\" loading=\"lazy\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/loan-watch-mills-makes-solihull-debut/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/loan-watch-mills-makes-solihull-debut/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-13T12:18:59Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b3",
            "teamId": "t94",
            "optaMatchID": "g2300095",
            "title": "Liam Rosenior Pre-Norwich (A) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior spoke to the media ahead of Hull City's Championship match away at Norwich City.",
            "content": "<p>The 38-year-old reflected on the goalless draw against Stoke City, praised the noise generated by the travelling faithful at the bet365 Stadium and previewed the trip to Carrow Road.</p>\n<p>Watch his pre-match press conference above!</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/liam-rosenior-pre-norwich-a-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/liam-rosenior-pre-norwich-a-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-13T13:55:30Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b4",
            "teamId": "t94",
            "optaMatchID": "g2300095",
            "title": "Connolly a doubt for Norwich",
            "type": [
                "Match Previews"
            ],
            "teaserText": "Forward Aaron Connolly is a fitness doubt for tomorrow’s game against Norwich City.",
            "content": "<p style=\"font-weight: 400;\">The 23-year-old was taken off with a knock after 64 minutes against Stoke City and will be assessed ahead of the trip to Carrow Road.</p>\n<p style=\"font-weight: 400;\">&ldquo;We&rsquo;re checking on him today (Monday),&rdquo; revealed Liam Rosenior.</p>\n<p style=\"font-weight: 400;\">&ldquo;We&rsquo;ll see how he is. We&rsquo;ll give him the maximum amount of time to recover to see if he&rsquo;s going to be fit for tomorrow. At the moment, I&rsquo;m not sure if he&rsquo;s going to be fit.&rdquo;</p>\n<p style=\"font-weight: 400;\">The Tigers will travel to Norfolk boosted by the return of Benjamin Tetteh who is available for selection again after serving his three-match suspension for a red card received in the defeat against Sheffield United in January.</p>\n<p style=\"font-weight: 400;\">&ldquo;Benjamin is back,&rdquo; said Rosenior.</p>\n<p style=\"font-weight: 400;\">&ldquo;Obviously, I&rsquo;ve got to be conscious that we were trying to build him back up to fitness with him being out for while. He came back in and got sent off. It&rsquo;s brilliant to have another fantastic option to work with.</p>\n<p style=\"font-weight: 400;\">&ldquo;Hopefully, we&rsquo;ll be able to build his levels up again and get him to where he needs to be.&rdquo;</p>\n<p style=\"font-weight: 400;\">With 63 minutes played at the bet365 Stadium, Malcolm Ebiowei came on to make his debut for the Tigers as he replaced Ozan Tufan. Dimitrios Pelkas also made a return to the squad as an unused substitute and with Allahyar Sayyadmanesh back in training, Rosenior is excited about the attacking options at his disposal.</p>\n<p style=\"font-weight: 400;\"><img src=\"https://www.wearehullcity.co.uk/contentassets/73042d49618f4a43aeeaf6f67409db17/imagego3hj.png\" alt=\"\" width=\"2500\" height=\"1667\" loading=\"lazy\"></p>\n<p style=\"font-weight: 400;\">&ldquo;This is what I was planning for and hoping for when I first joined,&rdquo; explained Rosenior.</p>\n<p style=\"font-weight: 400;\">&ldquo;I wanted to have a smaller squad, but a squad full of quality and I&rsquo;m starting to get that now.</p>\n<p style=\"font-weight: 400;\">&ldquo;Allahyar has come back and is in training. He&rsquo;s a little bit further away. These are the options I want.</p>\n<p style=\"font-weight: 400;\">&ldquo;It&rsquo;s great for me as it will allow me at times, especially throughout March and April, to freshen up the team. At the moment, we&rsquo;ve got to keep rolling and keep being consistent.&rdquo;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/connolly-a-doubt-for-norwich/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/connolly-a-doubt-for-norwich/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-13T12:42:42Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b6",
            "teamId": "t94",
            "optaMatchID": "g2300088",
            "title": "Rosenior's Stoke Reaction",
            "type": [
                "Interviews"
            ],
            "teaserText": "Liam Rosenior praised the mentality and spirit of his players after the gritty goalless stalemate at Stoke City.",
            "content": "<p>Aaron Connolly skied the best chance shortly after the restart having also gone close with a curling first-half effort, while &Oacute;scar Estupi&ntilde;&aacute;n&rsquo;s powerful long-range strike tested goalkeeper Matija &Scaron;arkić.</p>\n<p>The Tigers have now kept three consecutive clean sheets for the first time since January 2022 and lost only once in 10 league matches.</p>\n<p>&ldquo;It was a really hard-fought game,&rdquo; said Rosenior.</p>\n<p>&ldquo;Stoke gave us a lot of respect in the way they set-up; they tried to make it a scrap and a fight, and I thought Alex (Neil) set his team up really well in terms of doing that.</p>\n<p>&ldquo;They tried to take us out of our normal playing rhythm. They tried to press us from the front and stop us playing out.</p>\n<p>&ldquo;Once we got to grips with it after 10 minutes, I thought we had the better chances in the game. On another day, I think we end up winning the game.</p>\n<p>&ldquo;What I have to say &ndash; one, the fans were absolutely magnificent, and two, I think you&rsquo;re seeing a team.</p>\n<p>&ldquo;We&rsquo;re seeing a team that are very difficult to beat &ndash; that&rsquo;s three clean sheets in a row &ndash; and I couldn&rsquo;t be happier with their mentality and spirit.&rdquo;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-stoke-reaction/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/roseniors-stoke-reaction/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-11T21:38:25Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b8",
            "teamId": "t94",
            "optaMatchID": "g2300088",
            "title": "Team News: Stoke City (A)",
            "type": [
                "Club News"
            ],
            "teaserText": "Head Coach Liam Rosenior has named an unchanged starting XI for a third consecutive fixture as the Tigers face Stoke City at the bet365 Stadium.",
            "content": "<p>The only changes to the squad are on the bench as January signing Malcolm Ebiowei is named in the squad for the first time, replacing Xavier Simons.</p>\n<p>There is also a welcome return from injury for Dimitrios Pelkas, who is named amongst the substitutes in-place of Tobias Figueiredo.</p>\n<p><strong><em>Hull City (4-4-2): </em></strong><em>Ingram; Christie, A. Jones, McLoughlin, Elder; Docherty, Tufan, Seri &copy;, Slater; Connolly, Estupi&ntilde;&aacute;n.</em></p>\n<p><strong><em>Substitutes: </em></strong><em>Darlow, Coyle, Greaves, Ebiowei, Woods, Longman, Pelkas. </em></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/team-news-stoke-city-a/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/team-news-stoke-city-a/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-11T13:57:43Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02b9",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Mills and Nixon depart on loan",
            "type": [
                "Academy"
            ],
            "teaserText": "Under-21s duo Jevon Mills and Tom Nixon have left the club on loan, joining Solihull Moors and Boston United respectively.",
            "content": "<p>Mills, 19, joins Solihull for a 28-day loan while Nixon, 20, will stay with Boston until the end of the season.</p>\n<p>Making his senior Tigers debut in November 2021 in a 1-0 away win against Cardiff City, Mills has previously spent time out on loan at Falkirk and Gateshead, making six appearances for the latter earlier this season.</p>\n<p>Nixon returns to Boston after making five appearances for the Pilgrims already this campaign. Winning three and drawing the other two fixtures he featured in, the full-back&rsquo;s time at Boston was cut short due to injury.</p>\n<p>Solihull sit 11<sup>th</sup> in the National League table, three points off the play-off places, while Boston are two points adrift of safety in 22<sup>nd</sup> in National League North.</p>\n<p>The defensive duo will be looking to make their respected debuts on Saturday when Solihull travel to face Dagenham &amp; Redbridge and Boston host Darlington.</p>\n<p>We would like to wish both Jevon and Tom the best of luck for their loan spells.</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/63e64783df1648d08c2b54d590ccfe55/tom-nixon-wide.jpg\" alt=\"\" width=\"1920\" height=\"1080\" loading=\"lazy\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/mills-and-nixon-depart-on-loan/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/mills-and-nixon-depart-on-loan/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-10T16:58:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02ba",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Welcome to Hull City, Malcolm!",
            "type": [
                "Feature"
            ],
            "teaserText": "Malcolm Ebiowei joined the Tigers on loan until the end of the season from Crystal Palace in January!",
            "content": "<p>Here are the best pictures of the 19-year-old checking out his new surroundings at the MKM Stadium!</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/welcome-to-hull-city-malcolm/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/welcome-to-hull-city-malcolm/",
            "galletyUrls": [
                "https://www.wearehullcity.co.uk/api/image/feedassets/007052db-de74-453d-8c61-5d08b5b0d19d/Medium/malcom-ebiowei-001.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/fbdfdf74-a9ae-4a9a-9010-26511920475f/Medium/malcom-ebiowei-002.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/259541e4-887b-4384-a99d-4049b534a65a/Medium/malcom-ebiowei-003.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/9e83292f-f3c8-42a4-b939-7dba3e146fd6/Medium/malcom-ebiowei-004.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/b048ec5c-98e9-42a7-84e4-d38eb49026a4/Medium/malcom-ebiowei-005.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/2f20e9e7-9715-49ed-b614-c8f0d3623ee8/Medium/malcom-ebiowei-006.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/23b176b1-cdc3-4411-8b2b-cda6a13ce59e/Medium/malcom-ebiowei-007.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/acefc55f-1a70-4acf-8d51-3486e4be2c8c/Medium/malcom-ebiowei-008.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/51a454ee-cae4-47a3-8f0a-2ca92fd74178/Medium/malcom-ebiowei-009.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/2490e077-759a-4a41-a676-318a7682b27f/Medium/malcom-ebiowei-010.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/27e4d8a5-e342-4b18-bc3d-e2b17a8056dc/Medium/malcom-ebiowei-011.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/b5c8a61b-794f-44ec-87ed-d0d3eed68947/Medium/malcom-ebiowei-012.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/bdb72613-fe2d-4785-878d-cca5d65803b1/Medium/malcom-ebiowei-013.jpg,https://www.wearehullcity.co.uk/api/image/feedassets/9e6b71d9-a788-459d-8cba-726e13485d6e/Medium/malcom-ebiowei-014.jpg,"
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-10T16:00:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02bb",
            "teamId": "t94",
            "optaMatchID": "g2300088",
            "title": "Ozan Tufan Pre-Stoke City (A) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Turkish midfielder Ozan Tufan spoke candidly to the media via an interpreter about the tragic events in his homeland.",
            "content": "<p>The 27-year-old sent his heartfelt wishes to those affected and urged people to contribute to the humanitarian recovery efforts.</p>\n<table class=\"scroll-table\" style=\"height: 18px; width: 100%; border-collapse: collapse; border-style: hidden; margin-left: auto; margin-right: auto;\" border=\"0\">\n<tbody>\n<tr style=\"height: 18px;\">\n<td style=\"width: 42.7889%; background-color: #ffffff; height: 18px;\">&nbsp;</td>\n<td style=\"width: 17.8362%; border-style: solid; border-color: #e77400; background-color: #e77400; color: #ffffff; vertical-align: middle; height: 18px; text-align: center;\"><a href=\"https://donate.redcross.org.uk/appeal/turkey-syria-earthquake-appeal\">DONATE TO BRITISH RED CROSS</a></td>\n<td style=\"width: 39.3748%; background-color: #ffffff; height: 18px;\">&nbsp;</td>\n</tr>\n</tbody>\n</table>\n<p>Watch his press conference above.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/ozan-tufan-pre-stoke-city-a-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/ozan-tufan-pre-stoke-city-a-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-10T13:34:34Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02bc",
            "teamId": "t94",
            "optaMatchID": "g2300088",
            "title": "Justin Walker Pre-Stoke City (A) Press Conference",
            "type": [
                "Interviews"
            ],
            "teaserText": "Assistant head coach Justin Walker spoke to the media ahead of Hull City's Championship match away to Stoke City.",
            "content": "<p>The 47-year-old reflected on back-to-back home wins, discussed the improved form since his and Liam Rosenior's arrival and previewed the trip to the Bet365 Stadium.</p>\n<p>Watch his pre-match press conference above!</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/justin-walker-pre-stoke-city-a-press-conference/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/justin-walker-pre-stoke-city-a-press-conference/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-10T13:24:52Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02bd",
            "teamId": "t94",
            "optaMatchID": "g2300088",
            "title": "Ebiowei in contention for Stoke",
            "type": [
                "Match Previews"
            ],
            "teaserText": "Justin Walker says Malcolm Ebiowei could be in line to make his debut for the Tigers in this weekend’s game against Stoke City.",
            "content": "<p style=\"font-weight: 400;\">The 19-year-old linked up with the rest of Liam Rosenior&rsquo;s players earlier this week following his arrival from Premier League outfit Crystal Palace in January and will be hoping to feature at the bet365 Stadium.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s got a chance for this one,&rdquo; responded Walker when asked whether the teenager could feature against the Potters.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s trained for the last couple of days - he&rsquo;s a great guy. We&rsquo;ve worked with him before, and he was one of the first targets we spoke about when we came to the club.</p>\n<p style=\"font-weight: 400;\">&ldquo;He was someone we had in the back of our mind. We felt he could add something to this group. It&rsquo;s a great place for him to come and play and he already knows how we play.&rdquo;</p>\n<p style=\"font-weight: 400;\">Allahyar Sayyadmanesh has been sidelined with a hamstring injury since being taken off at half-time in the 4-1 win against Wigan Athletic and Walker has admitted this weekend&rsquo;s game might come too soon for him.</p>\n<p style=\"font-weight: 400;\">&ldquo;I think you saw what an impact he made when he came back,&rdquo; said Walker.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s an infectious character. I love seeing him as he brightens your day up. When he&rsquo;s playing, he always gives you everything. You&rsquo;ve just got to be mindful because he doesn&rsquo;t do anything in the half measures.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s that excited to get out there and play to his maximum for 90 minutes. We&rsquo;ve just got to be mindful and make sure this time we&rsquo;ve got a group which is competitive and we&rsquo;re not putting him at any risk whatsoever.&rdquo;</p>\n<p style=\"font-weight: 400;\">Adama Traor&eacute; played 45 minutes for the Under-21s on Tuesday as they fell to a 2-1 defeat against Barnsley in the Professional Development League and Walker is delighted with the progress the midfielder has made over recent weeks.</p>\n<p style=\"font-weight: 400;\">&ldquo;It was a shame he didn&rsquo;t score,&rdquo; added Walker.</p>\n<p style=\"font-weight: 400;\">&ldquo;That would&rsquo;ve gone down well with him. Having spent so much time out injured, it&rsquo;s good for him to be competitive again and to get a feel of being out with his teammates on the grass.</p>\n<p style=\"font-weight: 400;\">&ldquo;He&rsquo;s been training for a few weeks now and he&rsquo;s getting better each day. It&rsquo;s been a positive couple of weeks for him and hopefully, he&rsquo;ll be in and around the squad soon.&rdquo;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/ebiowei-in-contention-for-stoke/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/ebiowei-in-contention-for-stoke/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-10T10:25:36Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02be",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "No pay on the day at Stoke",
            "type": [
                "Ticket News"
            ],
            "teaserText": "Tickets for Hull City’s Sky Bet Championship fixture away at Stoke City will not be made available at the bet365 Stadium - meaning supporters must buy in advance.",
            "content": "<p>The game kicks off at 3pm at the bet365 Stadium on Saturday 11th February.</p>\n<p><strong>Ticket Prices:</strong></p>\n<p>&pound;25 - Adults</p>\n<p>&pound;19 - Senior (Over 65), Disabled Adults&nbsp;</p>\n<p>&pound;15 - Under 18s</p>\n<p>&pound;12 - Under 11s&nbsp;</p>\n<p>&nbsp;</p>\n<p>Tickets are available through the club&nbsp;<a href=\"https://wearehullcity.talent-sport.co.uk/PagesPublic/ProductBrowse/StandAndAreaSelection.aspx?stadium=ST&amp;product=HCST23&amp;campaign=&amp;type=H&amp;productsubtype=AWAY&amp;productIsHomeAsAway=N&amp;prodDesc=Stoke%20City%20vs.%20Hull%20City\">online</a>, over the phone on 01482 505600, and from Tiger Leisure until 12 noon (9am online) on Friday 10th February.</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/no-pay-on-the-day-at-stoke/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/no-pay-on-the-day-at-stoke/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-09T09:14:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02bf",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Panesar participates in PFA Asian inclusion mentoring scheme",
            "type": [
                "Academy"
            ],
            "teaserText": "Midfielder Aman Panesar was invited to participate in the PFA’s AIMS campaign at St George’s Park.",
            "content": "<p>Held on Sunday 29<sup>th</sup> January, this was the AIMS second annual &lsquo;Player Development Day&rsquo;, with over 100 players in attendance from 56 different clubs, which involved numerous different activities throughout the day.</p>\n<p>Led by Player Inclusion Executive at the PFA Riz Rehman, the pioneering work has gained significant momentum in developing a sustainable future for south Asian players in elite football, with the numbers of players within the Professional Development Phase being at an all-time high at any one given time.</p>\n<p>Being a ground-breaking period in AIMS third year of a five-year strategy, one in two professional club academies has at least one player with south Asian heritage signed on their books.</p>\n<p>Panesar was one of 28 scholars within the Professional Development Phase that attended, the highest number at any one given time. Speaking on the event, the Bradford-born midfielder said: &ldquo;It was really good and was really well done. It was really good to see a lot of other Asian people and it was good to act as an inspiration to the younger kids.</p>\n<p>&ldquo;It was a bit of an eye opener for the younger children and myself. We were working with professionals like Zidane Iqbal (Manchester United); players who were at the top level.</p>\n<p>&ldquo;The numbers are growing. I might see one or two Asian players every couple of games now and that&rsquo;s a lot more than it used to be. You can also see all the younger academy players representing some of the big clubs as well so it&rsquo;s good to see.&rdquo;</p>\n<p>Riz Rehman added: &ldquo;I am extremely proud of the ground-breaking work we have achieved over the last two and a half years &ndash; the AIMS network has grown considerably and it&rsquo;s important to focus on the progress that has been made.</p>\n<p>&ldquo;From 30 players at our first event in 2021 to now having over 100 players at St George's Park illustrates the success of the work and the buy-in we&rsquo;ve had from players and families.</p>\n<p>&ldquo;Days like these are not possible without the support from clubs, who excused their players from competitive academy fixtures and encouraged their attendance at the event. I was delighted to see so many staff from clubs in attendance who recognise the additional value the AIMS programme brings to their players' overall holistic development.&rdquo;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/5a7c7326af2e4f869557b1714121bb49/imagelgwqm.png\" alt=\"\" width=\"4970\" height=\"3313\"></p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/panesar-participates-in-pfa-asian-inclusion-mentoring-scheme/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/panesar-participates-in-pfa-asian-inclusion-mentoring-scheme/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-08T16:58:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        },
        {
            "id": "63fc20de5acc731fb99f02c0",
            "teamId": "t94",
            "optaMatchID": "",
            "title": "Vote for your January Player of the Month!",
            "type": [
                "Club News"
            ],
            "teaserText": "Supporters can now vote for their Hull City Player of the Month for January.",
            "content": "<p>Defenders Sean McLoughlin and Cyrus Christie, midfielder Jean Micha&euml;l Seri and striker &Oacute;scar Estupi&ntilde;&aacute;n make up the four-man shortlist.&nbsp;</p>\n<p>December's Player of the Month McLoughlin has made the four-man shortlist for a second consecutive month as his impressive performances continued in January. The 26-year-old centre-back was ever-present at the back in a defence that only conceded three goals in four league games.&nbsp;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/0e80c1eb2c9f408abb3e3b9ef6ebe629/imagew4ysl.png\" alt=\"\" width=\"2500\" height=\"1667\"></p>\n<p>Another standout performer was right-back Christie, who has now been nominated for a fourth consecutive month. Although the 30-year-old did not pick up any goals or assists in January, Christie still played a key role in the wins against Wigan and QPR.&nbsp;</p>\n<p>Midfielder Seri may not have featured against his former club Fulham in the FA Cup, but he remained ever-present in the Championship throughout January. The 31-year-old's performances continued to catch the eye and he picked up another assist for Aaron Connolly's second goal in the win against QPR. With 64 minutes played at the MKM, Seri's long ball released Connolly through on goal, who dispathced to secure himself a brace.&nbsp;</p>\n<p><img src=\"https://www.wearehullcity.co.uk/contentassets/0e80c1eb2c9f408abb3e3b9ef6ebe629/fil-wigan-hull-116.jpg\" alt=\"\" width=\"2500\" height=\"1664\"></p>\n<p>Our final nomination for Player of the Month is Estupi&ntilde;&aacute;n. The 26-year-old Colombian has netted 12 Championship goals, with two more coming in January against Wigan and Huddersfield. As it stands, only Middlesbrough's Chuba Akpom has scored more in the Championship with 16.&nbsp;</p>\n<p>Vote for your January Player of the Month by clicking <a href=\"https://www.surveymonkey.co.uk/r/POTMJanuary23\">here</a>.&nbsp;</p>",
            "articleURL": "https://www.wearehullcity.co.uk/news/2023/february/vote-for-your-january-player-of-the-month/",
            "imageUrl": "https://www.wearehullcity.co.uk/news/2023/february/vote-for-your-january-player-of-the-month/",
            "galletyUrls": [
                ""
            ],
            "videoUrl": "",
            "newsArticleID": "",
            "publishDate": "2023-02-08T14:00:00Z",
            "lastUpdateDate": "0001-01-01T00:00:00Z",
            "published": false
        }
    ],
    "metadata": {
        "createdAt": "2023-02-27T03:17:53.141814635Z",
        "totalItems": 50,
        "sort": "-published"
    }
}

## Get Article by ID

### Request where t94 is param {team} and 63fbfb4d26d4bc535d953e2d is param {id} this id will change when when try it because is a unique mongo hex id. So use a new one generated

curl --location --request GET 'localhost:4007/teams/t94/news/63fc11a3135b323a684f0742'

### Response

{"status":"succes","data":{"id":"63fc11a3135b323a684f0742","teamId":"t94","articleURL":"https://www.wearehullcity.co.uk/news/2023/february/roseniors-bristol-city-reaction/","newsArticleID":"443849","publishDate":"2023-02-25T21:23:37Z","type":["Interviews"],"teaserText":"Liam Rosenior felt a sluggish start set the tone for Hull City’s narrow 1-0 loss at Bristol City.","thumbnailImageURL":"https://www.wearehullcity.co.uk/api/image/feedassets/4739df60-379a-4fdd-bd92-23de9026785c/Medium/liam-rosenior-bristol-city-a-feb-2023.jpg","title":"Rosenior’s Bristol City Reaction","optaMatchID":"g2300117","lastUpdateDate":"2023-02-25T21:33:19Z","published":true},"metadata":{"createdAt":"2023-02-27T02:59:57.136305345Z"}

## Screenshots

![diagram](diagram.png)
