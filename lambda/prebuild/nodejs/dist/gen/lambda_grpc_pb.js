// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var lambda_pb = require('./lambda_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_lambda_CreateFunctionRequest(arg) {
  if (!(arg instanceof lambda_pb.CreateFunctionRequest)) {
    throw new Error('Expected argument of type lambda.CreateFunctionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_lambda_CreateFunctionRequest(buffer_arg) {
  return lambda_pb.CreateFunctionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_lambda_EventResponse(arg) {
  if (!(arg instanceof lambda_pb.EventResponse)) {
    throw new Error('Expected argument of type lambda.EventResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_lambda_EventResponse(buffer_arg) {
  return lambda_pb.EventResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_lambda_InvoqueFunctionRequest(arg) {
  if (!(arg instanceof lambda_pb.InvoqueFunctionRequest)) {
    throw new Error('Expected argument of type lambda.InvoqueFunctionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_lambda_InvoqueFunctionRequest(buffer_arg) {
  return lambda_pb.InvoqueFunctionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_lambda_LambdaResponse(arg) {
  if (!(arg instanceof lambda_pb.LambdaResponse)) {
    throw new Error('Expected argument of type lambda.LambdaResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_lambda_LambdaResponse(buffer_arg) {
  return lambda_pb.LambdaResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_lambda_ReceiveEventRequest(arg) {
  if (!(arg instanceof lambda_pb.ReceiveEventRequest)) {
    throw new Error('Expected argument of type lambda.ReceiveEventRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_lambda_ReceiveEventRequest(buffer_arg) {
  return lambda_pb.ReceiveEventRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_lambda_TestFunctionResquest(arg) {
  if (!(arg instanceof lambda_pb.TestFunctionResquest)) {
    throw new Error('Expected argument of type lambda.TestFunctionResquest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_lambda_TestFunctionResquest(buffer_arg) {
  return lambda_pb.TestFunctionResquest.deserializeBinary(new Uint8Array(buffer_arg));
}


var LambdaServiceService = exports.LambdaServiceService = {
  createFunction: {
    path: '/lambda.LambdaService/CreateFunction',
    requestStream: false,
    responseStream: false,
    requestType: lambda_pb.CreateFunctionRequest,
    responseType: lambda_pb.LambdaResponse,
    requestSerialize: serialize_lambda_CreateFunctionRequest,
    requestDeserialize: deserialize_lambda_CreateFunctionRequest,
    responseSerialize: serialize_lambda_LambdaResponse,
    responseDeserialize: deserialize_lambda_LambdaResponse,
  },
  testFunction: {
    path: '/lambda.LambdaService/TestFunction',
    requestStream: false,
    responseStream: false,
    requestType: lambda_pb.TestFunctionResquest,
    responseType: lambda_pb.LambdaResponse,
    requestSerialize: serialize_lambda_TestFunctionResquest,
    requestDeserialize: deserialize_lambda_TestFunctionResquest,
    responseSerialize: serialize_lambda_LambdaResponse,
    responseDeserialize: deserialize_lambda_LambdaResponse,
  },
  invoqueFunction: {
    path: '/lambda.LambdaService/InvoqueFunction',
    requestStream: false,
    responseStream: false,
    requestType: lambda_pb.InvoqueFunctionRequest,
    responseType: lambda_pb.LambdaResponse,
    requestSerialize: serialize_lambda_InvoqueFunctionRequest,
    requestDeserialize: deserialize_lambda_InvoqueFunctionRequest,
    responseSerialize: serialize_lambda_LambdaResponse,
    responseDeserialize: deserialize_lambda_LambdaResponse,
  },
  seedLambdaServer: {
    path: '/lambda.LambdaService/SeedLambdaServer',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: lambda_pb.LambdaResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_lambda_LambdaResponse,
    responseDeserialize: deserialize_lambda_LambdaResponse,
  },
  receiveEvents: {
    path: '/lambda.LambdaService/ReceiveEvents',
    requestStream: false,
    responseStream: true,
    requestType: lambda_pb.ReceiveEventRequest,
    responseType: lambda_pb.EventResponse,
    requestSerialize: serialize_lambda_ReceiveEventRequest,
    requestDeserialize: deserialize_lambda_ReceiveEventRequest,
    responseSerialize: serialize_lambda_EventResponse,
    responseDeserialize: deserialize_lambda_EventResponse,
  },
};

exports.LambdaServiceClient = grpc.makeGenericClientConstructor(LambdaServiceService);
