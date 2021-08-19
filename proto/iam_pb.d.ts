import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class Policy extends jspb.Message {
  getId(): number;
  setId(value: number): Policy;

  getName(): string;
  setName(value: string): Policy;

  getManifest(): string;
  setManifest(value: string): Policy;

  getDescription(): string;
  setDescription(value: string): Policy;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Policy.AsObject;
  static toObject(includeInstance: boolean, msg: Policy): Policy.AsObject;
  static serializeBinaryToWriter(message: Policy, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Policy;
  static deserializeBinaryFromReader(message: Policy, reader: jspb.BinaryReader): Policy;
}

export namespace Policy {
  export type AsObject = {
    id: number,
    name: string,
    manifest: string,
    description: string,
  }
}

export class PolicyCreation extends jspb.Message {
  getName(): string;
  setName(value: string): PolicyCreation;

  getManifest(): string;
  setManifest(value: string): PolicyCreation;

  getDescription(): string;
  setDescription(value: string): PolicyCreation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PolicyCreation.AsObject;
  static toObject(includeInstance: boolean, msg: PolicyCreation): PolicyCreation.AsObject;
  static serializeBinaryToWriter(message: PolicyCreation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PolicyCreation;
  static deserializeBinaryFromReader(message: PolicyCreation, reader: jspb.BinaryReader): PolicyCreation;
}

export namespace PolicyCreation {
  export type AsObject = {
    name: string,
    manifest: string,
    description: string,
  }
}

export class Role extends jspb.Message {
  getId(): number;
  setId(value: number): Role;

  getName(): string;
  setName(value: string): Role;

  getDescription(): string;
  setDescription(value: string): Role;

  getPoliciesList(): Array<Policy>;
  setPoliciesList(value: Array<Policy>): Role;
  clearPoliciesList(): Role;
  addPolicies(value?: Policy, index?: number): Policy;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Role.AsObject;
  static toObject(includeInstance: boolean, msg: Role): Role.AsObject;
  static serializeBinaryToWriter(message: Role, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Role;
  static deserializeBinaryFromReader(message: Role, reader: jspb.BinaryReader): Role;
}

export namespace Role {
  export type AsObject = {
    id: number,
    name: string,
    description: string,
    policiesList: Array<Policy.AsObject>,
  }
}

export class RootUser extends jspb.Message {
  getId(): number;
  setId(value: number): RootUser;

  getEmail(): string;
  setEmail(value: string): RootUser;

  getAccountId(): string;
  setAccountId(value: string): RootUser;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RootUser.AsObject;
  static toObject(includeInstance: boolean, msg: RootUser): RootUser.AsObject;
  static serializeBinaryToWriter(message: RootUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RootUser;
  static deserializeBinaryFromReader(message: RootUser, reader: jspb.BinaryReader): RootUser;
}

export namespace RootUser {
  export type AsObject = {
    id: number,
    email: string,
    accountId: string,
  }
}

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): User;

  getName(): string;
  setName(value: string): User;

  getDescription(): string;
  setDescription(value: string): User;

  getPoliciesList(): Array<Policy>;
  setPoliciesList(value: Array<Policy>): User;
  clearPoliciesList(): User;
  addPolicies(value?: Policy, index?: number): Policy;

  getArn(): string;
  setArn(value: string): User;

  getGroupId(): number;
  setGroupId(value: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    name: string,
    description: string,
    policiesList: Array<Policy.AsObject>,
    arn: string,
    groupId: number,
  }
}

export class Group extends jspb.Message {
  getId(): number;
  setId(value: number): Group;

  getName(): string;
  setName(value: string): Group;

  getDescription(): string;
  setDescription(value: string): Group;

  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): Group;
  clearUsersList(): Group;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Group.AsObject;
  static toObject(includeInstance: boolean, msg: Group): Group.AsObject;
  static serializeBinaryToWriter(message: Group, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Group;
  static deserializeBinaryFromReader(message: Group, reader: jspb.BinaryReader): Group;
}

export namespace Group {
  export type AsObject = {
    id: number,
    name: string,
    description: string,
    usersList: Array<User.AsObject>,
  }
}

export class AccessKey extends jspb.Message {
  getAccessKeyId(): string;
  setAccessKeyId(value: string): AccessKey;

  getSecretAccessKeyId(): string;
  setSecretAccessKeyId(value: string): AccessKey;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): AccessKey;
  hasCreatedAt(): boolean;
  clearCreatedAt(): AccessKey;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccessKey.AsObject;
  static toObject(includeInstance: boolean, msg: AccessKey): AccessKey.AsObject;
  static serializeBinaryToWriter(message: AccessKey, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccessKey;
  static deserializeBinaryFromReader(message: AccessKey, reader: jspb.BinaryReader): AccessKey;
}

export namespace AccessKey {
  export type AsObject = {
    accessKeyId: string,
    secretAccessKeyId: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class UserLoginRequest extends jspb.Message {
  getAccountId(): string;
  setAccountId(value: string): UserLoginRequest;

  getName(): string;
  setName(value: string): UserLoginRequest;

  getPassword(): string;
  setPassword(value: string): UserLoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserLoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserLoginRequest): UserLoginRequest.AsObject;
  static serializeBinaryToWriter(message: UserLoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserLoginRequest;
  static deserializeBinaryFromReader(message: UserLoginRequest, reader: jspb.BinaryReader): UserLoginRequest;
}

export namespace UserLoginRequest {
  export type AsObject = {
    accountId: string,
    name: string,
    password: string,
  }
}

export class RootUserLoginRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): RootUserLoginRequest;

  getPassword(): string;
  setPassword(value: string): RootUserLoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RootUserLoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RootUserLoginRequest): RootUserLoginRequest.AsObject;
  static serializeBinaryToWriter(message: RootUserLoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RootUserLoginRequest;
  static deserializeBinaryFromReader(message: RootUserLoginRequest, reader: jspb.BinaryReader): RootUserLoginRequest;
}

export namespace RootUserLoginRequest {
  export type AsObject = {
    email: string,
    password: string,
  }
}

export class LoginResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): LoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    token: string,
  }
}

