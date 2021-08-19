/**
 * @fileoverview gRPC-Web generated client stub for sqs
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as sqs_pb from './sqs_pb';


export class SQSServiceClient {
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

  methodInfoCreateQueue = new grpcWeb.AbstractClientBase.MethodInfo(
    sqs_pb.CreateQueueResponse,
    (request: sqs_pb.CreateQueueRequest) => {
      return request.serializeBinary();
    },
    sqs_pb.CreateQueueResponse.deserializeBinary
  );

  createQueue(
    request: sqs_pb.CreateQueueRequest,
    metadata: grpcWeb.Metadata | null): Promise<sqs_pb.CreateQueueResponse>;

  createQueue(
    request: sqs_pb.CreateQueueRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: sqs_pb.CreateQueueResponse) => void): grpcWeb.ClientReadableStream<sqs_pb.CreateQueueResponse>;

  createQueue(
    request: sqs_pb.CreateQueueRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: sqs_pb.CreateQueueResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/sqs.SQSService/CreateQueue',
        request,
        metadata || {},
        this.methodInfoCreateQueue,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/sqs.SQSService/CreateQueue',
    request,
    metadata || {},
    this.methodInfoCreateQueue);
  }

  methodInfoSendMessage = new grpcWeb.AbstractClientBase.MethodInfo(
    sqs_pb.SendMessageResponse,
    (request: sqs_pb.SendMessageRequest) => {
      return request.serializeBinary();
    },
    sqs_pb.SendMessageResponse.deserializeBinary
  );

  sendMessage(
    request: sqs_pb.SendMessageRequest,
    metadata: grpcWeb.Metadata | null): Promise<sqs_pb.SendMessageResponse>;

  sendMessage(
    request: sqs_pb.SendMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: sqs_pb.SendMessageResponse) => void): grpcWeb.ClientReadableStream<sqs_pb.SendMessageResponse>;

  sendMessage(
    request: sqs_pb.SendMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: sqs_pb.SendMessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/sqs.SQSService/SendMessage',
        request,
        metadata || {},
        this.methodInfoSendMessage,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/sqs.SQSService/SendMessage',
    request,
    metadata || {},
    this.methodInfoSendMessage);
  }

  methodInfoReceiveMessage = new grpcWeb.AbstractClientBase.MethodInfo(
    sqs_pb.ReceiveMessageResponse,
    (request: sqs_pb.ReceiveMessageRequest) => {
      return request.serializeBinary();
    },
    sqs_pb.ReceiveMessageResponse.deserializeBinary
  );

  receiveMessage(
    request: sqs_pb.ReceiveMessageRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/sqs.SQSService/ReceiveMessage',
      request,
      metadata || {},
      this.methodInfoReceiveMessage);
  }

  methodInfoDeleteMessage = new grpcWeb.AbstractClientBase.MethodInfo(
    sqs_pb.DeleteResponse,
    (request: sqs_pb.DeleteMessageRequest) => {
      return request.serializeBinary();
    },
    sqs_pb.DeleteResponse.deserializeBinary
  );

  deleteMessage(
    request: sqs_pb.DeleteMessageRequest,
    metadata: grpcWeb.Metadata | null): Promise<sqs_pb.DeleteResponse>;

  deleteMessage(
    request: sqs_pb.DeleteMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: sqs_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<sqs_pb.DeleteResponse>;

  deleteMessage(
    request: sqs_pb.DeleteMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: sqs_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/sqs.SQSService/DeleteMessage',
        request,
        metadata || {},
        this.methodInfoDeleteMessage,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/sqs.SQSService/DeleteMessage',
    request,
    metadata || {},
    this.methodInfoDeleteMessage);
  }

  methodInfoDeleteQueue = new grpcWeb.AbstractClientBase.MethodInfo(
    sqs_pb.DeleteResponse,
    (request: sqs_pb.DeleteQueueRequest) => {
      return request.serializeBinary();
    },
    sqs_pb.DeleteResponse.deserializeBinary
  );

  deleteQueue(
    request: sqs_pb.DeleteQueueRequest,
    metadata: grpcWeb.Metadata | null): Promise<sqs_pb.DeleteResponse>;

  deleteQueue(
    request: sqs_pb.DeleteQueueRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: sqs_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<sqs_pb.DeleteResponse>;

  deleteQueue(
    request: sqs_pb.DeleteQueueRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: sqs_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/sqs.SQSService/DeleteQueue',
        request,
        metadata || {},
        this.methodInfoDeleteQueue,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/sqs.SQSService/DeleteQueue',
    request,
    metadata || {},
    this.methodInfoDeleteQueue);
  }

}

