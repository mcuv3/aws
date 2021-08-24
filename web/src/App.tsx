import { grpc } from "@improbable-eng/grpc-web";
import React, { useEffect } from "react";
import "./App.css";
import { RootUserLoginRequest } from "./gen/iam_pb";
import { IAMService } from "./gen/iam_pb_service";

function App() {
  const fetchServer = () => {
    const req = new RootUserLoginRequest();
    req.setEmail("mcuve@outlook.com");
    req.setPassword("maotrix1");
    grpc.unary(IAMService.RootUserLogin, {
      request: req,
      host: "http://localhost:7000",
      debug:true,
      onEnd: (res) => {
        console.log(res.message?.toObject());
      }, 
    });
  };

  useEffect(() => {
    fetchServer();

  }, [])

  // fetchServer();

  return (
    <div className="App">
      <h1>Hello</h1>
    </div>
  );
}

export default App;
