import { Hono } from "hono";
import { HTTPException } from "hono/http-exception";
import { Bindings } from "./binding";
import { api } from "./api";
import { Variables } from "./variables";


const PREFIX = "Bearer";

const app = new Hono<{ Bindings: Bindings }>();
app.get("/", (c) => c.text("Hello Hono!"));
app.post("/newuser", async (c) => {
  const UUID = crypto.randomUUID();
  const data = await c.req.json();
  await c.env.rnapi.put(UUID, data.fcm_token);
  return c.json({ uuid: UUID });
});
const middleware = new Hono<{ Bindings: Bindings; Variables: Variables }>();
middleware.use("/call", async (c, next) => {
  const headerToken = c.req.header("Authorization");
  if (!headerToken) {
    // No Authorization header
    const res = new Response(
      JSON.stringify({ Unauthorized: "No Authorization header" }),
      {
        status: 401,
      }
    );
    throw new HTTPException(401, { res });
  } else {
    const token = headerToken.split(" ", 2);
    if (token.length != 2 || token[0] != PREFIX) {
      const res = new Response(
        JSON.stringify({ "Bad Request": "Bad Authorization header" }),
        {
          status: 400,
          headers: {"Content-Type": "application/json"}
        }
      );
      throw new HTTPException(400, { res });
    } else {
      const fcmkey = await c.env.rnapi.get(token[1]);
      if (fcmkey == null) {
        const res = new Response(
          JSON.stringify({ Unauthorized: "Invalid Bearer Token" }),
          {
            status: 401,
            headers: {"Content-Type": "application/json"}
          }
        );
        throw new HTTPException(401, { res });
      } else {
        c.set("fcmkey", fcmkey);
      }
    }
  }
  await next();
});
app.route("/api", middleware);
app.route("/api", api);

export default app;
