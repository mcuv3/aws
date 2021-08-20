// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var sqs_pb = require('./sqs_pb.js');

function serialize_sqs_CreateQueueRequest(arg) {
  if (!(arg instanceof sqs_pb.CreateQueueRequest)) {
    throw new Error('Expected argument of type sqs.CreateQueueRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_CreateQueueRequest(buffer_arg) {
  return sqs_pb.CreateQueueRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_CreateQueueResponse(arg) {
  if (!(arg instanceof sqs_pb.CreateQueueResponse)) {
    throw new Error('Expected argument of type sqs.CreateQueueResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_CreateQueueResponse(buffer_arg) {
  return sqs_pb.CreateQueueResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_DeleteMessageRequest(arg) {
  if (!(arg instanceof sqs_pb.DeleteMessageRequest)) {
    throw new Error('Expected argument of type sqs.DeleteMessageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_DeleteMessageRequest(buffer_arg) {
  return sqs_pb.DeleteMessageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_DeleteQueueRequest(arg) {
  if (!(arg instanceof sqs_pb.DeleteQueueRequest)) {
    throw new Error('Expected argument of type sqs.DeleteQueueRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_DeleteQueueRequest(buffer_arg) {
  return sqs_pb.DeleteQueueRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_DeleteResponse(arg) {
  if (!(arg instanceof sqs_pb.DeleteResponse)) {
    throw new Error('Expected argument of type sqs.DeleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_DeleteResponse(buffer_arg) {
  return sqs_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_ReceiveMessageRequest(arg) {
  if (!(arg instanceof sqs_pb.ReceiveMessageRequest)) {
    throw new Error('Expected argument of type sqs.ReceiveMessageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_ReceiveMessageRequest(buffer_arg) {
  return sqs_pb.ReceiveMessageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_ReceiveMessageResponse(arg) {
  if (!(arg instanceof sqs_pb.ReceiveMessageResponse)) {
    throw new Error('Expected argument of type sqs.ReceiveMessageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_ReceiveMessageResponse(buffer_arg) {
  return sqs_pb.ReceiveMessageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_SendMessageRequest(arg) {
  if (!(arg instanceof sqs_pb.SendMessageRequest)) {
    throw new Error('Expected argument of type sqs.SendMessageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_SendMessageRequest(buffer_arg) {
  return sqs_pb.SendMessageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sqs_SendMessageResponse(arg) {
  if (!(arg instanceof sqs_pb.SendMessageResponse)) {
    throw new Error('Expected argument of type sqs.SendMessageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sqs_SendMessageResponse(buffer_arg) {
  return sqs_pb.SendMessageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var SQSServiceService = exports.SQSServiceService = {
  createQueue: {
    path: '/sqs.SQSService/CreateQueue',
    requestStream: false,
    responseStream: false,
    requestType: sqs_pb.CreateQueueRequest,
    responseType: sqs_pb.CreateQueueResponse,
    requestSerialize: serialize_sqs_CreateQueueRequest,
    requestDeserialize: deserialize_sqs_CreateQueueRequest,
    responseSerialize: serialize_sqs_CreateQueueResponse,
    responseDeserialize: deserialize_sqs_CreateQueueResponse,
  },
  sendMessage: {
    path: '/sqs.SQSService/SendMessage',
    requestStream: false,
    responseStream: false,
    requestType: sqs_pb.SendMessageRequest,
    responseType: sqs_pb.SendMessageResponse,
    requestSerialize: serialize_sqs_SendMessageRequest,
    requestDeserialize: deserialize_sqs_SendMessageRequest,
    responseSerialize: serialize_sqs_SendMessageResponse,
    responseDeserialize: deserialize_sqs_SendMessageResponse,
  },
  receiveMessage: {
    path: '/sqs.SQSService/ReceiveMessage',
    requestStream: false,
    responseStream: true,
    requestType: sqs_pb.ReceiveMessageRequest,
    responseType: sqs_pb.ReceiveMessageResponse,
    requestSerialize: serialize_sqs_ReceiveMessageRequest,
    requestDeserialize: deserialize_sqs_ReceiveMessageRequest,
    responseSerialize: serialize_sqs_ReceiveMessageResponse,
    responseDeserialize: deserialize_sqs_ReceiveMessageResponse,
  },
  deleteMessage: {
    path: '/sqs.SQSService/DeleteMessage',
    requestStream: false,
    responseStream: false,
    requestType: sqs_pb.DeleteMessageRequest,
    responseType: sqs_pb.DeleteResponse,
    requestSerialize: serialize_sqs_DeleteMessageRequest,
    requestDeserialize: deserialize_sqs_DeleteMessageRequest,
    responseSerialize: serialize_sqs_DeleteResponse,
    responseDeserialize: deserialize_sqs_DeleteResponse,
  },
  deleteQueue: {
    path: '/sqs.SQSService/DeleteQueue',
    requestStream: false,
    responseStream: false,
    requestType: sqs_pb.DeleteQueueRequest,
    responseType: sqs_pb.DeleteResponse,
    requestSerialize: serialize_sqs_DeleteQueueRequest,
    requestDeserialize: deserialize_sqs_DeleteQueueRequest,
    responseSerialize: serialize_sqs_DeleteResponse,
    responseDeserialize: deserialize_sqs_DeleteResponse,
  },
};

exports.SQSServiceClient = grpc.makeGenericClientConstructor(SQSServiceService);
