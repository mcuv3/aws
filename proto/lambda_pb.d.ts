import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class CreateFunctionRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateFunctionRequest;

  getRuntime(): Runtime;
  setRuntime(value: Runtime): CreateFunctionRequest;

  getRam(): number;
  setRam(value: number): CreateFunctionRequest;

  getCode(): string;
  setCode(value: string): CreateFunctionRequest;

  getHandler(): string;
  setHandler(value: string): CreateFunctionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateFunctionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateFunctionRequest): CreateFunctionRequest.AsObject;
  static serializeBinaryToWriter(message: CreateFunctionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateFunctionRequest;
  static deserializeBinaryFromReader(message: CreateFunctionRequest, reader: jspb.BinaryReader): CreateFunctionRequest;
}

export namespace CreateFunctionRequest {
  export type AsObject = {
    name: string,
    runtime: Runtime,
    ram: number,
    code: string,
    handler: string,
  }
}

export class LambdaResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): LambdaResponse;

  getOk(): boolean;
  setOk(value: boolean): LambdaResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LambdaResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LambdaResponse): LambdaResponse.AsObject;
  static serializeBinaryToWriter(message: LambdaResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LambdaResponse;
  static deserializeBinaryFromReader(message: LambdaResponse, reader: jspb.BinaryReader): LambdaResponse;
}

export namespace LambdaResponse {
  export type AsObject = {
    message: string,
    ok: boolean,
  }
}

export class TestFunctionResquest extends jspb.Message {
  getFunctionName(): string;
  setFunctionName(value: string): TestFunctionResquest;

  getEventData(): string;
  setEventData(value: string): TestFunctionResquest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestFunctionResquest.AsObject;
  static toObject(includeInstance: boolean, msg: TestFunctionResquest): TestFunctionResquest.AsObject;
  static serializeBinaryToWriter(message: TestFunctionResquest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestFunctionResquest;
  static deserializeBinaryFromReader(message: TestFunctionResquest, reader: jspb.BinaryReader): TestFunctionResquest;
}

export namespace TestFunctionResquest {
  export type AsObject = {
    functionName: string,
    eventData: string,
  }
}

export class InvoqueFunctionRequest extends jspb.Message {
  getArn(): string;
  setArn(value: string): InvoqueFunctionRequest;

  getEventData(): string;
  setEventData(value: string): InvoqueFunctionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InvoqueFunctionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InvoqueFunctionRequest): InvoqueFunctionRequest.AsObject;
  static serializeBinaryToWriter(message: InvoqueFunctionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InvoqueFunctionRequest;
  static deserializeBinaryFromReader(message: InvoqueFunctionRequest, reader: jspb.BinaryReader): InvoqueFunctionRequest;
}

export namespace InvoqueFunctionRequest {
  export type AsObject = {
    arn: string,
    eventData: string,
  }
}

export class ReceiveEventRequest extends jspb.Message {
  getHash(): string;
  setHash(value: string): ReceiveEventRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveEventRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveEventRequest): ReceiveEventRequest.AsObject;
  static serializeBinaryToWriter(message: ReceiveEventRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveEventRequest;
  static deserializeBinaryFromReader(message: ReceiveEventRequest, reader: jspb.BinaryReader): ReceiveEventRequest;
}

export namespace ReceiveEventRequest {
  export type AsObject = {
    hash: string,
  }
}

export class EventResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): EventResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EventResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EventResponse): EventResponse.AsObject;
  static serializeBinaryToWriter(message: EventResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EventResponse;
  static deserializeBinaryFromReader(message: EventResponse, reader: jspb.BinaryReader): EventResponse;
}

export namespace EventResponse {
  export type AsObject = {
    message: string,
  }
}

export enum Runtime { 
  NODEJS14 = 0,
  PYTHON3 = 1,
  GOLANG = 2,
}
