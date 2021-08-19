/**
 * @fileoverview gRPC-Web generated client stub for iam
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as iam_pb from './iam_pb';


export class IAMServiceClient {
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

  methodInfoCreatePolicy = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.CreatePolicyResponse,
    (request: iam_pb.CreatePolicyRequest) => {
      return request.serializeBinary();
    },
    iam_pb.CreatePolicyResponse.deserializeBinary
  );

  createPolicy(
    request: iam_pb.CreatePolicyRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.CreatePolicyResponse>;

  createPolicy(
    request: iam_pb.CreatePolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.CreatePolicyResponse) => void): grpcWeb.ClientReadableStream<iam_pb.CreatePolicyResponse>;

  createPolicy(
    request: iam_pb.CreatePolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.CreatePolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/CreatePolicy',
        request,
        metadata || {},
        this.methodInfoCreatePolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/CreatePolicy',
    request,
    metadata || {},
    this.methodInfoCreatePolicy);
  }

  methodInfoGetPolicy = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.GetPolicyResponse,
    (request: iam_pb.GetPolicyRequest) => {
      return request.serializeBinary();
    },
    iam_pb.GetPolicyResponse.deserializeBinary
  );

  getPolicy(
    request: iam_pb.GetPolicyRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.GetPolicyResponse>;

  getPolicy(
    request: iam_pb.GetPolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.GetPolicyResponse) => void): grpcWeb.ClientReadableStream<iam_pb.GetPolicyResponse>;

  getPolicy(
    request: iam_pb.GetPolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.GetPolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/GetPolicy',
        request,
        metadata || {},
        this.methodInfoGetPolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/GetPolicy',
    request,
    metadata || {},
    this.methodInfoGetPolicy);
  }

  methodInfoUpdatePolicy = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.UpdatePolicyResponse,
    (request: iam_pb.UpdatePolicyRequest) => {
      return request.serializeBinary();
    },
    iam_pb.UpdatePolicyResponse.deserializeBinary
  );

  updatePolicy(
    request: iam_pb.UpdatePolicyRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.UpdatePolicyResponse>;

  updatePolicy(
    request: iam_pb.UpdatePolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.UpdatePolicyResponse) => void): grpcWeb.ClientReadableStream<iam_pb.UpdatePolicyResponse>;

  updatePolicy(
    request: iam_pb.UpdatePolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.UpdatePolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/UpdatePolicy',
        request,
        metadata || {},
        this.methodInfoUpdatePolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/UpdatePolicy',
    request,
    metadata || {},
    this.methodInfoUpdatePolicy);
  }

  methodInfoDeletePolicy = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.DeletePolicyResponse,
    (request: iam_pb.DeletePolicyRequest) => {
      return request.serializeBinary();
    },
    iam_pb.DeletePolicyResponse.deserializeBinary
  );

  deletePolicy(
    request: iam_pb.DeletePolicyRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.DeletePolicyResponse>;

  deletePolicy(
    request: iam_pb.DeletePolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.DeletePolicyResponse) => void): grpcWeb.ClientReadableStream<iam_pb.DeletePolicyResponse>;

  deletePolicy(
    request: iam_pb.DeletePolicyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.DeletePolicyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/DeletePolicy',
        request,
        metadata || {},
        this.methodInfoDeletePolicy,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/DeletePolicy',
    request,
    metadata || {},
    this.methodInfoDeletePolicy);
  }

  methodInfoCreateRole = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.CreateRoleResponse,
    (request: iam_pb.CreateRoleRequest) => {
      return request.serializeBinary();
    },
    iam_pb.CreateRoleResponse.deserializeBinary
  );

  createRole(
    request: iam_pb.CreateRoleRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.CreateRoleResponse>;

  createRole(
    request: iam_pb.CreateRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.CreateRoleResponse) => void): grpcWeb.ClientReadableStream<iam_pb.CreateRoleResponse>;

  createRole(
    request: iam_pb.CreateRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.CreateRoleResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/CreateRole',
        request,
        metadata || {},
        this.methodInfoCreateRole,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/CreateRole',
    request,
    metadata || {},
    this.methodInfoCreateRole);
  }

  methodInfoGetRole = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.GetRoleResponse,
    (request: iam_pb.GetRoleRequest) => {
      return request.serializeBinary();
    },
    iam_pb.GetRoleResponse.deserializeBinary
  );

  getRole(
    request: iam_pb.GetRoleRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.GetRoleResponse>;

  getRole(
    request: iam_pb.GetRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.GetRoleResponse) => void): grpcWeb.ClientReadableStream<iam_pb.GetRoleResponse>;

  getRole(
    request: iam_pb.GetRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.GetRoleResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/GetRole',
        request,
        metadata || {},
        this.methodInfoGetRole,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/GetRole',
    request,
    metadata || {},
    this.methodInfoGetRole);
  }

  methodInfoUpdateRole = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.UpdateRoleResponse,
    (request: iam_pb.UpdateRoleRequest) => {
      return request.serializeBinary();
    },
    iam_pb.UpdateRoleResponse.deserializeBinary
  );

  updateRole(
    request: iam_pb.UpdateRoleRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.UpdateRoleResponse>;

  updateRole(
    request: iam_pb.UpdateRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.UpdateRoleResponse) => void): grpcWeb.ClientReadableStream<iam_pb.UpdateRoleResponse>;

  updateRole(
    request: iam_pb.UpdateRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.UpdateRoleResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/UpdateRole',
        request,
        metadata || {},
        this.methodInfoUpdateRole,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/UpdateRole',
    request,
    metadata || {},
    this.methodInfoUpdateRole);
  }

  methodInfoDeleteRole = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.DeleteRoleResponse,
    (request: iam_pb.DeleteRoleRequest) => {
      return request.serializeBinary();
    },
    iam_pb.DeleteRoleResponse.deserializeBinary
  );

  deleteRole(
    request: iam_pb.DeleteRoleRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.DeleteRoleResponse>;

  deleteRole(
    request: iam_pb.DeleteRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.DeleteRoleResponse) => void): grpcWeb.ClientReadableStream<iam_pb.DeleteRoleResponse>;

  deleteRole(
    request: iam_pb.DeleteRoleRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.DeleteRoleResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/DeleteRole',
        request,
        metadata || {},
        this.methodInfoDeleteRole,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/DeleteRole',
    request,
    metadata || {},
    this.methodInfoDeleteRole);
  }

  methodInfoCreateUser = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.CreateUserResponse,
    (request: iam_pb.CreateUserRequest) => {
      return request.serializeBinary();
    },
    iam_pb.CreateUserResponse.deserializeBinary
  );

  createUser(
    request: iam_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.CreateUserResponse>;

  createUser(
    request: iam_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.CreateUserResponse) => void): grpcWeb.ClientReadableStream<iam_pb.CreateUserResponse>;

  createUser(
    request: iam_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.CreateUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/CreateUser',
        request,
        metadata || {},
        this.methodInfoCreateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/CreateUser',
    request,
    metadata || {},
    this.methodInfoCreateUser);
  }

  methodInfoGetUser = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.GetUserResponse,
    (request: iam_pb.GetUserRequest) => {
      return request.serializeBinary();
    },
    iam_pb.GetUserResponse.deserializeBinary
  );

  getUser(
    request: iam_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.GetUserResponse>;

  getUser(
    request: iam_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.GetUserResponse) => void): grpcWeb.ClientReadableStream<iam_pb.GetUserResponse>;

  getUser(
    request: iam_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.GetUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/GetUser',
        request,
        metadata || {},
        this.methodInfoGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/GetUser',
    request,
    metadata || {},
    this.methodInfoGetUser);
  }

  methodInfoUpdateUser = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.UpdateUserResponse,
    (request: iam_pb.UpdateUserRequest) => {
      return request.serializeBinary();
    },
    iam_pb.UpdateUserResponse.deserializeBinary
  );

  updateUser(
    request: iam_pb.UpdateUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.UpdateUserResponse>;

  updateUser(
    request: iam_pb.UpdateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.UpdateUserResponse) => void): grpcWeb.ClientReadableStream<iam_pb.UpdateUserResponse>;

  updateUser(
    request: iam_pb.UpdateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.UpdateUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/UpdateUser',
        request,
        metadata || {},
        this.methodInfoUpdateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/UpdateUser',
    request,
    metadata || {},
    this.methodInfoUpdateUser);
  }

  methodInfoDeleteUser = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: iam_pb.DeleteUserRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteUser(
    request: iam_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteUser(
    request: iam_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteUser(
    request: iam_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/DeleteUser',
        request,
        metadata || {},
        this.methodInfoDeleteUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/DeleteUser',
    request,
    metadata || {},
    this.methodInfoDeleteUser);
  }

  methodInfoCreateGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.CreateGroupResponse,
    (request: iam_pb.CreateGroupRequest) => {
      return request.serializeBinary();
    },
    iam_pb.CreateGroupResponse.deserializeBinary
  );

  createGroup(
    request: iam_pb.CreateGroupRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.CreateGroupResponse>;

  createGroup(
    request: iam_pb.CreateGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.CreateGroupResponse) => void): grpcWeb.ClientReadableStream<iam_pb.CreateGroupResponse>;

  createGroup(
    request: iam_pb.CreateGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.CreateGroupResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/CreateGroup',
        request,
        metadata || {},
        this.methodInfoCreateGroup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/CreateGroup',
    request,
    metadata || {},
    this.methodInfoCreateGroup);
  }

  methodInfoGetGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.GetGroupResponse,
    (request: iam_pb.GetGroupRequest) => {
      return request.serializeBinary();
    },
    iam_pb.GetGroupResponse.deserializeBinary
  );

  getGroup(
    request: iam_pb.GetGroupRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.GetGroupResponse>;

  getGroup(
    request: iam_pb.GetGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.GetGroupResponse) => void): grpcWeb.ClientReadableStream<iam_pb.GetGroupResponse>;

  getGroup(
    request: iam_pb.GetGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.GetGroupResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/GetGroup',
        request,
        metadata || {},
        this.methodInfoGetGroup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/GetGroup',
    request,
    metadata || {},
    this.methodInfoGetGroup);
  }

  methodInfoUpdateGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.UpdateGroupResponse,
    (request: iam_pb.UpdateGroupRequest) => {
      return request.serializeBinary();
    },
    iam_pb.UpdateGroupResponse.deserializeBinary
  );

  updateGroup(
    request: iam_pb.UpdateGroupRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.UpdateGroupResponse>;

  updateGroup(
    request: iam_pb.UpdateGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.UpdateGroupResponse) => void): grpcWeb.ClientReadableStream<iam_pb.UpdateGroupResponse>;

  updateGroup(
    request: iam_pb.UpdateGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.UpdateGroupResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/UpdateGroup',
        request,
        metadata || {},
        this.methodInfoUpdateGroup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/UpdateGroup',
    request,
    metadata || {},
    this.methodInfoUpdateGroup);
  }

  methodInfoDeleteGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.DeleteGroupResponse,
    (request: iam_pb.DeleteGroupRequest) => {
      return request.serializeBinary();
    },
    iam_pb.DeleteGroupResponse.deserializeBinary
  );

  deleteGroup(
    request: iam_pb.DeleteGroupRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.DeleteGroupResponse>;

  deleteGroup(
    request: iam_pb.DeleteGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.DeleteGroupResponse) => void): grpcWeb.ClientReadableStream<iam_pb.DeleteGroupResponse>;

  deleteGroup(
    request: iam_pb.DeleteGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.DeleteGroupResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/DeleteGroup',
        request,
        metadata || {},
        this.methodInfoDeleteGroup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/DeleteGroup',
    request,
    metadata || {},
    this.methodInfoDeleteGroup);
  }

  methodInfoCreateAccessKeys = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.CreateAccessKeysResponse,
    (request: iam_pb.CreateAccessKeyRequest) => {
      return request.serializeBinary();
    },
    iam_pb.CreateAccessKeysResponse.deserializeBinary
  );

  createAccessKeys(
    request: iam_pb.CreateAccessKeyRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.CreateAccessKeysResponse>;

  createAccessKeys(
    request: iam_pb.CreateAccessKeyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.CreateAccessKeysResponse) => void): grpcWeb.ClientReadableStream<iam_pb.CreateAccessKeysResponse>;

  createAccessKeys(
    request: iam_pb.CreateAccessKeyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.CreateAccessKeysResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/CreateAccessKeys',
        request,
        metadata || {},
        this.methodInfoCreateAccessKeys,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/CreateAccessKeys',
    request,
    metadata || {},
    this.methodInfoCreateAccessKeys);
  }

  methodInfoGetAccessKeys = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.GetAccessKeysResponse,
    (request: iam_pb.GetAccessKeysRequest) => {
      return request.serializeBinary();
    },
    iam_pb.GetAccessKeysResponse.deserializeBinary
  );

  getAccessKeys(
    request: iam_pb.GetAccessKeysRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.GetAccessKeysResponse>;

  getAccessKeys(
    request: iam_pb.GetAccessKeysRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.GetAccessKeysResponse) => void): grpcWeb.ClientReadableStream<iam_pb.GetAccessKeysResponse>;

  getAccessKeys(
    request: iam_pb.GetAccessKeysRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.GetAccessKeysResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/GetAccessKeys',
        request,
        metadata || {},
        this.methodInfoGetAccessKeys,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/GetAccessKeys',
    request,
    metadata || {},
    this.methodInfoGetAccessKeys);
  }

  methodInfoDeleteAccessKeys = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: iam_pb.DeleteAccessKeysRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteAccessKeys(
    request: iam_pb.DeleteAccessKeysRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteAccessKeys(
    request: iam_pb.DeleteAccessKeysRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteAccessKeys(
    request: iam_pb.DeleteAccessKeysRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/DeleteAccessKeys',
        request,
        metadata || {},
        this.methodInfoDeleteAccessKeys,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/DeleteAccessKeys',
    request,
    metadata || {},
    this.methodInfoDeleteAccessKeys);
  }

  methodInfoUserLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.LoginResponse,
    (request: iam_pb.UserLoginRequest) => {
      return request.serializeBinary();
    },
    iam_pb.LoginResponse.deserializeBinary
  );

  userLogin(
    request: iam_pb.UserLoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.LoginResponse>;

  userLogin(
    request: iam_pb.UserLoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<iam_pb.LoginResponse>;

  userLogin(
    request: iam_pb.UserLoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/UserLogin',
        request,
        metadata || {},
        this.methodInfoUserLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/UserLogin',
    request,
    metadata || {},
    this.methodInfoUserLogin);
  }

  methodInfoRootUserLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.LoginResponse,
    (request: iam_pb.RootUserLoginRequest) => {
      return request.serializeBinary();
    },
    iam_pb.LoginResponse.deserializeBinary
  );

  rootUserLogin(
    request: iam_pb.RootUserLoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.LoginResponse>;

  rootUserLogin(
    request: iam_pb.RootUserLoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<iam_pb.LoginResponse>;

  rootUserLogin(
    request: iam_pb.RootUserLoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/RootUserLogin',
        request,
        metadata || {},
        this.methodInfoRootUserLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/RootUserLogin',
    request,
    metadata || {},
    this.methodInfoRootUserLogin);
  }

  methodInfoSignUp = new grpcWeb.AbstractClientBase.MethodInfo(
    iam_pb.SignUpResponse,
    (request: iam_pb.SignUpRequest) => {
      return request.serializeBinary();
    },
    iam_pb.SignUpResponse.deserializeBinary
  );

  signUp(
    request: iam_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | null): Promise<iam_pb.SignUpResponse>;

  signUp(
    request: iam_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: iam_pb.SignUpResponse) => void): grpcWeb.ClientReadableStream<iam_pb.SignUpResponse>;

  signUp(
    request: iam_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: iam_pb.SignUpResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/iam.IAMService/SignUp',
        request,
        metadata || {},
        this.methodInfoSignUp,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/iam.IAMService/SignUp',
    request,
    metadata || {},
    this.methodInfoSignUp);
  }

}

