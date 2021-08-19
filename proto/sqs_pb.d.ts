import * as jspb from 'google-protobuf'



export class Message extends jspb.Message {
  getId(): string;
  setId(value: string): Message;

  getMessage(): string;
  setMessage(value: string): Message;

  getQueueName(): string;
  setQueueName(value: string): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    id: string,
    message: string,
    queueName: string,
  }
}

export class Queue extends jspb.Message {
  getName(): string;
  setName(value: string): Queue;

  getArn(): string;
  setArn(value: string): Queue;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Queue.AsObject;
  static toObject(includeInstance: boolean, msg: Queue): Queue.AsObject;
  static serializeBinaryToWriter(message: Queue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Queue;
  static deserializeBinaryFromReader(message: Queue, reader: jspb.BinaryReader): Queue;
}

export namespace Queue {
  export type AsObject = {
    name: string,
    arn: string,
  }
}

export class ConfigurationQueue extends jspb.Message {
  getVisibilitytimeout(): number;
  setVisibilitytimeout(value: number): ConfigurationQueue;

  getVisibilitytime(): Time;
  setVisibilitytime(value: Time): ConfigurationQueue;

  getMessageretention(): number;
  setMessageretention(value: number): ConfigurationQueue;

  getMessageretentiontime(): Time;
  setMessageretentiontime(value: Time): ConfigurationQueue;

  getDeliverydelay(): number;
  setDeliverydelay(value: number): ConfigurationQueue;

  getDeliverydelaytime(): Time;
  setDeliverydelaytime(value: Time): ConfigurationQueue;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConfigurationQueue.AsObject;
  static toObject(includeInstance: boolean, msg: ConfigurationQueue): ConfigurationQueue.AsObject;
  static serializeBinaryToWriter(message: ConfigurationQueue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConfigurationQueue;
  static deserializeBinaryFromReader(message: ConfigurationQueue, reader: jspb.BinaryReader): ConfigurationQueue;
}

export namespace ConfigurationQueue {
  export type AsObject = {
    visibilitytimeout: number,
    visibilitytime: Time,
    messageretention: number,
    messageretentiontime: Time,
    deliverydelay: number,
    deliverydelaytime: Time,
  }
}

export class CreateQueueRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateQueueRequest;

  getConf(): ConfigurationQueue | undefined;
  setConf(value?: ConfigurationQueue): CreateQueueRequest;
  hasConf(): boolean;
  clearConf(): CreateQueueRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateQueueRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateQueueRequest): CreateQueueRequest.AsObject;
  static serializeBinaryToWriter(message: CreateQueueRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateQueueRequest;
  static deserializeBinaryFromReader(message: CreateQueueRequest, reader: jspb.BinaryReader): CreateQueueRequest;
}

export namespace CreateQueueRequest {
  export type AsObject = {
    name: string,
    conf?: ConfigurationQueue.AsObject,
  }
}

export class CreateQueueResponse extends jspb.Message {
  getQueue(): Queue | undefined;
  setQueue(value?: Queue): CreateQueueResponse;
  hasQueue(): boolean;
  clearQueue(): CreateQueueResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateQueueResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateQueueResponse): CreateQueueResponse.AsObject;
  static serializeBinaryToWriter(message: CreateQueueResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateQueueResponse;
  static deserializeBinaryFromReader(message: CreateQueueResponse, reader: jspb.BinaryReader): CreateQueueResponse;
}

export namespace CreateQueueResponse {
  export type AsObject = {
    queue?: Queue.AsObject,
  }
}

export class SendMessageRequest extends jspb.Message {
  getQueueName(): string;
  setQueueName(value: string): SendMessageRequest;

  getMessage(): string;
  setMessage(value: string): SendMessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SendMessageRequest): SendMessageRequest.AsObject;
  static serializeBinaryToWriter(message: SendMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendMessageRequest;
  static deserializeBinaryFromReader(message: SendMessageRequest, reader: jspb.BinaryReader): SendMessageRequest;
}

export namespace SendMessageRequest {
  export type AsObject = {
    queueName: string,
    message: string,
  }
}

export class SendMessageResponse extends jspb.Message {
  getMessage(): Message | undefined;
  setMessage(value?: Message): SendMessageResponse;
  hasMessage(): boolean;
  clearMessage(): SendMessageResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendMessageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SendMessageResponse): SendMessageResponse.AsObject;
  static serializeBinaryToWriter(message: SendMessageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendMessageResponse;
  static deserializeBinaryFromReader(message: SendMessageResponse, reader: jspb.BinaryReader): SendMessageResponse;
}

export namespace SendMessageResponse {
  export type AsObject = {
    message?: Message.AsObject,
  }
}

export class ReceiveMessageRequest extends jspb.Message {
  getQueueName(): string;
  setQueueName(value: string): ReceiveMessageRequest;

  getBatchLimit(): number;
  setBatchLimit(value: number): ReceiveMessageRequest;

  getWaitLimit(): number;
  setWaitLimit(value: number): ReceiveMessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveMessageRequest): ReceiveMessageRequest.AsObject;
  static serializeBinaryToWriter(message: ReceiveMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveMessageRequest;
  static deserializeBinaryFromReader(message: ReceiveMessageRequest, reader: jspb.BinaryReader): ReceiveMessageRequest;
}

export namespace ReceiveMessageRequest {
  export type AsObject = {
    queueName: string,
    batchLimit: number,
    waitLimit: number,
  }
}

export class ReceiveMessageResponse extends jspb.Message {
  getMessage(): Message | undefined;
  setMessage(value?: Message): ReceiveMessageResponse;
  hasMessage(): boolean;
  clearMessage(): ReceiveMessageResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveMessageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveMessageResponse): ReceiveMessageResponse.AsObject;
  static serializeBinaryToWriter(message: ReceiveMessageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveMessageResponse;
  static deserializeBinaryFromReader(message: ReceiveMessageResponse, reader: jspb.BinaryReader): ReceiveMessageResponse;
}

export namespace ReceiveMessageResponse {
  export type AsObject = {
    message?: Message.AsObject,
  }
}

export class DeleteMessageRequest extends jspb.Message {
  getQueueName(): string;
  setQueueName(value: string): DeleteMessageRequest;

  getMessageId(): string;
  setMessageId(value: string): DeleteMessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteMessageRequest): DeleteMessageRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteMessageRequest;
  static deserializeBinaryFromReader(message: DeleteMessageRequest, reader: jspb.BinaryReader): DeleteMessageRequest;
}

export namespace DeleteMessageRequest {
  export type AsObject = {
    queueName: string,
    messageId: string,
  }
}

export class DeleteResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): DeleteResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteResponse;
  static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
  export type AsObject = {
    message: string,
  }
}

export class DeleteQueueRequest extends jspb.Message {
  getQueueName(): string;
  setQueueName(value: string): DeleteQueueRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteQueueRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteQueueRequest): DeleteQueueRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteQueueRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteQueueRequest;
  static deserializeBinaryFromReader(message: DeleteQueueRequest, reader: jspb.BinaryReader): DeleteQueueRequest;
}

export namespace DeleteQueueRequest {
  export type AsObject = {
    queueName: string,
  }
}

export enum Time { 
  SECOND = 0,
  MINUTE = 1,
  HOUR = 2,
}
