// package: iam
// file: iam.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as iam_pb from "./iam_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface IIAMServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createPolicy: IIAMServiceService_ICreatePolicy;
    getPolicy: IIAMServiceService_IGetPolicy;
    updatePolicy: IIAMServiceService_IUpdatePolicy;
    deletePolicy: IIAMServiceService_IDeletePolicy;
    createRole: IIAMServiceService_ICreateRole;
    getRole: IIAMServiceService_IGetRole;
    updateRole: IIAMServiceService_IUpdateRole;
    deleteRole: IIAMServiceService_IDeleteRole;
    createUser: IIAMServiceService_ICreateUser;
    getUser: IIAMServiceService_IGetUser;
    updateUser: IIAMServiceService_IUpdateUser;
    deleteUser: IIAMServiceService_IDeleteUser;
    createGroup: IIAMServiceService_ICreateGroup;
    getGroup: IIAMServiceService_IGetGroup;
    updateGroup: IIAMServiceService_IUpdateGroup;
    deleteGroup: IIAMServiceService_IDeleteGroup;
    createAccessKeys: IIAMServiceService_ICreateAccessKeys;
    getAccessKeys: IIAMServiceService_IGetAccessKeys;
    deleteAccessKeys: IIAMServiceService_IDeleteAccessKeys;
    userLogin: IIAMServiceService_IUserLogin;
    rootUserLogin: IIAMServiceService_IRootUserLogin;
    signUp: IIAMServiceService_ISignUp;
}

