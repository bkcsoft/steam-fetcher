# Steam Fetcher

Small HTTP Service that searches for Steam-apps and returns their app-id and name (since SteamAPI doesn't have a search function)

## Endpoints

- `/search?name={query}`

 Searches for an App

 #### Params

 `name - string` What do you think it does?!

 #### Response
```json
{
	"apps": [
		{
			"appid": integer,
			"name": string
		}
	]
}
```
