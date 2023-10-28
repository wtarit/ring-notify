import { Hono } from "hono";
import { Bindings } from "./binding";
import { Variables } from "./variables";
import GoogleAuth, { GoogleKey } from "cloudflare-workers-and-google-oauth";

const api = new Hono<{ Bindings: Bindings; Variables: Variables }>();

api.post("/call", async (c) => {
  const data = await c.req.json();
  const scopes: string[] = [
    "https://www.googleapis.com/auth/firebase.messaging",
  ];

  const googleAuth: GoogleKey = JSON.parse(c.env.GCP_SERVICE_ACCOUNT);
  // Initialize the service

  const oauth = new GoogleAuth(googleAuth, scopes);

  const token = await oauth.getGoogleAuthToken();
  const res = await fetch(
    "https://fcm.googleapis.com/v1/projects/ring-notify/messages:send",
    {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        message: {
          token: c.get("fcmkey"),
          data: data,
        },
      }),
    }
  );
  // console.log(await c.env.rnapi.get("test"));
  return c.json({ status: "success" });
});

export { api };
