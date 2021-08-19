# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import lambda_pb2 as lambda__pb2


class LambdaServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateFunction = channel.unary_unary(
                '/lambda.LambdaService/CreateFunction',
                request_serializer=lambda__pb2.CreateFunctionRequest.SerializeToString,
                response_deserializer=lambda__pb2.LambdaResponse.FromString,
                )
        self.TestFunction = channel.unary_unary(
                '/lambda.LambdaService/TestFunction',
                request_serializer=lambda__pb2.TestFunctionResquest.SerializeToString,
                response_deserializer=lambda__pb2.LambdaResponse.FromString,
                )
        self.InvoqueFunction = channel.unary_unary(
                '/lambda.LambdaService/InvoqueFunction',
                request_serializer=lambda__pb2.InvoqueFunctionRequest.SerializeToString,
                response_deserializer=lambda__pb2.LambdaResponse.FromString,
                )
        self.SeedLambdaServer = channel.unary_unary(
                '/lambda.LambdaService/SeedLambdaServer',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=lambda__pb2.LambdaResponse.FromString,
                )
        self.ReceiveEvents = channel.unary_stream(
                '/lambda.LambdaService/ReceiveEvents',
                request_serializer=lambda__pb2.ReceiveEventRequest.SerializeToString,
                response_deserializer=lambda__pb2.EventResponse.FromString,
                )


class LambdaServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateFunction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TestFunction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def InvoqueFunction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SeedLambdaServer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReceiveEvents(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LambdaServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateFunction': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateFunction,
                    request_deserializer=lambda__pb2.CreateFunctionRequest.FromString,
                    response_serializer=lambda__pb2.LambdaResponse.SerializeToString,
            ),
            'TestFunction': grpc.unary_unary_rpc_method_handler(
                    servicer.TestFunction,
                    request_deserializer=lambda__pb2.TestFunctionResquest.FromString,
                    response_serializer=lambda__pb2.LambdaResponse.SerializeToString,
            ),
            'InvoqueFunction': grpc.unary_unary_rpc_method_handler(
                    servicer.InvoqueFunction,
                    request_deserializer=lambda__pb2.InvoqueFunctionRequest.FromString,
                    response_serializer=lambda__pb2.LambdaResponse.SerializeToString,
            ),
            'SeedLambdaServer': grpc.unary_unary_rpc_method_handler(
                    servicer.SeedLambdaServer,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=lambda__pb2.LambdaResponse.SerializeToString,
            ),
            'ReceiveEvents': grpc.unary_stream_rpc_method_handler(
                    servicer.ReceiveEvents,
                    request_deserializer=lambda__pb2.ReceiveEventRequest.FromString,
                    response_serializer=lambda__pb2.EventResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'lambda.LambdaService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class LambdaService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateFunction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/lambda.LambdaService/CreateFunction',
            lambda__pb2.CreateFunctionRequest.SerializeToString,
            lambda__pb2.LambdaResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TestFunction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/lambda.LambdaService/TestFunction',
            lambda__pb2.TestFunctionResquest.SerializeToString,
            lambda__pb2.LambdaResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def InvoqueFunction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/lambda.LambdaService/InvoqueFunction',
            lambda__pb2.InvoqueFunctionRequest.SerializeToString,
            lambda__pb2.LambdaResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SeedLambdaServer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/lambda.LambdaService/SeedLambdaServer',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            lambda__pb2.LambdaResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ReceiveEvents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/lambda.LambdaService/ReceiveEvents',
            lambda__pb2.ReceiveEventRequest.SerializeToString,
            lambda__pb2.EventResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