export class SignUpRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): SignUpRequest;

  getPassword(): string;
  setPassword(value: string): SignUpRequest;

  getConfirmPassword(): string;
  setConfirmPassword(value: string): SignUpRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignUpRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SignUpRequest): SignUpRequest.AsObject;
  static serializeBinaryToWriter(message: SignUpRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignUpRequest;
  static deserializeBinaryFromReader(message: SignUpRequest, reader: jspb.BinaryReader): SignUpRequest;
}

export namespace SignUpRequest {
  export type AsObject = {
    email: string,
    password: string,
    confirmPassword: string,
  }
}

export class SignUpResponse extends jspb.Message {
  getSucceed(): boolean;
  setSucceed(value: boolean): SignUpResponse;

  getAccountId(): string;
  setAccountId(value: string): SignUpResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignUpResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SignUpResponse): SignUpResponse.AsObject;
  static serializeBinaryToWriter(message: SignUpResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignUpResponse;
  static deserializeBinaryFromReader(message: SignUpResponse, reader: jspb.BinaryReader): SignUpResponse;
}

export namespace SignUpResponse {
  export type AsObject = {
    succeed: boolean,
    accountId: string,
  }
}

export class CreatePolicyRequest extends jspb.Message {
  getPolicy(): Policy | undefined;
  setPolicy(value?: Policy): CreatePolicyRequest;
  hasPolicy(): boolean;
  clearPolicy(): CreatePolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePolicyRequest): CreatePolicyRequest.AsObject;
  static serializeBinaryToWriter(message: CreatePolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePolicyRequest;
  static deserializeBinaryFromReader(message: CreatePolicyRequest, reader: jspb.BinaryReader): CreatePolicyRequest;
}

export namespace CreatePolicyRequest {
  export type AsObject = {
    policy?: Policy.AsObject,
  }
}

export class CreatePolicyResponse extends jspb.Message {
  getCreated(): Policy | undefined;
  setCreated(value?: Policy): CreatePolicyResponse;
  hasCreated(): boolean;
  clearCreated(): CreatePolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePolicyResponse): CreatePolicyResponse.AsObject;
  static serializeBinaryToWriter(message: CreatePolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePolicyResponse;
  static deserializeBinaryFromReader(message: CreatePolicyResponse, reader: jspb.BinaryReader): CreatePolicyResponse;
}

