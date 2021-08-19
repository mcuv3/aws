/**
 * @fileoverview gRPC-Web generated client stub for lambda
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as lambda_pb from './lambda_pb';


export class LambdaServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCreateFunction = new grpcWeb.AbstractClientBase.MethodInfo(
    lambda_pb.LambdaResponse,
    (request: lambda_pb.CreateFunctionRequest) => {
      return request.serializeBinary();
    },
    lambda_pb.LambdaResponse.deserializeBinary
  );

  createFunction(
    request: lambda_pb.CreateFunctionRequest,
    metadata: grpcWeb.Metadata | null): Promise<lambda_pb.LambdaResponse>;

  createFunction(
    request: lambda_pb.CreateFunctionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void): grpcWeb.ClientReadableStream<lambda_pb.LambdaResponse>;

  createFunction(
    request: lambda_pb.CreateFunctionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lambda.LambdaService/CreateFunction',
        request,
        metadata || {},
        this.methodInfoCreateFunction,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lambda.LambdaService/CreateFunction',
    request,
    metadata || {},
    this.methodInfoCreateFunction);
  }

  methodInfoTestFunction = new grpcWeb.AbstractClientBase.MethodInfo(
    lambda_pb.LambdaResponse,
    (request: lambda_pb.TestFunctionResquest) => {
      return request.serializeBinary();
    },
    lambda_pb.LambdaResponse.deserializeBinary
  );

  testFunction(
    request: lambda_pb.TestFunctionResquest,
    metadata: grpcWeb.Metadata | null): Promise<lambda_pb.LambdaResponse>;

  testFunction(
    request: lambda_pb.TestFunctionResquest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void): grpcWeb.ClientReadableStream<lambda_pb.LambdaResponse>;

  testFunction(
    request: lambda_pb.TestFunctionResquest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lambda.LambdaService/TestFunction',
        request,
        metadata || {},
        this.methodInfoTestFunction,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lambda.LambdaService/TestFunction',
    request,
    metadata || {},
    this.methodInfoTestFunction);
  }

  methodInfoInvoqueFunction = new grpcWeb.AbstractClientBase.MethodInfo(
    lambda_pb.LambdaResponse,
    (request: lambda_pb.InvoqueFunctionRequest) => {
      return request.serializeBinary();
    },
    lambda_pb.LambdaResponse.deserializeBinary
  );

  invoqueFunction(
    request: lambda_pb.InvoqueFunctionRequest,
    metadata: grpcWeb.Metadata | null): Promise<lambda_pb.LambdaResponse>;

  invoqueFunction(
    request: lambda_pb.InvoqueFunctionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void): grpcWeb.ClientReadableStream<lambda_pb.LambdaResponse>;

  invoqueFunction(
    request: lambda_pb.InvoqueFunctionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lambda.LambdaService/InvoqueFunction',
        request,
        metadata || {},
        this.methodInfoInvoqueFunction,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lambda.LambdaService/InvoqueFunction',
    request,
    metadata || {},
    this.methodInfoInvoqueFunction);
  }

  methodInfoSeedLambdaServer = new grpcWeb.AbstractClientBase.MethodInfo(
    lambda_pb.LambdaResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    lambda_pb.LambdaResponse.deserializeBinary
  );

  seedLambdaServer(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<lambda_pb.LambdaResponse>;

  seedLambdaServer(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void): grpcWeb.ClientReadableStream<lambda_pb.LambdaResponse>;

  seedLambdaServer(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: lambda_pb.LambdaResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/lambda.LambdaService/SeedLambdaServer',
        request,
        metadata || {},
        this.methodInfoSeedLambdaServer,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/lambda.LambdaService/SeedLambdaServer',
    request,
    metadata || {},
    this.methodInfoSeedLambdaServer);
  }

  methodInfoReceiveEvents = new grpcWeb.AbstractClientBase.MethodInfo(
    lambda_pb.EventResponse,
    (request: lambda_pb.ReceiveEventRequest) => {
      return request.serializeBinary();
    },
    lambda_pb.EventResponse.deserializeBinary
  );

  receiveEvents(
    request: lambda_pb.ReceiveEventRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/lambda.LambdaService/ReceiveEvents',
      request,
      metadata || {},
      this.methodInfoReceiveEvents);
  }

}

