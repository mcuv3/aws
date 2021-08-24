// package: sqs
// file: sqs.proto

import * as sqs_pb from "./sqs_pb";
import {grpc} from "@improbable-eng/grpc-web";

type SQSServiceCreateQueue = {
  readonly methodName: string;
  readonly service: typeof SQSService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof sqs_pb.CreateQueueRequest;
  readonly responseType: typeof sqs_pb.CreateQueueResponse;
};

type SQSServiceSendMessage = {
  readonly methodName: string;
  readonly service: typeof SQSService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof sqs_pb.SendMessageRequest;
  readonly responseType: typeof sqs_pb.SendMessageResponse;
};

type SQSServiceReceiveMessage = {
  readonly methodName: string;
  readonly service: typeof SQSService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof sqs_pb.ReceiveMessageRequest;
  readonly responseType: typeof sqs_pb.ReceiveMessageResponse;
};

type SQSServiceDeleteMessage = {
  readonly methodName: string;
  readonly service: typeof SQSService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof sqs_pb.DeleteMessageRequest;
  readonly responseType: typeof sqs_pb.DeleteResponse;
};

type SQSServiceDeleteQueue = {
  readonly methodName: string;
  readonly service: typeof SQSService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof sqs_pb.DeleteQueueRequest;
  readonly responseType: typeof sqs_pb.DeleteResponse;
};

export class SQSService {
  static readonly serviceName: string;
  static readonly CreateQueue: SQSServiceCreateQueue;
  static readonly SendMessage: SQSServiceSendMessage;
  static readonly ReceiveMessage: SQSServiceReceiveMessage;
  static readonly DeleteMessage: SQSServiceDeleteMessage;
  static readonly DeleteQueue: SQSServiceDeleteQueue;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class SQSServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createQueue(
    requestMessage: sqs_pb.CreateQueueRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.CreateQueueResponse|null) => void
  ): UnaryResponse;
  createQueue(
    requestMessage: sqs_pb.CreateQueueRequest,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.CreateQueueResponse|null) => void
  ): UnaryResponse;
  sendMessage(
    requestMessage: sqs_pb.SendMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.SendMessageResponse|null) => void
  ): UnaryResponse;
  sendMessage(
    requestMessage: sqs_pb.SendMessageRequest,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.SendMessageResponse|null) => void
  ): UnaryResponse;
  receiveMessage(requestMessage: sqs_pb.ReceiveMessageRequest, metadata?: grpc.Metadata): ResponseStream<sqs_pb.ReceiveMessageResponse>;
  deleteMessage(
    requestMessage: sqs_pb.DeleteMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  deleteMessage(
    requestMessage: sqs_pb.DeleteMessageRequest,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  deleteQueue(
    requestMessage: sqs_pb.DeleteQueueRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.DeleteResponse|null) => void
  ): UnaryResponse;
  deleteQueue(
    requestMessage: sqs_pb.DeleteQueueRequest,
    callback: (error: ServiceError|null, responseMessage: sqs_pb.DeleteResponse|null) => void
  ): UnaryResponse;
}