export namespace CreatePolicyResponse {
  export type AsObject = {
    created?: Policy.AsObject,
  }
}

export class GetPolicyRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetPolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPolicyRequest): GetPolicyRequest.AsObject;
  static serializeBinaryToWriter(message: GetPolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPolicyRequest;
  static deserializeBinaryFromReader(message: GetPolicyRequest, reader: jspb.BinaryReader): GetPolicyRequest;
}

export namespace GetPolicyRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetPolicyResponse extends jspb.Message {
  getPolicy(): Policy | undefined;
  setPolicy(value?: Policy): GetPolicyResponse;
  hasPolicy(): boolean;
  clearPolicy(): GetPolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPolicyResponse): GetPolicyResponse.AsObject;
  static serializeBinaryToWriter(message: GetPolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPolicyResponse;
  static deserializeBinaryFromReader(message: GetPolicyResponse, reader: jspb.BinaryReader): GetPolicyResponse;
}

export namespace GetPolicyResponse {
  export type AsObject = {
    policy?: Policy.AsObject,
  }
}

export class UpdatePolicyRequest extends jspb.Message {
  getUpdated(): Policy | undefined;
  setUpdated(value?: Policy): UpdatePolicyRequest;
  hasUpdated(): boolean;
  clearUpdated(): UpdatePolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePolicyRequest): UpdatePolicyRequest.AsObject;
  static serializeBinaryToWriter(message: UpdatePolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePolicyRequest;
  static deserializeBinaryFromReader(message: UpdatePolicyRequest, reader: jspb.BinaryReader): UpdatePolicyRequest;
}

export namespace UpdatePolicyRequest {
  export type AsObject = {
    updated?: Policy.AsObject,
  }
}

export class UpdatePolicyResponse extends jspb.Message {
  getResult(): Policy | undefined;
  setResult(value?: Policy): UpdatePolicyResponse;
  hasResult(): boolean;
  clearResult(): UpdatePolicyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePolicyResponse): UpdatePolicyResponse.AsObject;
  static serializeBinaryToWriter(message: UpdatePolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePolicyResponse;
  static deserializeBinaryFromReader(message: UpdatePolicyResponse, reader: jspb.BinaryReader): UpdatePolicyResponse;
}

export namespace UpdatePolicyResponse {
  export type AsObject = {
    result?: Policy.AsObject,
  }
}

export class DeletePolicyRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeletePolicyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePolicyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePolicyRequest): DeletePolicyRequest.AsObject;
  static serializeBinaryToWriter(message: DeletePolicyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePolicyRequest;
  static deserializeBinaryFromReader(message: DeletePolicyRequest, reader: jspb.BinaryReader): DeletePolicyRequest;
}

export namespace DeletePolicyRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeletePolicyResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePolicyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePolicyResponse): DeletePolicyResponse.AsObject;
  static serializeBinaryToWriter(message: DeletePolicyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePolicyResponse;
  static deserializeBinaryFromReader(message: DeletePolicyResponse, reader: jspb.BinaryReader): DeletePolicyResponse;
}

export namespace DeletePolicyResponse {
  export type AsObject = {
  }
}

export class CreateRoleRequest extends jspb.Message {
  getRole(): Role | undefined;
  setRole(value?: Role): CreateRoleRequest;
  hasRole(): boolean;
  clearRole(): CreateRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoleRequest): CreateRoleRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoleRequest;
  static deserializeBinaryFromReader(message: CreateRoleRequest, reader: jspb.BinaryReader): CreateRoleRequest;
}

export namespace CreateRoleRequest {
  export type AsObject = {
    role?: Role.AsObject,
  }
}

export class CreateRoleResponse extends jspb.Message {
  getCreated(): Role | undefined;
  setCreated(value?: Role): CreateRoleResponse;
  hasCreated(): boolean;
  clearCreated(): CreateRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoleResponse): CreateRoleResponse.AsObject;
  static serializeBinaryToWriter(message: CreateRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoleResponse;
  static deserializeBinaryFromReader(message: CreateRoleResponse, reader: jspb.BinaryReader): CreateRoleResponse;
}

export namespace CreateRoleResponse {
  export type AsObject = {
    created?: Role.AsObject,
  }
}

