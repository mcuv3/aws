// package: lambda
// file: lambda.proto

var lambda_pb = require("./lambda_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var LambdaService = (function () {
  function LambdaService() {}
  LambdaService.serviceName = "lambda.LambdaService";
  return LambdaService;
}());

LambdaService.CreateFunction = {
  methodName: "CreateFunction",
  service: LambdaService,
  requestStream: false,
  responseStream: false,
  requestType: lambda_pb.CreateFunctionRequest,
  responseType: lambda_pb.LambdaResponse
};

LambdaService.TestFunction = {
  methodName: "TestFunction",
  service: LambdaService,
  requestStream: false,
  responseStream: false,
  requestType: lambda_pb.TestFunctionResquest,
  responseType: lambda_pb.LambdaResponse
};

LambdaService.InvoqueFunction = {
  methodName: "InvoqueFunction",
  service: LambdaService,
  requestStream: false,
  responseStream: false,
  requestType: lambda_pb.InvoqueFunctionRequest,
  responseType: lambda_pb.LambdaResponse
};

LambdaService.SeedLambdaServer = {
  methodName: "SeedLambdaServer",
  service: LambdaService,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: lambda_pb.LambdaResponse
};

LambdaService.ReceiveEvents = {
  methodName: "ReceiveEvents",
  service: LambdaService,
  requestStream: false,
  responseStream: true,
  requestType: lambda_pb.ReceiveEventRequest,
  responseType: lambda_pb.EventResponse
};

LambdaService.UpdateLambda = {
  methodName: "UpdateLambda",
  service: LambdaService,
  requestStream: false,
  responseStream: false,
  requestType: lambda_pb.UpdateLambdaRequest,
  responseType: lambda_pb.LambdaResponse
};

exports.LambdaService = LambdaService;

function LambdaServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

LambdaServiceClient.prototype.createFunction = function createFunction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LambdaService.CreateFunction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

LambdaServiceClient.prototype.testFunction = function testFunction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LambdaService.TestFunction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

LambdaServiceClient.prototype.invoqueFunction = function invoqueFunction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LambdaService.InvoqueFunction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

LambdaServiceClient.prototype.seedLambdaServer = function seedLambdaServer(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LambdaService.SeedLambdaServer, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

LambdaServiceClient.prototype.receiveEvents = function receiveEvents(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(LambdaService.ReceiveEvents, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

LambdaServiceClient.prototype.updateLambda = function updateLambda(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(LambdaService.UpdateLambda, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.LambdaServiceClient = LambdaServiceClient;

