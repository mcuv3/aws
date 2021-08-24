// package: sqs
// file: sqs.proto

var sqs_pb = require("./sqs_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var SQSService = (function () {
  function SQSService() {}
  SQSService.serviceName = "sqs.SQSService";
  return SQSService;
}());

SQSService.CreateQueue = {
  methodName: "CreateQueue",
  service: SQSService,
  requestStream: false,
  responseStream: false,
  requestType: sqs_pb.CreateQueueRequest,
  responseType: sqs_pb.CreateQueueResponse
};

SQSService.SendMessage = {
  methodName: "SendMessage",
  service: SQSService,
  requestStream: false,
  responseStream: false,
  requestType: sqs_pb.SendMessageRequest,
  responseType: sqs_pb.SendMessageResponse
};

SQSService.ReceiveMessage = {
  methodName: "ReceiveMessage",
  service: SQSService,
  requestStream: false,
  responseStream: true,
  requestType: sqs_pb.ReceiveMessageRequest,
  responseType: sqs_pb.ReceiveMessageResponse
};

SQSService.DeleteMessage = {
  methodName: "DeleteMessage",
  service: SQSService,
  requestStream: false,
  responseStream: false,
  requestType: sqs_pb.DeleteMessageRequest,
  responseType: sqs_pb.DeleteResponse
};

SQSService.DeleteQueue = {
  methodName: "DeleteQueue",
  service: SQSService,
  requestStream: false,
  responseStream: false,
  requestType: sqs_pb.DeleteQueueRequest,
  responseType: sqs_pb.DeleteResponse
};

exports.SQSService = SQSService;

function SQSServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

SQSServiceClient.prototype.createQueue = function createQueue(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SQSService.CreateQueue, {
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

SQSServiceClient.prototype.sendMessage = function sendMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SQSService.SendMessage, {
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

SQSServiceClient.prototype.receiveMessage = function receiveMessage(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(SQSService.ReceiveMessage, {
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

SQSServiceClient.prototype.deleteMessage = function deleteMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SQSService.DeleteMessage, {
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

SQSServiceClient.prototype.deleteQueue = function deleteQueue(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SQSService.DeleteQueue, {
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

exports.SQSServiceClient = SQSServiceClient;

