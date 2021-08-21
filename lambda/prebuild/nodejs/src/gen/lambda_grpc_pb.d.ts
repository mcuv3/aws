// package: lambda
// file: lambda.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as lambda_pb from "./lambda_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface ILambdaServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createFunction: ILambdaServiceService_ICreateFunction;
    testFunction: ILambdaServiceService_ITestFunction;
    invoqueFunction: ILambdaServiceService_IInvoqueFunction;
    seedLambdaServer: ILambdaServiceService_ISeedLambdaServer;
    receiveEvents: ILambdaServiceService_IReceiveEvents;
    updateLambda: ILambdaServiceService_IUpdateLambda;
}

interface ILambdaServiceService_ICreateFunction extends grpc.MethodDefinition<lambda_pb.CreateFunctionRequest, lambda_pb.LambdaResponse> {
    path: "/lambda.LambdaService/CreateFunction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<lambda_pb.CreateFunctionRequest>;
    requestDeserialize: grpc.deserialize<lambda_pb.CreateFunctionRequest>;
    responseSerialize: grpc.serialize<lambda_pb.LambdaResponse>;
    responseDeserialize: grpc.deserialize<lambda_pb.LambdaResponse>;
}
interface ILambdaServiceService_ITestFunction extends grpc.MethodDefinition<lambda_pb.TestFunctionResquest, lambda_pb.LambdaResponse> {
    path: "/lambda.LambdaService/TestFunction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<lambda_pb.TestFunctionResquest>;
    requestDeserialize: grpc.deserialize<lambda_pb.TestFunctionResquest>;
    responseSerialize: grpc.serialize<lambda_pb.LambdaResponse>;
    responseDeserialize: grpc.deserialize<lambda_pb.LambdaResponse>;
}
interface ILambdaServiceService_IInvoqueFunction extends grpc.MethodDefinition<lambda_pb.InvoqueFunctionRequest, lambda_pb.LambdaResponse> {
    path: "/lambda.LambdaService/InvoqueFunction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<lambda_pb.InvoqueFunctionRequest>;
    requestDeserialize: grpc.deserialize<lambda_pb.InvoqueFunctionRequest>;
    responseSerialize: grpc.serialize<lambda_pb.LambdaResponse>;
    responseDeserialize: grpc.deserialize<lambda_pb.LambdaResponse>;
}
interface ILambdaServiceService_ISeedLambdaServer extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, lambda_pb.LambdaResponse> {
    path: "/lambda.LambdaService/SeedLambdaServer";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<lambda_pb.LambdaResponse>;
    responseDeserialize: grpc.deserialize<lambda_pb.LambdaResponse>;
}
interface ILambdaServiceService_IReceiveEvents extends grpc.MethodDefinition<lambda_pb.ReceiveEventRequest, lambda_pb.EventResponse> {
    path: "/lambda.LambdaService/ReceiveEvents";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<lambda_pb.ReceiveEventRequest>;
    requestDeserialize: grpc.deserialize<lambda_pb.ReceiveEventRequest>;
    responseSerialize: grpc.serialize<lambda_pb.EventResponse>;
    responseDeserialize: grpc.deserialize<lambda_pb.EventResponse>;
}
interface ILambdaServiceService_IUpdateLambda extends grpc.MethodDefinition<lambda_pb.UpdateLambdaRequest, lambda_pb.LambdaResponse> {
    path: "/lambda.LambdaService/UpdateLambda";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<lambda_pb.UpdateLambdaRequest>;
    requestDeserialize: grpc.deserialize<lambda_pb.UpdateLambdaRequest>;
    responseSerialize: grpc.serialize<lambda_pb.LambdaResponse>;
    responseDeserialize: grpc.deserialize<lambda_pb.LambdaResponse>;
}

export const LambdaServiceService: ILambdaServiceService;

export interface ILambdaServiceServer extends grpc.UntypedServiceImplementation {
    createFunction: grpc.handleUnaryCall<lambda_pb.CreateFunctionRequest, lambda_pb.LambdaResponse>;
    testFunction: grpc.handleUnaryCall<lambda_pb.TestFunctionResquest, lambda_pb.LambdaResponse>;
    invoqueFunction: grpc.handleUnaryCall<lambda_pb.InvoqueFunctionRequest, lambda_pb.LambdaResponse>;
    seedLambdaServer: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, lambda_pb.LambdaResponse>;
    receiveEvents: grpc.handleServerStreamingCall<lambda_pb.ReceiveEventRequest, lambda_pb.EventResponse>;
    updateLambda: grpc.handleUnaryCall<lambda_pb.UpdateLambdaRequest, lambda_pb.LambdaResponse>;
}

export interface ILambdaServiceClient {
    createFunction(request: lambda_pb.CreateFunctionRequest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    createFunction(request: lambda_pb.CreateFunctionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    createFunction(request: lambda_pb.CreateFunctionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    testFunction(request: lambda_pb.TestFunctionResquest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    testFunction(request: lambda_pb.TestFunctionResquest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    testFunction(request: lambda_pb.TestFunctionResquest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    invoqueFunction(request: lambda_pb.InvoqueFunctionRequest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    invoqueFunction(request: lambda_pb.InvoqueFunctionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    invoqueFunction(request: lambda_pb.InvoqueFunctionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    seedLambdaServer(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    seedLambdaServer(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    seedLambdaServer(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    receiveEvents(request: lambda_pb.ReceiveEventRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<lambda_pb.EventResponse>;
    receiveEvents(request: lambda_pb.ReceiveEventRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<lambda_pb.EventResponse>;
    updateLambda(request: lambda_pb.UpdateLambdaRequest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    updateLambda(request: lambda_pb.UpdateLambdaRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    updateLambda(request: lambda_pb.UpdateLambdaRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
}

export class LambdaServiceClient extends grpc.Client implements ILambdaServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createFunction(request: lambda_pb.CreateFunctionRequest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public createFunction(request: lambda_pb.CreateFunctionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public createFunction(request: lambda_pb.CreateFunctionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public testFunction(request: lambda_pb.TestFunctionResquest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public testFunction(request: lambda_pb.TestFunctionResquest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public testFunction(request: lambda_pb.TestFunctionResquest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public invoqueFunction(request: lambda_pb.InvoqueFunctionRequest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public invoqueFunction(request: lambda_pb.InvoqueFunctionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public invoqueFunction(request: lambda_pb.InvoqueFunctionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public seedLambdaServer(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public seedLambdaServer(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public seedLambdaServer(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public receiveEvents(request: lambda_pb.ReceiveEventRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<lambda_pb.EventResponse>;
    public receiveEvents(request: lambda_pb.ReceiveEventRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<lambda_pb.EventResponse>;
    public updateLambda(request: lambda_pb.UpdateLambdaRequest, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public updateLambda(request: lambda_pb.UpdateLambdaRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
    public updateLambda(request: lambda_pb.UpdateLambdaRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: lambda_pb.LambdaResponse) => void): grpc.ClientUnaryCall;
}