interface IIAMServiceService_ICreatePolicy extends grpc.MethodDefinition<iam_pb.CreatePolicyRequest, iam_pb.CreatePolicyResponse> {
    path: "/iam.IAMService/CreatePolicy";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.CreatePolicyRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.CreatePolicyRequest>;
    responseSerialize: grpc.serialize<iam_pb.CreatePolicyResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.CreatePolicyResponse>;
}
interface IIAMServiceService_IGetPolicy extends grpc.MethodDefinition<iam_pb.GetPolicyRequest, iam_pb.GetPolicyResponse> {
    path: "/iam.IAMService/GetPolicy";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.GetPolicyRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.GetPolicyRequest>;
    responseSerialize: grpc.serialize<iam_pb.GetPolicyResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.GetPolicyResponse>;
}
interface IIAMServiceService_IUpdatePolicy extends grpc.MethodDefinition<iam_pb.UpdatePolicyRequest, iam_pb.UpdatePolicyResponse> {
    path: "/iam.IAMService/UpdatePolicy";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.UpdatePolicyRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.UpdatePolicyRequest>;
    responseSerialize: grpc.serialize<iam_pb.UpdatePolicyResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.UpdatePolicyResponse>;
}
interface IIAMServiceService_IDeletePolicy extends grpc.MethodDefinition<iam_pb.DeletePolicyRequest, iam_pb.DeletePolicyResponse> {
    path: "/iam.IAMService/DeletePolicy";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.DeletePolicyRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.DeletePolicyRequest>;
    responseSerialize: grpc.serialize<iam_pb.DeletePolicyResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.DeletePolicyResponse>;
}
interface IIAMServiceService_ICreateRole extends grpc.MethodDefinition<iam_pb.CreateRoleRequest, iam_pb.CreateRoleResponse> {
    path: "/iam.IAMService/CreateRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.CreateRoleRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.CreateRoleRequest>;
    responseSerialize: grpc.serialize<iam_pb.CreateRoleResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.CreateRoleResponse>;
}
interface IIAMServiceService_IGetRole extends grpc.MethodDefinition<iam_pb.GetRoleRequest, iam_pb.GetRoleResponse> {
    path: "/iam.IAMService/GetRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.GetRoleRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.GetRoleRequest>;
    responseSerialize: grpc.serialize<iam_pb.GetRoleResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.GetRoleResponse>;
}
interface IIAMServiceService_IUpdateRole extends grpc.MethodDefinition<iam_pb.UpdateRoleRequest, iam_pb.UpdateRoleResponse> {
    path: "/iam.IAMService/UpdateRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.UpdateRoleRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.UpdateRoleRequest>;
    responseSerialize: grpc.serialize<iam_pb.UpdateRoleResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.UpdateRoleResponse>;
}
interface IIAMServiceService_IDeleteRole extends grpc.MethodDefinition<iam_pb.DeleteRoleRequest, iam_pb.DeleteRoleResponse> {
    path: "/iam.IAMService/DeleteRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.DeleteRoleRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.DeleteRoleRequest>;
    responseSerialize: grpc.serialize<iam_pb.DeleteRoleResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.DeleteRoleResponse>;
}
interface IIAMServiceService_ICreateUser extends grpc.MethodDefinition<iam_pb.CreateUserRequest, iam_pb.CreateUserResponse> {
    path: "/iam.IAMService/CreateUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.CreateUserRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.CreateUserRequest>;
    responseSerialize: grpc.serialize<iam_pb.CreateUserResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.CreateUserResponse>;
}
interface IIAMServiceService_IGetUser extends grpc.MethodDefinition<iam_pb.GetUserRequest, iam_pb.GetUserResponse> {
    path: "/iam.IAMService/GetUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.GetUserRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.GetUserRequest>;
    responseSerialize: grpc.serialize<iam_pb.GetUserResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.GetUserResponse>;
}
interface IIAMServiceService_IUpdateUser extends grpc.MethodDefinition<iam_pb.UpdateUserRequest, iam_pb.UpdateUserResponse> {
    path: "/iam.IAMService/UpdateUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.UpdateUserRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.UpdateUserRequest>;
    responseSerialize: grpc.serialize<iam_pb.UpdateUserResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.UpdateUserResponse>;
}
interface IIAMServiceService_IDeleteUser extends grpc.MethodDefinition<iam_pb.DeleteUserRequest, google_protobuf_empty_pb.Empty> {
    path: "/iam.IAMService/DeleteUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.DeleteUserRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.DeleteUserRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface IIAMServiceService_ICreateGroup extends grpc.MethodDefinition<iam_pb.CreateGroupRequest, iam_pb.CreateGroupResponse> {
    path: "/iam.IAMService/CreateGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.CreateGroupRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.CreateGroupRequest>;
    responseSerialize: grpc.serialize<iam_pb.CreateGroupResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.CreateGroupResponse>;
}
interface IIAMServiceService_IGetGroup extends grpc.MethodDefinition<iam_pb.GetGroupRequest, iam_pb.GetGroupResponse> {
    path: "/iam.IAMService/GetGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.GetGroupRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.GetGroupRequest>;
    responseSerialize: grpc.serialize<iam_pb.GetGroupResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.GetGroupResponse>;
}
interface IIAMServiceService_IUpdateGroup extends grpc.MethodDefinition<iam_pb.UpdateGroupRequest, iam_pb.UpdateGroupResponse> {
    path: "/iam.IAMService/UpdateGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.UpdateGroupRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.UpdateGroupRequest>;
    responseSerialize: grpc.serialize<iam_pb.UpdateGroupResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.UpdateGroupResponse>;
}
interface IIAMServiceService_IDeleteGroup extends grpc.MethodDefinition<iam_pb.DeleteGroupRequest, iam_pb.DeleteGroupResponse> {
    path: "/iam.IAMService/DeleteGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.DeleteGroupRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.DeleteGroupRequest>;
    responseSerialize: grpc.serialize<iam_pb.DeleteGroupResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.DeleteGroupResponse>;
}
interface IIAMServiceService_ICreateAccessKeys extends grpc.MethodDefinition<iam_pb.CreateAccessKeyRequest, iam_pb.CreateAccessKeysResponse> {
    path: "/iam.IAMService/CreateAccessKeys";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.CreateAccessKeyRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.CreateAccessKeyRequest>;
    responseSerialize: grpc.serialize<iam_pb.CreateAccessKeysResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.CreateAccessKeysResponse>;
}
interface IIAMServiceService_IGetAccessKeys extends grpc.MethodDefinition<iam_pb.GetAccessKeysRequest, iam_pb.GetAccessKeysResponse> {
    path: "/iam.IAMService/GetAccessKeys";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.GetAccessKeysRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.GetAccessKeysRequest>;
    responseSerialize: grpc.serialize<iam_pb.GetAccessKeysResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.GetAccessKeysResponse>;
}
interface IIAMServiceService_IDeleteAccessKeys extends grpc.MethodDefinition<iam_pb.DeleteAccessKeysRequest, google_protobuf_empty_pb.Empty> {
    path: "/iam.IAMService/DeleteAccessKeys";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.DeleteAccessKeysRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.DeleteAccessKeysRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface IIAMServiceService_IUserLogin extends grpc.MethodDefinition<iam_pb.UserLoginRequest, iam_pb.LoginResponse> {
    path: "/iam.IAMService/UserLogin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.UserLoginRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.UserLoginRequest>;
    responseSerialize: grpc.serialize<iam_pb.LoginResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.LoginResponse>;
}
interface IIAMServiceService_IRootUserLogin extends grpc.MethodDefinition<iam_pb.RootUserLoginRequest, iam_pb.LoginResponse> {
    path: "/iam.IAMService/RootUserLogin";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.RootUserLoginRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.RootUserLoginRequest>;
    responseSerialize: grpc.serialize<iam_pb.LoginResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.LoginResponse>;
}
interface IIAMServiceService_ISignUp extends grpc.MethodDefinition<iam_pb.SignUpRequest, iam_pb.SignUpResponse> {
    path: "/iam.IAMService/SignUp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<iam_pb.SignUpRequest>;
    requestDeserialize: grpc.deserialize<iam_pb.SignUpRequest>;
    responseSerialize: grpc.serialize<iam_pb.SignUpResponse>;
    responseDeserialize: grpc.deserialize<iam_pb.SignUpResponse>;
}

export const IAMServiceService: IIAMServiceService;

export interface IIAMServiceServer extends grpc.UntypedServiceImplementation {
    createPolicy: grpc.handleUnaryCall<iam_pb.CreatePolicyRequest, iam_pb.CreatePolicyResponse>;
    getPolicy: grpc.handleUnaryCall<iam_pb.GetPolicyRequest, iam_pb.GetPolicyResponse>;
    updatePolicy: grpc.handleUnaryCall<iam_pb.UpdatePolicyRequest, iam_pb.UpdatePolicyResponse>;
    deletePolicy: grpc.handleUnaryCall<iam_pb.DeletePolicyRequest, iam_pb.DeletePolicyResponse>;
    createRole: grpc.handleUnaryCall<iam_pb.CreateRoleRequest, iam_pb.CreateRoleResponse>;
    getRole: grpc.handleUnaryCall<iam_pb.GetRoleRequest, iam_pb.GetRoleResponse>;
    updateRole: grpc.handleUnaryCall<iam_pb.UpdateRoleRequest, iam_pb.UpdateRoleResponse>;
    deleteRole: grpc.handleUnaryCall<iam_pb.DeleteRoleRequest, iam_pb.DeleteRoleResponse>;
    createUser: grpc.handleUnaryCall<iam_pb.CreateUserRequest, iam_pb.CreateUserResponse>;
    getUser: grpc.handleUnaryCall<iam_pb.GetUserRequest, iam_pb.GetUserResponse>;
    updateUser: grpc.handleUnaryCall<iam_pb.UpdateUserRequest, iam_pb.UpdateUserResponse>;
    deleteUser: grpc.handleUnaryCall<iam_pb.DeleteUserRequest, google_protobuf_empty_pb.Empty>;
    createGroup: grpc.handleUnaryCall<iam_pb.CreateGroupRequest, iam_pb.CreateGroupResponse>;
    getGroup: grpc.handleUnaryCall<iam_pb.GetGroupRequest, iam_pb.GetGroupResponse>;
    updateGroup: grpc.handleUnaryCall<iam_pb.UpdateGroupRequest, iam_pb.UpdateGroupResponse>;
    deleteGroup: grpc.handleUnaryCall<iam_pb.DeleteGroupRequest, iam_pb.DeleteGroupResponse>;
    createAccessKeys: grpc.handleUnaryCall<iam_pb.CreateAccessKeyRequest, iam_pb.CreateAccessKeysResponse>;
    getAccessKeys: grpc.handleUnaryCall<iam_pb.GetAccessKeysRequest, iam_pb.GetAccessKeysResponse>;
    deleteAccessKeys: grpc.handleUnaryCall<iam_pb.DeleteAccessKeysRequest, google_protobuf_empty_pb.Empty>;
    userLogin: grpc.handleUnaryCall<iam_pb.UserLoginRequest, iam_pb.LoginResponse>;
    rootUserLogin: grpc.handleUnaryCall<iam_pb.RootUserLoginRequest, iam_pb.LoginResponse>;
    signUp: grpc.handleUnaryCall<iam_pb.SignUpRequest, iam_pb.SignUpResponse>;
}

export interface IIAMServiceClient {
    createPolicy(request: iam_pb.CreatePolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreatePolicyResponse) => void): grpc.ClientUnaryCall;
    createPolicy(request: iam_pb.CreatePolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreatePolicyResponse) => void): grpc.ClientUnaryCall;
    createPolicy(request: iam_pb.CreatePolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreatePolicyResponse) => void): grpc.ClientUnaryCall;
    getPolicy(request: iam_pb.GetPolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetPolicyResponse) => void): grpc.ClientUnaryCall;
    getPolicy(request: iam_pb.GetPolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetPolicyResponse) => void): grpc.ClientUnaryCall;
    getPolicy(request: iam_pb.GetPolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetPolicyResponse) => void): grpc.ClientUnaryCall;
    updatePolicy(request: iam_pb.UpdatePolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdatePolicyResponse) => void): grpc.ClientUnaryCall;
    updatePolicy(request: iam_pb.UpdatePolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdatePolicyResponse) => void): grpc.ClientUnaryCall;
    updatePolicy(request: iam_pb.UpdatePolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdatePolicyResponse) => void): grpc.ClientUnaryCall;
    deletePolicy(request: iam_pb.DeletePolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.DeletePolicyResponse) => void): grpc.ClientUnaryCall;
    deletePolicy(request: iam_pb.DeletePolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.DeletePolicyResponse) => void): grpc.ClientUnaryCall;
    deletePolicy(request: iam_pb.DeletePolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.DeletePolicyResponse) => void): grpc.ClientUnaryCall;
    createRole(request: iam_pb.CreateRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    createRole(request: iam_pb.CreateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    createRole(request: iam_pb.CreateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    getRole(request: iam_pb.GetRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetRoleResponse) => void): grpc.ClientUnaryCall;
    getRole(request: iam_pb.GetRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetRoleResponse) => void): grpc.ClientUnaryCall;
    getRole(request: iam_pb.GetRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetRoleResponse) => void): grpc.ClientUnaryCall;
    updateRole(request: iam_pb.UpdateRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    updateRole(request: iam_pb.UpdateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    updateRole(request: iam_pb.UpdateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    deleteRole(request: iam_pb.DeleteRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteRoleResponse) => void): grpc.ClientUnaryCall;
    deleteRole(request: iam_pb.DeleteRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteRoleResponse) => void): grpc.ClientUnaryCall;
    deleteRole(request: iam_pb.DeleteRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteRoleResponse) => void): grpc.ClientUnaryCall;
    createUser(request: iam_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: iam_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: iam_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: iam_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: iam_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: iam_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    updateUser(request: iam_pb.UpdateUserRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    updateUser(request: iam_pb.UpdateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    updateUser(request: iam_pb.UpdateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    deleteUser(request: iam_pb.DeleteUserRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteUser(request: iam_pb.DeleteUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteUser(request: iam_pb.DeleteUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    createGroup(request: iam_pb.CreateGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    createGroup(request: iam_pb.CreateGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    createGroup(request: iam_pb.CreateGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    getGroup(request: iam_pb.GetGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetGroupResponse) => void): grpc.ClientUnaryCall;
    getGroup(request: iam_pb.GetGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetGroupResponse) => void): grpc.ClientUnaryCall;
    getGroup(request: iam_pb.GetGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetGroupResponse) => void): grpc.ClientUnaryCall;
    updateGroup(request: iam_pb.UpdateGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateGroupResponse) => void): grpc.ClientUnaryCall;
    updateGroup(request: iam_pb.UpdateGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateGroupResponse) => void): grpc.ClientUnaryCall;
    updateGroup(request: iam_pb.UpdateGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateGroupResponse) => void): grpc.ClientUnaryCall;
    deleteGroup(request: iam_pb.DeleteGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteGroupResponse) => void): grpc.ClientUnaryCall;
    deleteGroup(request: iam_pb.DeleteGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteGroupResponse) => void): grpc.ClientUnaryCall;
    deleteGroup(request: iam_pb.DeleteGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteGroupResponse) => void): grpc.ClientUnaryCall;
    createAccessKeys(request: iam_pb.CreateAccessKeyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateAccessKeysResponse) => void): grpc.ClientUnaryCall;
    createAccessKeys(request: iam_pb.CreateAccessKeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateAccessKeysResponse) => void): grpc.ClientUnaryCall;
    createAccessKeys(request: iam_pb.CreateAccessKeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateAccessKeysResponse) => void): grpc.ClientUnaryCall;
    getAccessKeys(request: iam_pb.GetAccessKeysRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetAccessKeysResponse) => void): grpc.ClientUnaryCall;
    getAccessKeys(request: iam_pb.GetAccessKeysRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetAccessKeysResponse) => void): grpc.ClientUnaryCall;
    getAccessKeys(request: iam_pb.GetAccessKeysRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetAccessKeysResponse) => void): grpc.ClientUnaryCall;
    deleteAccessKeys(request: iam_pb.DeleteAccessKeysRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteAccessKeys(request: iam_pb.DeleteAccessKeysRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteAccessKeys(request: iam_pb.DeleteAccessKeysRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    userLogin(request: iam_pb.UserLoginRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    userLogin(request: iam_pb.UserLoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    userLogin(request: iam_pb.UserLoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    rootUserLogin(request: iam_pb.RootUserLoginRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    rootUserLogin(request: iam_pb.RootUserLoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    rootUserLogin(request: iam_pb.RootUserLoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    signUp(request: iam_pb.SignUpRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.SignUpResponse) => void): grpc.ClientUnaryCall;
    signUp(request: iam_pb.SignUpRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.SignUpResponse) => void): grpc.ClientUnaryCall;
    signUp(request: iam_pb.SignUpRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.SignUpResponse) => void): grpc.ClientUnaryCall;
}

export class IAMServiceClient extends grpc.Client implements IIAMServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createPolicy(request: iam_pb.CreatePolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreatePolicyResponse) => void): grpc.ClientUnaryCall;
    public createPolicy(request: iam_pb.CreatePolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreatePolicyResponse) => void): grpc.ClientUnaryCall;
    public createPolicy(request: iam_pb.CreatePolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreatePolicyResponse) => void): grpc.ClientUnaryCall;
    public getPolicy(request: iam_pb.GetPolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetPolicyResponse) => void): grpc.ClientUnaryCall;
    public getPolicy(request: iam_pb.GetPolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetPolicyResponse) => void): grpc.ClientUnaryCall;
    public getPolicy(request: iam_pb.GetPolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetPolicyResponse) => void): grpc.ClientUnaryCall;
    public updatePolicy(request: iam_pb.UpdatePolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdatePolicyResponse) => void): grpc.ClientUnaryCall;
    public updatePolicy(request: iam_pb.UpdatePolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdatePolicyResponse) => void): grpc.ClientUnaryCall;
    public updatePolicy(request: iam_pb.UpdatePolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdatePolicyResponse) => void): grpc.ClientUnaryCall;
    public deletePolicy(request: iam_pb.DeletePolicyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.DeletePolicyResponse) => void): grpc.ClientUnaryCall;
    public deletePolicy(request: iam_pb.DeletePolicyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.DeletePolicyResponse) => void): grpc.ClientUnaryCall;
    public deletePolicy(request: iam_pb.DeletePolicyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.DeletePolicyResponse) => void): grpc.ClientUnaryCall;
    public createRole(request: iam_pb.CreateRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    public createRole(request: iam_pb.CreateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    public createRole(request: iam_pb.CreateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    public getRole(request: iam_pb.GetRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetRoleResponse) => void): grpc.ClientUnaryCall;
    public getRole(request: iam_pb.GetRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetRoleResponse) => void): grpc.ClientUnaryCall;
    public getRole(request: iam_pb.GetRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetRoleResponse) => void): grpc.ClientUnaryCall;
    public updateRole(request: iam_pb.UpdateRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    public updateRole(request: iam_pb.UpdateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    public updateRole(request: iam_pb.UpdateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    public deleteRole(request: iam_pb.DeleteRoleRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteRoleResponse) => void): grpc.ClientUnaryCall;
    public deleteRole(request: iam_pb.DeleteRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteRoleResponse) => void): grpc.ClientUnaryCall;
    public deleteRole(request: iam_pb.DeleteRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteRoleResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: iam_pb.CreateUserRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: iam_pb.CreateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: iam_pb.CreateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateUserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: iam_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: iam_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: iam_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetUserResponse) => void): grpc.ClientUnaryCall;
    public updateUser(request: iam_pb.UpdateUserRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    public updateUser(request: iam_pb.UpdateUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    public updateUser(request: iam_pb.UpdateUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateUserResponse) => void): grpc.ClientUnaryCall;
    public deleteUser(request: iam_pb.DeleteUserRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteUser(request: iam_pb.DeleteUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteUser(request: iam_pb.DeleteUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public createGroup(request: iam_pb.CreateGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    public createGroup(request: iam_pb.CreateGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    public createGroup(request: iam_pb.CreateGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateGroupResponse) => void): grpc.ClientUnaryCall;
    public getGroup(request: iam_pb.GetGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetGroupResponse) => void): grpc.ClientUnaryCall;
    public getGroup(request: iam_pb.GetGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetGroupResponse) => void): grpc.ClientUnaryCall;
    public getGroup(request: iam_pb.GetGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetGroupResponse) => void): grpc.ClientUnaryCall;
    public updateGroup(request: iam_pb.UpdateGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateGroupResponse) => void): grpc.ClientUnaryCall;
    public updateGroup(request: iam_pb.UpdateGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateGroupResponse) => void): grpc.ClientUnaryCall;
    public updateGroup(request: iam_pb.UpdateGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.UpdateGroupResponse) => void): grpc.ClientUnaryCall;
    public deleteGroup(request: iam_pb.DeleteGroupRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteGroupResponse) => void): grpc.ClientUnaryCall;
    public deleteGroup(request: iam_pb.DeleteGroupRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteGroupResponse) => void): grpc.ClientUnaryCall;
    public deleteGroup(request: iam_pb.DeleteGroupRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.DeleteGroupResponse) => void): grpc.ClientUnaryCall;
    public createAccessKeys(request: iam_pb.CreateAccessKeyRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateAccessKeysResponse) => void): grpc.ClientUnaryCall;
    public createAccessKeys(request: iam_pb.CreateAccessKeyRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateAccessKeysResponse) => void): grpc.ClientUnaryCall;
    public createAccessKeys(request: iam_pb.CreateAccessKeyRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.CreateAccessKeysResponse) => void): grpc.ClientUnaryCall;
    public getAccessKeys(request: iam_pb.GetAccessKeysRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.GetAccessKeysResponse) => void): grpc.ClientUnaryCall;
    public getAccessKeys(request: iam_pb.GetAccessKeysRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.GetAccessKeysResponse) => void): grpc.ClientUnaryCall;
    public getAccessKeys(request: iam_pb.GetAccessKeysRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.GetAccessKeysResponse) => void): grpc.ClientUnaryCall;
    public deleteAccessKeys(request: iam_pb.DeleteAccessKeysRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteAccessKeys(request: iam_pb.DeleteAccessKeysRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteAccessKeys(request: iam_pb.DeleteAccessKeysRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public userLogin(request: iam_pb.UserLoginRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public userLogin(request: iam_pb.UserLoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public userLogin(request: iam_pb.UserLoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public rootUserLogin(request: iam_pb.RootUserLoginRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public rootUserLogin(request: iam_pb.RootUserLoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public rootUserLogin(request: iam_pb.RootUserLoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public signUp(request: iam_pb.SignUpRequest, callback: (error: grpc.ServiceError | null, response: iam_pb.SignUpResponse) => void): grpc.ClientUnaryCall;
    public signUp(request: iam_pb.SignUpRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: iam_pb.SignUpResponse) => void): grpc.ClientUnaryCall;
    public signUp(request: iam_pb.SignUpRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: iam_pb.SignUpResponse) => void): grpc.ClientUnaryCall;
}
