// package: lambda
// file: lambda.proto

import * as lambda_pb from "./lambda_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type LambdaServiceCreateFunction = {
  readonly methodName: string;
  readonly service: typeof LambdaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof lambda_pb.CreateFunctionRequest;
  readonly responseType: typeof lambda_pb.LambdaResponse;
};

type LambdaServiceTestFunction = {
  readonly methodName: string;
  readonly service: typeof LambdaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof lambda_pb.TestFunctionResquest;
  readonly responseType: typeof lambda_pb.LambdaResponse;
};

type LambdaServiceInvoqueFunction = {
  readonly methodName: string;
  readonly service: typeof LambdaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof lambda_pb.InvoqueFunctionRequest;
  readonly responseType: typeof lambda_pb.LambdaResponse;
};

type LambdaServiceSeedLambdaServer = {
  readonly methodName: string;
  readonly service: typeof LambdaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof lambda_pb.LambdaResponse;
};

type LambdaServiceReceiveEvents = {
  readonly methodName: string;
  readonly service: typeof LambdaService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof lambda_pb.ReceiveEventRequest;
  readonly responseType: typeof lambda_pb.EventResponse;
};

type LambdaServiceUpdateLambda = {
  readonly methodName: string;
  readonly service: typeof LambdaService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof lambda_pb.UpdateLambdaRequest;
  readonly responseType: typeof lambda_pb.LambdaResponse;
};

export class LambdaService {
  static readonly serviceName: string;
  static readonly CreateFunction: LambdaServiceCreateFunction;
  static readonly TestFunction: LambdaServiceTestFunction;
  static readonly InvoqueFunction: LambdaServiceInvoqueFunction;
  static readonly SeedLambdaServer: LambdaServiceSeedLambdaServer;
  static readonly ReceiveEvents: LambdaServiceReceiveEvents;
  static readonly UpdateLambda: LambdaServiceUpdateLambda;
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

export class LambdaServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createFunction(
    requestMessage: lambda_pb.CreateFunctionRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  createFunction(
    requestMessage: lambda_pb.CreateFunctionRequest,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  testFunction(
    requestMessage: lambda_pb.TestFunctionResquest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  testFunction(
    requestMessage: lambda_pb.TestFunctionResquest,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  invoqueFunction(
    requestMessage: lambda_pb.InvoqueFunctionRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  invoqueFunction(
    requestMessage: lambda_pb.InvoqueFunctionRequest,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  seedLambdaServer(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  seedLambdaServer(
    requestMessage: google_protobuf_empty_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  receiveEvents(requestMessage: lambda_pb.ReceiveEventRequest, metadata?: grpc.Metadata): ResponseStream<lambda_pb.EventResponse>;
  updateLambda(
    requestMessage: lambda_pb.UpdateLambdaRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
  updateLambda(
    requestMessage: lambda_pb.UpdateLambdaRequest,
    callback: (error: ServiceError|null, responseMessage: lambda_pb.LambdaResponse|null) => void
  ): UnaryResponse;
}