export class GetRoleRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRoleRequest): GetRoleRequest.AsObject;
  static serializeBinaryToWriter(message: GetRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRoleRequest;
  static deserializeBinaryFromReader(message: GetRoleRequest, reader: jspb.BinaryReader): GetRoleRequest;
}

export namespace GetRoleRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetRoleResponse extends jspb.Message {
  getRole(): Role | undefined;
  setRole(value?: Role): GetRoleResponse;
  hasRole(): boolean;
  clearRole(): GetRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetRoleResponse): GetRoleResponse.AsObject;
  static serializeBinaryToWriter(message: GetRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRoleResponse;
  static deserializeBinaryFromReader(message: GetRoleResponse, reader: jspb.BinaryReader): GetRoleResponse;
}

export namespace GetRoleResponse {
  export type AsObject = {
    role?: Role.AsObject,
  }
}

export class UpdateRoleRequest extends jspb.Message {
  getUpdated(): Role | undefined;
  setUpdated(value?: Role): UpdateRoleRequest;
  hasUpdated(): boolean;
  clearUpdated(): UpdateRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRoleRequest): UpdateRoleRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRoleRequest;
  static deserializeBinaryFromReader(message: UpdateRoleRequest, reader: jspb.BinaryReader): UpdateRoleRequest;
}

export namespace UpdateRoleRequest {
  export type AsObject = {
    updated?: Role.AsObject,
  }
}

export class UpdateRoleResponse extends jspb.Message {
  getResult(): Role | undefined;
  setResult(value?: Role): UpdateRoleResponse;
  hasResult(): boolean;
  clearResult(): UpdateRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRoleResponse): UpdateRoleResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRoleResponse;
  static deserializeBinaryFromReader(message: UpdateRoleResponse, reader: jspb.BinaryReader): UpdateRoleResponse;
}

export namespace UpdateRoleResponse {
  export type AsObject = {
    result?: Role.AsObject,
  }
}

export class DeleteRoleRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRoleRequest): DeleteRoleRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRoleRequest;
  static deserializeBinaryFromReader(message: DeleteRoleRequest, reader: jspb.BinaryReader): DeleteRoleRequest;
}

export namespace DeleteRoleRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteRoleResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRoleResponse): DeleteRoleResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRoleResponse;
  static deserializeBinaryFromReader(message: DeleteRoleResponse, reader: jspb.BinaryReader): DeleteRoleResponse;
}

export namespace DeleteRoleResponse {
  export type AsObject = {
  }
}

export class CreateUserRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateUserRequest;

  getDescription(): string;
  setDescription(value: string): CreateUserRequest;

  getPassword(): string;
  setPassword(value: string): CreateUserRequest;

  getPolicesList(): Array<Policy>;
  setPolicesList(value: Array<Policy>): CreateUserRequest;
  clearPolicesList(): CreateUserRequest;
  addPolices(value?: Policy, index?: number): Policy;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    name: string,
    description: string,
    password: string,
    policesList: Array<Policy.AsObject>,
  }
}

export class CreateUserResponse extends jspb.Message {
  getCreated(): User | undefined;
  setCreated(value?: User): CreateUserResponse;
  hasCreated(): boolean;
  clearCreated(): CreateUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
  static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
  static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
  export type AsObject = {
    created?: User.AsObject,
  }
}

export class GetUserRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetUserResponse extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): GetUserResponse;
  hasUser(): boolean;
  clearUser(): GetUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserResponse): GetUserResponse.AsObject;
  static serializeBinaryToWriter(message: GetUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserResponse;
  static deserializeBinaryFromReader(message: GetUserResponse, reader: jspb.BinaryReader): GetUserResponse;
}

export namespace GetUserResponse {
  export type AsObject = {
    user?: User.AsObject,
  }
}

export class UpdateUserRequest extends jspb.Message {
  getUpdated(): User | undefined;
  setUpdated(value?: User): UpdateUserRequest;
  hasUpdated(): boolean;
  clearUpdated(): UpdateUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
  static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    updated?: User.AsObject,
  }
}

export class UpdateUserResponse extends jspb.Message {
  getResult(): User | undefined;
  setResult(value?: User): UpdateUserResponse;
  hasResult(): boolean;
  clearResult(): UpdateUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserResponse): UpdateUserResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserResponse;
  static deserializeBinaryFromReader(message: UpdateUserResponse, reader: jspb.BinaryReader): UpdateUserResponse;
}

