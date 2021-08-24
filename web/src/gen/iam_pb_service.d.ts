// package: iam
// file: iam.proto

import * as iam_pb from "./iam_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type IAMServiceCreatePolicy = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.CreatePolicyRequest;
  readonly responseType: typeof iam_pb.CreatePolicyResponse;
};

type IAMServiceGetPolicy = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.GetPolicyRequest;
  readonly responseType: typeof iam_pb.GetPolicyResponse;
};

type IAMServiceUpdatePolicy = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.UpdatePolicyRequest;
  readonly responseType: typeof iam_pb.UpdatePolicyResponse;
};

type IAMServiceDeletePolicy = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.DeletePolicyRequest;
  readonly responseType: typeof iam_pb.DeletePolicyResponse;
};

type IAMServiceCreateRole = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.CreateRoleRequest;
  readonly responseType: typeof iam_pb.CreateRoleResponse;
};

type IAMServiceGetRole = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.GetRoleRequest;
  readonly responseType: typeof iam_pb.GetRoleResponse;
};

type IAMServiceUpdateRole = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.UpdateRoleRequest;
  readonly responseType: typeof iam_pb.UpdateRoleResponse;
};

type IAMServiceDeleteRole = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.DeleteRoleRequest;
  readonly responseType: typeof iam_pb.DeleteRoleResponse;
};

type IAMServiceCreateUser = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.CreateUserRequest;
  readonly responseType: typeof iam_pb.CreateUserResponse;
};

type IAMServiceGetUser = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.GetUserRequest;
  readonly responseType: typeof iam_pb.GetUserResponse;
};

type IAMServiceUpdateUser = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.UpdateUserRequest;
  readonly responseType: typeof iam_pb.UpdateUserResponse;
};

type IAMServiceDeleteUser = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.DeleteUserRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type IAMServiceCreateGroup = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.CreateGroupRequest;
  readonly responseType: typeof iam_pb.CreateGroupResponse;
};

type IAMServiceGetGroup = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.GetGroupRequest;
  readonly responseType: typeof iam_pb.GetGroupResponse;
};

type IAMServiceUpdateGroup = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.UpdateGroupRequest;
  readonly responseType: typeof iam_pb.UpdateGroupResponse;
};

type IAMServiceDeleteGroup = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.DeleteGroupRequest;
  readonly responseType: typeof iam_pb.DeleteGroupResponse;
};

type IAMServiceCreateAccessKeys = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.CreateAccessKeyRequest;
  readonly responseType: typeof iam_pb.CreateAccessKeysResponse;
};

type IAMServiceGetAccessKeys = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.GetAccessKeysRequest;
  readonly responseType: typeof iam_pb.GetAccessKeysResponse;
};

type IAMServiceDeleteAccessKeys = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.DeleteAccessKeysRequest;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type IAMServiceUserLogin = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.UserLoginRequest;
  readonly responseType: typeof iam_pb.LoginResponse;
};

type IAMServiceRootUserLogin = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.RootUserLoginRequest;
  readonly responseType: typeof iam_pb.LoginResponse;
};

type IAMServiceSignUp = {
  readonly methodName: string;
  readonly service: typeof IAMService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof iam_pb.SignUpRequest;
  readonly responseType: typeof iam_pb.SignUpResponse;
};

