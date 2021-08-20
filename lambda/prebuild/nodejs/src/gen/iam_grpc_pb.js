// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var iam_pb = require('./iam_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateAccessKeyRequest(arg) {
  if (!(arg instanceof iam_pb.CreateAccessKeyRequest)) {
    throw new Error('Expected argument of type iam.CreateAccessKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateAccessKeyRequest(buffer_arg) {
  return iam_pb.CreateAccessKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateAccessKeysResponse(arg) {
  if (!(arg instanceof iam_pb.CreateAccessKeysResponse)) {
    throw new Error('Expected argument of type iam.CreateAccessKeysResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateAccessKeysResponse(buffer_arg) {
  return iam_pb.CreateAccessKeysResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateGroupRequest(arg) {
  if (!(arg instanceof iam_pb.CreateGroupRequest)) {
    throw new Error('Expected argument of type iam.CreateGroupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateGroupRequest(buffer_arg) {
  return iam_pb.CreateGroupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateGroupResponse(arg) {
  if (!(arg instanceof iam_pb.CreateGroupResponse)) {
    throw new Error('Expected argument of type iam.CreateGroupResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateGroupResponse(buffer_arg) {
  return iam_pb.CreateGroupResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreatePolicyRequest(arg) {
  if (!(arg instanceof iam_pb.CreatePolicyRequest)) {
    throw new Error('Expected argument of type iam.CreatePolicyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreatePolicyRequest(buffer_arg) {
  return iam_pb.CreatePolicyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreatePolicyResponse(arg) {
  if (!(arg instanceof iam_pb.CreatePolicyResponse)) {
    throw new Error('Expected argument of type iam.CreatePolicyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreatePolicyResponse(buffer_arg) {
  return iam_pb.CreatePolicyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateRoleRequest(arg) {
  if (!(arg instanceof iam_pb.CreateRoleRequest)) {
    throw new Error('Expected argument of type iam.CreateRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateRoleRequest(buffer_arg) {
  return iam_pb.CreateRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateRoleResponse(arg) {
  if (!(arg instanceof iam_pb.CreateRoleResponse)) {
    throw new Error('Expected argument of type iam.CreateRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateRoleResponse(buffer_arg) {
  return iam_pb.CreateRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateUserRequest(arg) {
  if (!(arg instanceof iam_pb.CreateUserRequest)) {
    throw new Error('Expected argument of type iam.CreateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateUserRequest(buffer_arg) {
  return iam_pb.CreateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_CreateUserResponse(arg) {
  if (!(arg instanceof iam_pb.CreateUserResponse)) {
    throw new Error('Expected argument of type iam.CreateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_CreateUserResponse(buffer_arg) {
  return iam_pb.CreateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeleteAccessKeysRequest(arg) {
  if (!(arg instanceof iam_pb.DeleteAccessKeysRequest)) {
    throw new Error('Expected argument of type iam.DeleteAccessKeysRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeleteAccessKeysRequest(buffer_arg) {
  return iam_pb.DeleteAccessKeysRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeleteGroupRequest(arg) {
  if (!(arg instanceof iam_pb.DeleteGroupRequest)) {
    throw new Error('Expected argument of type iam.DeleteGroupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeleteGroupRequest(buffer_arg) {
  return iam_pb.DeleteGroupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeleteGroupResponse(arg) {
  if (!(arg instanceof iam_pb.DeleteGroupResponse)) {
    throw new Error('Expected argument of type iam.DeleteGroupResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeleteGroupResponse(buffer_arg) {
  return iam_pb.DeleteGroupResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeletePolicyRequest(arg) {
  if (!(arg instanceof iam_pb.DeletePolicyRequest)) {
    throw new Error('Expected argument of type iam.DeletePolicyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeletePolicyRequest(buffer_arg) {
  return iam_pb.DeletePolicyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeletePolicyResponse(arg) {
  if (!(arg instanceof iam_pb.DeletePolicyResponse)) {
    throw new Error('Expected argument of type iam.DeletePolicyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeletePolicyResponse(buffer_arg) {
  return iam_pb.DeletePolicyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeleteRoleRequest(arg) {
  if (!(arg instanceof iam_pb.DeleteRoleRequest)) {
    throw new Error('Expected argument of type iam.DeleteRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeleteRoleRequest(buffer_arg) {
  return iam_pb.DeleteRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeleteRoleResponse(arg) {
  if (!(arg instanceof iam_pb.DeleteRoleResponse)) {
    throw new Error('Expected argument of type iam.DeleteRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeleteRoleResponse(buffer_arg) {
  return iam_pb.DeleteRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_DeleteUserRequest(arg) {
  if (!(arg instanceof iam_pb.DeleteUserRequest)) {
    throw new Error('Expected argument of type iam.DeleteUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_DeleteUserRequest(buffer_arg) {
  return iam_pb.DeleteUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetAccessKeysRequest(arg) {
  if (!(arg instanceof iam_pb.GetAccessKeysRequest)) {
    throw new Error('Expected argument of type iam.GetAccessKeysRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetAccessKeysRequest(buffer_arg) {
  return iam_pb.GetAccessKeysRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetAccessKeysResponse(arg) {
  if (!(arg instanceof iam_pb.GetAccessKeysResponse)) {
    throw new Error('Expected argument of type iam.GetAccessKeysResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetAccessKeysResponse(buffer_arg) {
  return iam_pb.GetAccessKeysResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetGroupRequest(arg) {
  if (!(arg instanceof iam_pb.GetGroupRequest)) {
    throw new Error('Expected argument of type iam.GetGroupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetGroupRequest(buffer_arg) {
  return iam_pb.GetGroupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetGroupResponse(arg) {
  if (!(arg instanceof iam_pb.GetGroupResponse)) {
    throw new Error('Expected argument of type iam.GetGroupResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetGroupResponse(buffer_arg) {
  return iam_pb.GetGroupResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetPolicyRequest(arg) {
  if (!(arg instanceof iam_pb.GetPolicyRequest)) {
    throw new Error('Expected argument of type iam.GetPolicyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetPolicyRequest(buffer_arg) {
  return iam_pb.GetPolicyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetPolicyResponse(arg) {
  if (!(arg instanceof iam_pb.GetPolicyResponse)) {
    throw new Error('Expected argument of type iam.GetPolicyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetPolicyResponse(buffer_arg) {
  return iam_pb.GetPolicyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetRoleRequest(arg) {
  if (!(arg instanceof iam_pb.GetRoleRequest)) {
    throw new Error('Expected argument of type iam.GetRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetRoleRequest(buffer_arg) {
  return iam_pb.GetRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetRoleResponse(arg) {
  if (!(arg instanceof iam_pb.GetRoleResponse)) {
    throw new Error('Expected argument of type iam.GetRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetRoleResponse(buffer_arg) {
  return iam_pb.GetRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetUserRequest(arg) {
  if (!(arg instanceof iam_pb.GetUserRequest)) {
    throw new Error('Expected argument of type iam.GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetUserRequest(buffer_arg) {
  return iam_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_GetUserResponse(arg) {
  if (!(arg instanceof iam_pb.GetUserResponse)) {
    throw new Error('Expected argument of type iam.GetUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_GetUserResponse(buffer_arg) {
  return iam_pb.GetUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_LoginResponse(arg) {
  if (!(arg instanceof iam_pb.LoginResponse)) {
    throw new Error('Expected argument of type iam.LoginResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_LoginResponse(buffer_arg) {
  return iam_pb.LoginResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_RootUserLoginRequest(arg) {
  if (!(arg instanceof iam_pb.RootUserLoginRequest)) {
    throw new Error('Expected argument of type iam.RootUserLoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_RootUserLoginRequest(buffer_arg) {
  return iam_pb.RootUserLoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_SignUpRequest(arg) {
  if (!(arg instanceof iam_pb.SignUpRequest)) {
    throw new Error('Expected argument of type iam.SignUpRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_SignUpRequest(buffer_arg) {
  return iam_pb.SignUpRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_SignUpResponse(arg) {
  if (!(arg instanceof iam_pb.SignUpResponse)) {
    throw new Error('Expected argument of type iam.SignUpResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_SignUpResponse(buffer_arg) {
  return iam_pb.SignUpResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdateGroupRequest(arg) {
  if (!(arg instanceof iam_pb.UpdateGroupRequest)) {
    throw new Error('Expected argument of type iam.UpdateGroupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdateGroupRequest(buffer_arg) {
  return iam_pb.UpdateGroupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdateGroupResponse(arg) {
  if (!(arg instanceof iam_pb.UpdateGroupResponse)) {
    throw new Error('Expected argument of type iam.UpdateGroupResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdateGroupResponse(buffer_arg) {
  return iam_pb.UpdateGroupResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdatePolicyRequest(arg) {
  if (!(arg instanceof iam_pb.UpdatePolicyRequest)) {
    throw new Error('Expected argument of type iam.UpdatePolicyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdatePolicyRequest(buffer_arg) {
  return iam_pb.UpdatePolicyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdatePolicyResponse(arg) {
  if (!(arg instanceof iam_pb.UpdatePolicyResponse)) {
    throw new Error('Expected argument of type iam.UpdatePolicyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdatePolicyResponse(buffer_arg) {
  return iam_pb.UpdatePolicyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdateRoleRequest(arg) {
  if (!(arg instanceof iam_pb.UpdateRoleRequest)) {
    throw new Error('Expected argument of type iam.UpdateRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdateRoleRequest(buffer_arg) {
  return iam_pb.UpdateRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdateRoleResponse(arg) {
  if (!(arg instanceof iam_pb.UpdateRoleResponse)) {
    throw new Error('Expected argument of type iam.UpdateRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdateRoleResponse(buffer_arg) {
  return iam_pb.UpdateRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdateUserRequest(arg) {
  if (!(arg instanceof iam_pb.UpdateUserRequest)) {
    throw new Error('Expected argument of type iam.UpdateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdateUserRequest(buffer_arg) {
  return iam_pb.UpdateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UpdateUserResponse(arg) {
  if (!(arg instanceof iam_pb.UpdateUserResponse)) {
    throw new Error('Expected argument of type iam.UpdateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UpdateUserResponse(buffer_arg) {
  return iam_pb.UpdateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_iam_UserLoginRequest(arg) {
  if (!(arg instanceof iam_pb.UserLoginRequest)) {
    throw new Error('Expected argument of type iam.UserLoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_iam_UserLoginRequest(buffer_arg) {
  return iam_pb.UserLoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var IAMServiceService = exports.IAMServiceService = {
  createPolicy: {
    path: '/iam.IAMService/CreatePolicy',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.CreatePolicyRequest,
    responseType: iam_pb.CreatePolicyResponse,
    requestSerialize: serialize_iam_CreatePolicyRequest,
    requestDeserialize: deserialize_iam_CreatePolicyRequest,
    responseSerialize: serialize_iam_CreatePolicyResponse,
    responseDeserialize: deserialize_iam_CreatePolicyResponse,
  },
  getPolicy: {
    path: '/iam.IAMService/GetPolicy',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.GetPolicyRequest,
    responseType: iam_pb.GetPolicyResponse,
    requestSerialize: serialize_iam_GetPolicyRequest,
    requestDeserialize: deserialize_iam_GetPolicyRequest,
    responseSerialize: serialize_iam_GetPolicyResponse,
    responseDeserialize: deserialize_iam_GetPolicyResponse,
  },
  updatePolicy: {
    path: '/iam.IAMService/UpdatePolicy',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.UpdatePolicyRequest,
    responseType: iam_pb.UpdatePolicyResponse,
    requestSerialize: serialize_iam_UpdatePolicyRequest,
    requestDeserialize: deserialize_iam_UpdatePolicyRequest,
    responseSerialize: serialize_iam_UpdatePolicyResponse,
    responseDeserialize: deserialize_iam_UpdatePolicyResponse,
  },
  deletePolicy: {
    path: '/iam.IAMService/DeletePolicy',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.DeletePolicyRequest,
    responseType: iam_pb.DeletePolicyResponse,
    requestSerialize: serialize_iam_DeletePolicyRequest,
    requestDeserialize: deserialize_iam_DeletePolicyRequest,
    responseSerialize: serialize_iam_DeletePolicyResponse,
    responseDeserialize: deserialize_iam_DeletePolicyResponse,
  },
  createRole: {
    path: '/iam.IAMService/CreateRole',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.CreateRoleRequest,
    responseType: iam_pb.CreateRoleResponse,
    requestSerialize: serialize_iam_CreateRoleRequest,
    requestDeserialize: deserialize_iam_CreateRoleRequest,
    responseSerialize: serialize_iam_CreateRoleResponse,
    responseDeserialize: deserialize_iam_CreateRoleResponse,
  },
  getRole: {
    path: '/iam.IAMService/GetRole',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.GetRoleRequest,
    responseType: iam_pb.GetRoleResponse,
    requestSerialize: serialize_iam_GetRoleRequest,
    requestDeserialize: deserialize_iam_GetRoleRequest,
    responseSerialize: serialize_iam_GetRoleResponse,
    responseDeserialize: deserialize_iam_GetRoleResponse,
  },
  updateRole: {
    path: '/iam.IAMService/UpdateRole',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.UpdateRoleRequest,
    responseType: iam_pb.UpdateRoleResponse,
    requestSerialize: serialize_iam_UpdateRoleRequest,
    requestDeserialize: deserialize_iam_UpdateRoleRequest,
    responseSerialize: serialize_iam_UpdateRoleResponse,
    responseDeserialize: deserialize_iam_UpdateRoleResponse,
  },
  deleteRole: {
    path: '/iam.IAMService/DeleteRole',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.DeleteRoleRequest,
    responseType: iam_pb.DeleteRoleResponse,
    requestSerialize: serialize_iam_DeleteRoleRequest,
    requestDeserialize: deserialize_iam_DeleteRoleRequest,
    responseSerialize: serialize_iam_DeleteRoleResponse,
    responseDeserialize: deserialize_iam_DeleteRoleResponse,
  },
  createUser: {
    path: '/iam.IAMService/CreateUser',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.CreateUserRequest,
    responseType: iam_pb.CreateUserResponse,
    requestSerialize: serialize_iam_CreateUserRequest,
    requestDeserialize: deserialize_iam_CreateUserRequest,
    responseSerialize: serialize_iam_CreateUserResponse,
    responseDeserialize: deserialize_iam_CreateUserResponse,
  },
  getUser: {
    path: '/iam.IAMService/GetUser',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.GetUserRequest,
    responseType: iam_pb.GetUserResponse,
    requestSerialize: serialize_iam_GetUserRequest,
    requestDeserialize: deserialize_iam_GetUserRequest,
    responseSerialize: serialize_iam_GetUserResponse,
    responseDeserialize: deserialize_iam_GetUserResponse,
  },
  updateUser: {
    path: '/iam.IAMService/UpdateUser',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.UpdateUserRequest,
    responseType: iam_pb.UpdateUserResponse,
    requestSerialize: serialize_iam_UpdateUserRequest,
    requestDeserialize: deserialize_iam_UpdateUserRequest,
    responseSerialize: serialize_iam_UpdateUserResponse,
    responseDeserialize: deserialize_iam_UpdateUserResponse,
  },
  deleteUser: {
    path: '/iam.IAMService/DeleteUser',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.DeleteUserRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_iam_DeleteUserRequest,
    requestDeserialize: deserialize_iam_DeleteUserRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  createGroup: {
    path: '/iam.IAMService/CreateGroup',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.CreateGroupRequest,
    responseType: iam_pb.CreateGroupResponse,
    requestSerialize: serialize_iam_CreateGroupRequest,
    requestDeserialize: deserialize_iam_CreateGroupRequest,
    responseSerialize: serialize_iam_CreateGroupResponse,
    responseDeserialize: deserialize_iam_CreateGroupResponse,
  },
  getGroup: {
    path: '/iam.IAMService/GetGroup',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.GetGroupRequest,
    responseType: iam_pb.GetGroupResponse,
    requestSerialize: serialize_iam_GetGroupRequest,
    requestDeserialize: deserialize_iam_GetGroupRequest,
    responseSerialize: serialize_iam_GetGroupResponse,
    responseDeserialize: deserialize_iam_GetGroupResponse,
  },
  updateGroup: {
    path: '/iam.IAMService/UpdateGroup',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.UpdateGroupRequest,
    responseType: iam_pb.UpdateGroupResponse,
    requestSerialize: serialize_iam_UpdateGroupRequest,
    requestDeserialize: deserialize_iam_UpdateGroupRequest,
    responseSerialize: serialize_iam_UpdateGroupResponse,
    responseDeserialize: deserialize_iam_UpdateGroupResponse,
  },
  deleteGroup: {
    path: '/iam.IAMService/DeleteGroup',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.DeleteGroupRequest,
    responseType: iam_pb.DeleteGroupResponse,
    requestSerialize: serialize_iam_DeleteGroupRequest,
    requestDeserialize: deserialize_iam_DeleteGroupRequest,
    responseSerialize: serialize_iam_DeleteGroupResponse,
    responseDeserialize: deserialize_iam_DeleteGroupResponse,
  },
  createAccessKeys: {
    path: '/iam.IAMService/CreateAccessKeys',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.CreateAccessKeyRequest,
    responseType: iam_pb.CreateAccessKeysResponse,
    requestSerialize: serialize_iam_CreateAccessKeyRequest,
    requestDeserialize: deserialize_iam_CreateAccessKeyRequest,
    responseSerialize: serialize_iam_CreateAccessKeysResponse,
    responseDeserialize: deserialize_iam_CreateAccessKeysResponse,
  },
  getAccessKeys: {
    path: '/iam.IAMService/GetAccessKeys',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.GetAccessKeysRequest,
    responseType: iam_pb.GetAccessKeysResponse,
    requestSerialize: serialize_iam_GetAccessKeysRequest,
    requestDeserialize: deserialize_iam_GetAccessKeysRequest,
    responseSerialize: serialize_iam_GetAccessKeysResponse,
    responseDeserialize: deserialize_iam_GetAccessKeysResponse,
  },
  deleteAccessKeys: {
    path: '/iam.IAMService/DeleteAccessKeys',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.DeleteAccessKeysRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_iam_DeleteAccessKeysRequest,
    requestDeserialize: deserialize_iam_DeleteAccessKeysRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  userLogin: {
    path: '/iam.IAMService/UserLogin',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.UserLoginRequest,
    responseType: iam_pb.LoginResponse,
    requestSerialize: serialize_iam_UserLoginRequest,
    requestDeserialize: deserialize_iam_UserLoginRequest,
    responseSerialize: serialize_iam_LoginResponse,
    responseDeserialize: deserialize_iam_LoginResponse,
  },
  rootUserLogin: {
    path: '/iam.IAMService/RootUserLogin',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.RootUserLoginRequest,
    responseType: iam_pb.LoginResponse,
    requestSerialize: serialize_iam_RootUserLoginRequest,
    requestDeserialize: deserialize_iam_RootUserLoginRequest,
    responseSerialize: serialize_iam_LoginResponse,
    responseDeserialize: deserialize_iam_LoginResponse,
  },
  signUp: {
    path: '/iam.IAMService/SignUp',
    requestStream: false,
    responseStream: false,
    requestType: iam_pb.SignUpRequest,
    responseType: iam_pb.SignUpResponse,
    requestSerialize: serialize_iam_SignUpRequest,
    requestDeserialize: deserialize_iam_SignUpRequest,
    responseSerialize: serialize_iam_SignUpResponse,
    responseDeserialize: deserialize_iam_SignUpResponse,
  },
};

exports.IAMServiceClient = grpc.makeGenericClientConstructor(IAMServiceService);