export namespace UpdateUserResponse {
  export type AsObject = {
    result?: User.AsObject,
  }
}

export class DeleteUserRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
  static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
  export type AsObject = {
    id: number,
  }
}

export class CreateGroupRequest extends jspb.Message {
  getGroup(): Group | undefined;
  setGroup(value?: Group): CreateGroupRequest;
  hasGroup(): boolean;
  clearGroup(): CreateGroupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateGroupRequest): CreateGroupRequest.AsObject;
  static serializeBinaryToWriter(message: CreateGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateGroupRequest;
  static deserializeBinaryFromReader(message: CreateGroupRequest, reader: jspb.BinaryReader): CreateGroupRequest;
}

export namespace CreateGroupRequest {
  export type AsObject = {
    group?: Group.AsObject,
  }
}

export class CreateGroupResponse extends jspb.Message {
  getCreated(): Group | undefined;
  setCreated(value?: Group): CreateGroupResponse;
  hasCreated(): boolean;
  clearCreated(): CreateGroupResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateGroupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateGroupResponse): CreateGroupResponse.AsObject;
  static serializeBinaryToWriter(message: CreateGroupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateGroupResponse;
  static deserializeBinaryFromReader(message: CreateGroupResponse, reader: jspb.BinaryReader): CreateGroupResponse;
}

export namespace CreateGroupResponse {
  export type AsObject = {
    created?: Group.AsObject,
  }
}

export class GetGroupRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetGroupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetGroupRequest): GetGroupRequest.AsObject;
  static serializeBinaryToWriter(message: GetGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetGroupRequest;
  static deserializeBinaryFromReader(message: GetGroupRequest, reader: jspb.BinaryReader): GetGroupRequest;
}

export namespace GetGroupRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetGroupResponse extends jspb.Message {
  getGroup(): Group | undefined;
  setGroup(value?: Group): GetGroupResponse;
  hasGroup(): boolean;
  clearGroup(): GetGroupResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetGroupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetGroupResponse): GetGroupResponse.AsObject;
  static serializeBinaryToWriter(message: GetGroupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetGroupResponse;
  static deserializeBinaryFromReader(message: GetGroupResponse, reader: jspb.BinaryReader): GetGroupResponse;
}

export namespace GetGroupResponse {
  export type AsObject = {
    group?: Group.AsObject,
  }
}

export class UpdateGroupRequest extends jspb.Message {
  getUpdated(): Group | undefined;
  setUpdated(value?: Group): UpdateGroupRequest;
  hasUpdated(): boolean;
  clearUpdated(): UpdateGroupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupRequest): UpdateGroupRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupRequest;
  static deserializeBinaryFromReader(message: UpdateGroupRequest, reader: jspb.BinaryReader): UpdateGroupRequest;
}

export namespace UpdateGroupRequest {
  export type AsObject = {
    updated?: Group.AsObject,
  }
}

export class UpdateGroupResponse extends jspb.Message {
  getResult(): Group | undefined;
  setResult(value?: Group): UpdateGroupResponse;
  hasResult(): boolean;
  clearResult(): UpdateGroupResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupResponse): UpdateGroupResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateGroupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupResponse;
  static deserializeBinaryFromReader(message: UpdateGroupResponse, reader: jspb.BinaryReader): UpdateGroupResponse;
}

export namespace UpdateGroupResponse {
  export type AsObject = {
    result?: Group.AsObject,
  }
}

export class DeleteGroupRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteGroupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteGroupRequest): DeleteGroupRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteGroupRequest;
  static deserializeBinaryFromReader(message: DeleteGroupRequest, reader: jspb.BinaryReader): DeleteGroupRequest;
}

export namespace DeleteGroupRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteGroupResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteGroupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteGroupResponse): DeleteGroupResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteGroupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteGroupResponse;
  static deserializeBinaryFromReader(message: DeleteGroupResponse, reader: jspb.BinaryReader): DeleteGroupResponse;
}

export namespace DeleteGroupResponse {
  export type AsObject = {
  }
}

