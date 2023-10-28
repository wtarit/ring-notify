# Run using Cloudflare Wrangler
Before you run you will need to add content of your service-account.json file into `.dev.vars` since it is used for firebase cloud messaging. https://developers.cloudflare.com/workers/configuration/secrets/

Since google auth library doesn't work on cloudflare worker, I used [cloudflare-workers-and-google-oauth](https://ryan-schachte.com/blog/cf-workers-auth) to handle creating short lived short-lived OAuth 2.0 access token used to call Firebase Cloud Messaging API https://firebase.google.com/docs/cloud-messaging/migrate-v1 
```
wrangler dev src/index.ts
```

```
npm install
npm run dev
```

```
npm run deploy
```
