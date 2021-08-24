// package: sqs
// file: sqs.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as sqs_pb from "./sqs_pb";

interface ISQSServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createQueue: ISQSServiceService_ICreateQueue;
    sendMessage: ISQSServiceService_ISendMessage;
    receiveMessage: ISQSServiceService_IReceiveMessage;
    deleteMessage: ISQSServiceService_IDeleteMessage;
    deleteQueue: ISQSServiceService_IDeleteQueue;
}

interface ISQSServiceService_ICreateQueue extends grpc.MethodDefinition<sqs_pb.CreateQueueRequest, sqs_pb.CreateQueueResponse> {
    path: "/sqs.SQSService/CreateQueue";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sqs_pb.CreateQueueRequest>;
    requestDeserialize: grpc.deserialize<sqs_pb.CreateQueueRequest>;
    responseSerialize: grpc.serialize<sqs_pb.CreateQueueResponse>;
    responseDeserialize: grpc.deserialize<sqs_pb.CreateQueueResponse>;
}
interface ISQSServiceService_ISendMessage extends grpc.MethodDefinition<sqs_pb.SendMessageRequest, sqs_pb.SendMessageResponse> {
    path: "/sqs.SQSService/SendMessage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sqs_pb.SendMessageRequest>;
    requestDeserialize: grpc.deserialize<sqs_pb.SendMessageRequest>;
    responseSerialize: grpc.serialize<sqs_pb.SendMessageResponse>;
    responseDeserialize: grpc.deserialize<sqs_pb.SendMessageResponse>;
}
interface ISQSServiceService_IReceiveMessage extends grpc.MethodDefinition<sqs_pb.ReceiveMessageRequest, sqs_pb.ReceiveMessageResponse> {
    path: "/sqs.SQSService/ReceiveMessage";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<sqs_pb.ReceiveMessageRequest>;
    requestDeserialize: grpc.deserialize<sqs_pb.ReceiveMessageRequest>;
    responseSerialize: grpc.serialize<sqs_pb.ReceiveMessageResponse>;
    responseDeserialize: grpc.deserialize<sqs_pb.ReceiveMessageResponse>;
}
interface ISQSServiceService_IDeleteMessage extends grpc.MethodDefinition<sqs_pb.DeleteMessageRequest, sqs_pb.DeleteResponse> {
    path: "/sqs.SQSService/DeleteMessage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sqs_pb.DeleteMessageRequest>;
    requestDeserialize: grpc.deserialize<sqs_pb.DeleteMessageRequest>;
    responseSerialize: grpc.serialize<sqs_pb.DeleteResponse>;
    responseDeserialize: grpc.deserialize<sqs_pb.DeleteResponse>;
}
interface ISQSServiceService_IDeleteQueue extends grpc.MethodDefinition<sqs_pb.DeleteQueueRequest, sqs_pb.DeleteResponse> {
    path: "/sqs.SQSService/DeleteQueue";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<sqs_pb.DeleteQueueRequest>;
    requestDeserialize: grpc.deserialize<sqs_pb.DeleteQueueRequest>;
    responseSerialize: grpc.serialize<sqs_pb.DeleteResponse>;
    responseDeserialize: grpc.deserialize<sqs_pb.DeleteResponse>;
}

export const SQSServiceService: ISQSServiceService;

export interface ISQSServiceServer {
    createQueue: grpc.handleUnaryCall<sqs_pb.CreateQueueRequest, sqs_pb.CreateQueueResponse>;
    sendMessage: grpc.handleUnaryCall<sqs_pb.SendMessageRequest, sqs_pb.SendMessageResponse>;
    receiveMessage: grpc.handleServerStreamingCall<sqs_pb.ReceiveMessageRequest, sqs_pb.ReceiveMessageResponse>;
    deleteMessage: grpc.handleUnaryCall<sqs_pb.DeleteMessageRequest, sqs_pb.DeleteResponse>;
    deleteQueue: grpc.handleUnaryCall<sqs_pb.DeleteQueueRequest, sqs_pb.DeleteResponse>;
}

export interface ISQSServiceClient {
    createQueue(request: sqs_pb.CreateQueueRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.CreateQueueResponse) => void): grpc.ClientUnaryCall;
    createQueue(request: sqs_pb.CreateQueueRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.CreateQueueResponse) => void): grpc.ClientUnaryCall;
    createQueue(request: sqs_pb.CreateQueueRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.CreateQueueResponse) => void): grpc.ClientUnaryCall;
    sendMessage(request: sqs_pb.SendMessageRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.SendMessageResponse) => void): grpc.ClientUnaryCall;
    sendMessage(request: sqs_pb.SendMessageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.SendMessageResponse) => void): grpc.ClientUnaryCall;
    sendMessage(request: sqs_pb.SendMessageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.SendMessageResponse) => void): grpc.ClientUnaryCall;
    receiveMessage(request: sqs_pb.ReceiveMessageRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<sqs_pb.ReceiveMessageResponse>;
    receiveMessage(request: sqs_pb.ReceiveMessageRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<sqs_pb.ReceiveMessageResponse>;
    deleteMessage(request: sqs_pb.DeleteMessageRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteMessage(request: sqs_pb.DeleteMessageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteMessage(request: sqs_pb.DeleteMessageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteQueue(request: sqs_pb.DeleteQueueRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteQueue(request: sqs_pb.DeleteQueueRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    deleteQueue(request: sqs_pb.DeleteQueueRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
}

export class SQSServiceClient extends grpc.Client implements ISQSServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createQueue(request: sqs_pb.CreateQueueRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.CreateQueueResponse) => void): grpc.ClientUnaryCall;
    public createQueue(request: sqs_pb.CreateQueueRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.CreateQueueResponse) => void): grpc.ClientUnaryCall;
    public createQueue(request: sqs_pb.CreateQueueRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.CreateQueueResponse) => void): grpc.ClientUnaryCall;
    public sendMessage(request: sqs_pb.SendMessageRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.SendMessageResponse) => void): grpc.ClientUnaryCall;
    public sendMessage(request: sqs_pb.SendMessageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.SendMessageResponse) => void): grpc.ClientUnaryCall;
    public sendMessage(request: sqs_pb.SendMessageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.SendMessageResponse) => void): grpc.ClientUnaryCall;
    public receiveMessage(request: sqs_pb.ReceiveMessageRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<sqs_pb.ReceiveMessageResponse>;
    public receiveMessage(request: sqs_pb.ReceiveMessageRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<sqs_pb.ReceiveMessageResponse>;
    public deleteMessage(request: sqs_pb.DeleteMessageRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteMessage(request: sqs_pb.DeleteMessageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteMessage(request: sqs_pb.DeleteMessageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteQueue(request: sqs_pb.DeleteQueueRequest, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteQueue(request: sqs_pb.DeleteQueueRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public deleteQueue(request: sqs_pb.DeleteQueueRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: sqs_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
}