export class CreateAccessKeyRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): CreateAccessKeyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAccessKeyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAccessKeyRequest): CreateAccessKeyRequest.AsObject;
  static serializeBinaryToWriter(message: CreateAccessKeyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAccessKeyRequest;
  static deserializeBinaryFromReader(message: CreateAccessKeyRequest, reader: jspb.BinaryReader): CreateAccessKeyRequest;
}

export namespace CreateAccessKeyRequest {
  export type AsObject = {
    userId: number,
  }
}

export class CreateAccessKeysResponse extends jspb.Message {
  getCreated(): AccessKey | undefined;
  setCreated(value?: AccessKey): CreateAccessKeysResponse;
  hasCreated(): boolean;
  clearCreated(): CreateAccessKeysResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAccessKeysResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAccessKeysResponse): CreateAccessKeysResponse.AsObject;
  static serializeBinaryToWriter(message: CreateAccessKeysResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAccessKeysResponse;
  static deserializeBinaryFromReader(message: CreateAccessKeysResponse, reader: jspb.BinaryReader): CreateAccessKeysResponse;
}

export namespace CreateAccessKeysResponse {
  export type AsObject = {
    created?: AccessKey.AsObject,
  }
}

export class GetAccessKeysRequest extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): GetAccessKeysRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAccessKeysRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAccessKeysRequest): GetAccessKeysRequest.AsObject;
  static serializeBinaryToWriter(message: GetAccessKeysRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAccessKeysRequest;
  static deserializeBinaryFromReader(message: GetAccessKeysRequest, reader: jspb.BinaryReader): GetAccessKeysRequest;
}

export namespace GetAccessKeysRequest {
  export type AsObject = {
    userId: number,
  }
}

export class GetAccessKeysResponse extends jspb.Message {
  getAccesskeys(): AccessKey | undefined;
  setAccesskeys(value?: AccessKey): GetAccessKeysResponse;
  hasAccesskeys(): boolean;
  clearAccesskeys(): GetAccessKeysResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAccessKeysResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAccessKeysResponse): GetAccessKeysResponse.AsObject;
  static serializeBinaryToWriter(message: GetAccessKeysResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAccessKeysResponse;
  static deserializeBinaryFromReader(message: GetAccessKeysResponse, reader: jspb.BinaryReader): GetAccessKeysResponse;
}

export namespace GetAccessKeysResponse {
  export type AsObject = {
    accesskeys?: AccessKey.AsObject,
  }
}

export class UpdateAccessKeysRequest extends jspb.Message {
  getUpdated(): AccessKey | undefined;
  setUpdated(value?: AccessKey): UpdateAccessKeysRequest;
  hasUpdated(): boolean;
  clearUpdated(): UpdateAccessKeysRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAccessKeysRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAccessKeysRequest): UpdateAccessKeysRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateAccessKeysRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAccessKeysRequest;
  static deserializeBinaryFromReader(message: UpdateAccessKeysRequest, reader: jspb.BinaryReader): UpdateAccessKeysRequest;
}

export namespace UpdateAccessKeysRequest {
  export type AsObject = {
    updated?: AccessKey.AsObject,
  }
}

export class UpdateAccessKeysResponse extends jspb.Message {
  getResult(): AccessKey | undefined;
  setResult(value?: AccessKey): UpdateAccessKeysResponse;
  hasResult(): boolean;
  clearResult(): UpdateAccessKeysResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateAccessKeysResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateAccessKeysResponse): UpdateAccessKeysResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateAccessKeysResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateAccessKeysResponse;
  static deserializeBinaryFromReader(message: UpdateAccessKeysResponse, reader: jspb.BinaryReader): UpdateAccessKeysResponse;
}

export namespace UpdateAccessKeysResponse {
  export type AsObject = {
    result?: AccessKey.AsObject,
  }
}

export class DeleteAccessKeysRequest extends jspb.Message {
  getAccessKeyId(): string;
  setAccessKeyId(value: string): DeleteAccessKeysRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteAccessKeysRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteAccessKeysRequest): DeleteAccessKeysRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteAccessKeysRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteAccessKeysRequest;
  static deserializeBinaryFromReader(message: DeleteAccessKeysRequest, reader: jspb.BinaryReader): DeleteAccessKeysRequest;
}

export namespace DeleteAccessKeysRequest {
  export type AsObject = {
    accessKeyId: string,
  }
}

