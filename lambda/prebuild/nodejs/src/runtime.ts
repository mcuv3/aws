import { credentials } from "@grpc/grpc-js";
import { LambdaServiceClient } from "./gen/lambda_grpc_pb";
import { EventResponse, ReceiveEventRequest } from "./gen/lambda_pb";

function getNewClient(code: Function) {
  return (resolve: (v: string) => void, reject: (v: string) => void) => {
    const client = new LambdaServiceClient(
      process.env.HOST!,
      credentials.createInsecure()
    );
    const req = new ReceiveEventRequest();
    req.setHash(process.env.HASH!);

    const res = client.receiveEvents(req);

    res.on("error", (err) => {
      console.log("Something went wrong ", err);
      reject(err.message);
    });

    res.on("data", (data: EventResponse) => {
      code(data);
    });

    res.on("end", () => console.log("Ended"));

    res.on("close", () => {
      // resolve("Closed");
      console.log("Connection closed");
    });
  };
}

(async () => {
  // TODO: implement grpc client to receive messages from host
  const eventData = process.env.EVENT_DATA;
  const handler = process.env.HANDLER;

  const path = require("path");
  const ext = path.extname(handler);
  const fn = ext.replace(".", "");

  if (!handler) {
    throw new Error("Handler not set");
  }
  const pt = handler.replace(ext, "");

  const mod = require(pt);

  const code = mod[fn];

  if (typeof code !== "function") {
    console.log(`${fn} is not a function sorry`);
  }

  code(eventData);

  if (process.env.LISTEN) {
    await new Promise(getNewClient(code));
  }
})();
