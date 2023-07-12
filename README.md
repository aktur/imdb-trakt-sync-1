[![sync](https://github.com/aktur/imdb-trakt-sync-1/actions/workflows/sync.yml/badge.svg?event=schedule)](https://github.com/aktur/imdb-trakt-sync-1/actions/workflows/sync.yml)
# imdb-trakt-sync
GoLang app that can sync [IMDb](https://www.imdb.com/) and [Trakt](https://trakt.tv/dashboard) user data - watchlist, 
lists, ratings and history.  
To achieve its goals the application is using the [Trakt API](https://trakt.docs.apiary.io/) and web scraping the IMDb website.  
Keep in mind that this application is performing a one-way sync from IMDb to Trakt.  
There are 3 possible modes to run this application and more details can be found in the [.env.example](.env.example) file.  
As much as I wanted to provide a two-way sync functionality, this will not be possible until IMDb decides to expose a public API.

## Sync your local fork with upstream

- git remote add upstream https://github.com/cecobask/imdb-trakt-sync.git
- git fetch upstream main\n
- git rebase upstream/main
- *resolve conflicts is any*
- git rebase --continue *(in case there were conflicts)*
- git push origin main --force

# Usage
The application can be setup to run automatically, based on a custom schedule (_default: once every 3 hours_) using 
`GitHub Actions` or locally on your machine. Follow the relevant section below, based on how you want to use the application. 

## Run the application using GitHub Actions
1. [Fork the repository](https://github.com/aktur/imdb-trakt-sync-1/fork) to your account
2. [Create a Trakt API application](https://trakt.tv/oauth/applications). Give it a name and use `urn:ietf:wg:oauth:2.0:oob` 
as redirect uri. The rest of the fields can be left empty
3. Configure GitHub repository secrets:
   - All the secrets (a.k.a. environment variables) are listed in the [.env.example](.env.example) file
   - Open the repository secrets dashboard of your fork
   - Create repository secrets: `Settings` > `Secrets` > `Actions` > `New repository secret`
   - Repeat the previous step for each secret individually
   - The repository secrets dashboard should look similar to this:![repository_secrets.png](assets/repository_secrets.png)
4. Enable GitHub Actions for the fork repository
5. Enable the `sync` workflow, as scheduled workflows are disabled by default in fork repositories
6. The `sync` workflow can be triggered manually right away to test if it works. Alternatively, wait for GitHub actions 
to automatically trigger it every 3 hours

### Refresh TRAKT_ACCESS_TOKEN

Every 3 minths a new trakt token (TRAKT_ACCESS_TOKEN secret in GitHub) must be generated. Go to https://trakt.tv/oauth/applications/103585 and click authenticate. It will redirect it
to page containig code. Copy this code and execute:

```bash
export TRAKT_CODE=3E52968B
curl --include \
     --request POST \
     --header "Content-Type: application/json" \
     --data-binary "{
    \"code\": \"$TRAKT_CODE\",
    \"client_id\": \"b73ecf5f8fc994d35b59f2ca925df6c6521bccf9b922de5ac96f507ac75a9384\",
    \"client_secret\": \"355a2f1d353049451ab2267ff337a83b54a2a002be32b1af9a6684294df29539\",
    \"redirect_uri\": \"urn:ietf:wg:oauth:2.0:oob\",
    \"grant_type\": \"authorization_code\"
}" \
'https://api.trakt.tv/oauth/token'
```
Put "access_token" in the corresponding secret on https://github.com/aktur/imdb-trakt-sync-1/settings/secrets/actions
## Run the application locally
1. Clone the repository to your machine
2. [Create a Trakt API application](https://trakt.tv/oauth/applications). Give it a name and use `urn:ietf:wg:oauth:2.0:oob`
   as redirect uri. The rest of the fields can be left empty.
3. Make a copy of the [.env.example](.env.example) file and name it `.env`
4. Populate all the environment variables in that file using the existing values as reference
5. Make sure you have GoLang installed on your machine. If you do not have it, [this is how you can install it](https://go.dev/doc/install).
6. Open a terminal window in the repository folder and run the application using the command `go run cmd/syncer/main.go`