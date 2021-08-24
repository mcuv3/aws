// package: iam
// file: iam.proto

var iam_pb = require("./iam_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var IAMService = (function () {
  function IAMService() {}
  IAMService.serviceName = "iam.IAMService";
  return IAMService;
}());

IAMService.CreatePolicy = {
  methodName: "CreatePolicy",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.CreatePolicyRequest,
  responseType: iam_pb.CreatePolicyResponse
};

IAMService.GetPolicy = {
  methodName: "GetPolicy",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.GetPolicyRequest,
  responseType: iam_pb.GetPolicyResponse
};

IAMService.UpdatePolicy = {
  methodName: "UpdatePolicy",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.UpdatePolicyRequest,
  responseType: iam_pb.UpdatePolicyResponse
};

IAMService.DeletePolicy = {
  methodName: "DeletePolicy",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.DeletePolicyRequest,
  responseType: iam_pb.DeletePolicyResponse
};

IAMService.CreateRole = {
  methodName: "CreateRole",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.CreateRoleRequest,
  responseType: iam_pb.CreateRoleResponse
};

IAMService.GetRole = {
  methodName: "GetRole",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.GetRoleRequest,
  responseType: iam_pb.GetRoleResponse
};

IAMService.UpdateRole = {
  methodName: "UpdateRole",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.UpdateRoleRequest,
  responseType: iam_pb.UpdateRoleResponse
};

IAMService.DeleteRole = {
  methodName: "DeleteRole",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.DeleteRoleRequest,
  responseType: iam_pb.DeleteRoleResponse
};

IAMService.CreateUser = {
  methodName: "CreateUser",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.CreateUserRequest,
  responseType: iam_pb.CreateUserResponse
};

IAMService.GetUser = {
  methodName: "GetUser",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.GetUserRequest,
  responseType: iam_pb.GetUserResponse
};

IAMService.UpdateUser = {
  methodName: "UpdateUser",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.UpdateUserRequest,
  responseType: iam_pb.UpdateUserResponse
};

IAMService.DeleteUser = {
  methodName: "DeleteUser",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.DeleteUserRequest,
  responseType: google_protobuf_empty_pb.Empty
};

IAMService.CreateGroup = {
  methodName: "CreateGroup",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.CreateGroupRequest,
  responseType: iam_pb.CreateGroupResponse
};

IAMService.GetGroup = {
  methodName: "GetGroup",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.GetGroupRequest,
  responseType: iam_pb.GetGroupResponse
};

IAMService.UpdateGroup = {
  methodName: "UpdateGroup",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.UpdateGroupRequest,
  responseType: iam_pb.UpdateGroupResponse
};

IAMService.DeleteGroup = {
  methodName: "DeleteGroup",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.DeleteGroupRequest,
  responseType: iam_pb.DeleteGroupResponse
};

IAMService.CreateAccessKeys = {
  methodName: "CreateAccessKeys",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.CreateAccessKeyRequest,
  responseType: iam_pb.CreateAccessKeysResponse
};

IAMService.GetAccessKeys = {
  methodName: "GetAccessKeys",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.GetAccessKeysRequest,
  responseType: iam_pb.GetAccessKeysResponse
};

IAMService.DeleteAccessKeys = {
  methodName: "DeleteAccessKeys",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.DeleteAccessKeysRequest,
  responseType: google_protobuf_empty_pb.Empty
};

IAMService.UserLogin = {
  methodName: "UserLogin",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.UserLoginRequest,
  responseType: iam_pb.LoginResponse
};

IAMService.RootUserLogin = {
  methodName: "RootUserLogin",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.RootUserLoginRequest,
  responseType: iam_pb.LoginResponse
};

IAMService.SignUp = {
  methodName: "SignUp",
  service: IAMService,
  requestStream: false,
  responseStream: false,
  requestType: iam_pb.SignUpRequest,
  responseType: iam_pb.SignUpResponse
};

exports.IAMService = IAMService;

function IAMServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

IAMServiceClient.prototype.createPolicy = function createPolicy(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.CreatePolicy, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.getPolicy = function getPolicy(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.GetPolicy, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.updatePolicy = function updatePolicy(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.UpdatePolicy, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.deletePolicy = function deletePolicy(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.DeletePolicy, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.createRole = function createRole(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.CreateRole, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.getRole = function getRole(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.GetRole, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.updateRole = function updateRole(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.UpdateRole, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.deleteRole = function deleteRole(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.DeleteRole, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.createUser = function createUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.CreateUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.getUser = function getUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.GetUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.updateUser = function updateUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.UpdateUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.deleteUser = function deleteUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.DeleteUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.createGroup = function createGroup(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.CreateGroup, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.getGroup = function getGroup(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.GetGroup, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.updateGroup = function updateGroup(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.UpdateGroup, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.deleteGroup = function deleteGroup(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.DeleteGroup, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.createAccessKeys = function createAccessKeys(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.CreateAccessKeys, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.getAccessKeys = function getAccessKeys(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.GetAccessKeys, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.deleteAccessKeys = function deleteAccessKeys(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.DeleteAccessKeys, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.userLogin = function userLogin(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.UserLogin, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.rootUserLogin = function rootUserLogin(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.RootUserLogin, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

IAMServiceClient.prototype.signUp = function signUp(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(IAMService.SignUp, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.IAMServiceClient = IAMServiceClient;