export class IAMService {
  static readonly serviceName: string;
  static readonly CreatePolicy: IAMServiceCreatePolicy;
  static readonly GetPolicy: IAMServiceGetPolicy;
  static readonly UpdatePolicy: IAMServiceUpdatePolicy;
  static readonly DeletePolicy: IAMServiceDeletePolicy;
  static readonly CreateRole: IAMServiceCreateRole;
  static readonly GetRole: IAMServiceGetRole;
  static readonly UpdateRole: IAMServiceUpdateRole;
  static readonly DeleteRole: IAMServiceDeleteRole;
  static readonly CreateUser: IAMServiceCreateUser;
  static readonly GetUser: IAMServiceGetUser;
  static readonly UpdateUser: IAMServiceUpdateUser;
  static readonly DeleteUser: IAMServiceDeleteUser;
  static readonly CreateGroup: IAMServiceCreateGroup;
  static readonly GetGroup: IAMServiceGetGroup;
  static readonly UpdateGroup: IAMServiceUpdateGroup;
  static readonly DeleteGroup: IAMServiceDeleteGroup;
  static readonly CreateAccessKeys: IAMServiceCreateAccessKeys;
  static readonly GetAccessKeys: IAMServiceGetAccessKeys;
  static readonly DeleteAccessKeys: IAMServiceDeleteAccessKeys;
  static readonly UserLogin: IAMServiceUserLogin;
  static readonly RootUserLogin: IAMServiceRootUserLogin;
  static readonly SignUp: IAMServiceSignUp;
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

export class IAMServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createPolicy(
    requestMessage: iam_pb.CreatePolicyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreatePolicyResponse|null) => void
  ): UnaryResponse;
  createPolicy(
    requestMessage: iam_pb.CreatePolicyRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreatePolicyResponse|null) => void
  ): UnaryResponse;
  getPolicy(
    requestMessage: iam_pb.GetPolicyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetPolicyResponse|null) => void
  ): UnaryResponse;
  getPolicy(
    requestMessage: iam_pb.GetPolicyRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetPolicyResponse|null) => void
  ): UnaryResponse;
  updatePolicy(
    requestMessage: iam_pb.UpdatePolicyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdatePolicyResponse|null) => void
  ): UnaryResponse;
  updatePolicy(
    requestMessage: iam_pb.UpdatePolicyRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdatePolicyResponse|null) => void
  ): UnaryResponse;
  deletePolicy(
    requestMessage: iam_pb.DeletePolicyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.DeletePolicyResponse|null) => void
  ): UnaryResponse;
  deletePolicy(
    requestMessage: iam_pb.DeletePolicyRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.DeletePolicyResponse|null) => void
  ): UnaryResponse;
  createRole(
    requestMessage: iam_pb.CreateRoleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateRoleResponse|null) => void
  ): UnaryResponse;
  createRole(
    requestMessage: iam_pb.CreateRoleRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateRoleResponse|null) => void
  ): UnaryResponse;
  getRole(
    requestMessage: iam_pb.GetRoleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetRoleResponse|null) => void
  ): UnaryResponse;
  getRole(
    requestMessage: iam_pb.GetRoleRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetRoleResponse|null) => void
  ): UnaryResponse;
  updateRole(
    requestMessage: iam_pb.UpdateRoleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdateRoleResponse|null) => void
  ): UnaryResponse;
  updateRole(
    requestMessage: iam_pb.UpdateRoleRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdateRoleResponse|null) => void
  ): UnaryResponse;
  deleteRole(
    requestMessage: iam_pb.DeleteRoleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.DeleteRoleResponse|null) => void
  ): UnaryResponse;
  deleteRole(
    requestMessage: iam_pb.DeleteRoleRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.DeleteRoleResponse|null) => void
  ): UnaryResponse;
  createUser(
    requestMessage: iam_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateUserResponse|null) => void
  ): UnaryResponse;
  createUser(
    requestMessage: iam_pb.CreateUserRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateUserResponse|null) => void
  ): UnaryResponse;
  getUser(
    requestMessage: iam_pb.GetUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetUserResponse|null) => void
  ): UnaryResponse;
  getUser(
    requestMessage: iam_pb.GetUserRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetUserResponse|null) => void
  ): UnaryResponse;
  updateUser(
    requestMessage: iam_pb.UpdateUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdateUserResponse|null) => void
  ): UnaryResponse;
  updateUser(
    requestMessage: iam_pb.UpdateUserRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdateUserResponse|null) => void
  ): UnaryResponse;
  deleteUser(
    requestMessage: iam_pb.DeleteUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  deleteUser(
    requestMessage: iam_pb.DeleteUserRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  createGroup(
    requestMessage: iam_pb.CreateGroupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateGroupResponse|null) => void
  ): UnaryResponse;
  createGroup(
    requestMessage: iam_pb.CreateGroupRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateGroupResponse|null) => void
  ): UnaryResponse;
  getGroup(
    requestMessage: iam_pb.GetGroupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetGroupResponse|null) => void
  ): UnaryResponse;
  getGroup(
    requestMessage: iam_pb.GetGroupRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetGroupResponse|null) => void
  ): UnaryResponse;
  updateGroup(
    requestMessage: iam_pb.UpdateGroupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdateGroupResponse|null) => void
  ): UnaryResponse;
  updateGroup(
    requestMessage: iam_pb.UpdateGroupRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.UpdateGroupResponse|null) => void
  ): UnaryResponse;
  deleteGroup(
    requestMessage: iam_pb.DeleteGroupRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.DeleteGroupResponse|null) => void
  ): UnaryResponse;
  deleteGroup(
    requestMessage: iam_pb.DeleteGroupRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.DeleteGroupResponse|null) => void
  ): UnaryResponse;
  createAccessKeys(
    requestMessage: iam_pb.CreateAccessKeyRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateAccessKeysResponse|null) => void
  ): UnaryResponse;
  createAccessKeys(
    requestMessage: iam_pb.CreateAccessKeyRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.CreateAccessKeysResponse|null) => void
  ): UnaryResponse;
  getAccessKeys(
    requestMessage: iam_pb.GetAccessKeysRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetAccessKeysResponse|null) => void
  ): UnaryResponse;
  getAccessKeys(
    requestMessage: iam_pb.GetAccessKeysRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.GetAccessKeysResponse|null) => void
  ): UnaryResponse;
  deleteAccessKeys(
    requestMessage: iam_pb.DeleteAccessKeysRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  deleteAccessKeys(
    requestMessage: iam_pb.DeleteAccessKeysRequest,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  userLogin(
    requestMessage: iam_pb.UserLoginRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.LoginResponse|null) => void
  ): UnaryResponse;
  userLogin(
    requestMessage: iam_pb.UserLoginRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.LoginResponse|null) => void
  ): UnaryResponse;
  rootUserLogin(
    requestMessage: iam_pb.RootUserLoginRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.LoginResponse|null) => void
  ): UnaryResponse;
  rootUserLogin(
    requestMessage: iam_pb.RootUserLoginRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.LoginResponse|null) => void
  ): UnaryResponse;
  signUp(
    requestMessage: iam_pb.SignUpRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: iam_pb.SignUpResponse|null) => void
  ): UnaryResponse;
  signUp(
    requestMessage: iam_pb.SignUpRequest,
    callback: (error: ServiceError|null, responseMessage: iam_pb.SignUpResponse|null) => void
  ): UnaryResponse;
}

