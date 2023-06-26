package org.grpc_service;

import io.grpc.stub.StreamObserver;
import org.grpc_service.grpc.Chat;
import org.grpc_service.grpc.ChatServiceGrpc;

public class ChatService extends ChatServiceGrpc.ChatServiceImplBase {
    static int count;

    @Override
    public void sayHello(Chat.Message request, StreamObserver<Chat.Message> responseObserver) {
        System.out.println(count++);
        System.out.println(request.getBody());

        Chat.Message.Builder response = Chat.Message
                .newBuilder()
                .setBody("got [" + count + "] : " + request.getBody());


        responseObserver.onNext(response.build());
        responseObserver.onCompleted();
    }

}
