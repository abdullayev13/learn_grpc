package org.grpc_client;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import org.grpc_client.grpc.Chat;
import org.grpc_client.grpc.ChatServiceGrpc;

public class Main {
    static ChatServiceGrpc.ChatServiceBlockingStub stub;

    public static void main(String[] args) {
        System.out.println("hello");
        ManagedChannel channel = ManagedChannelBuilder
                .forAddress("localhost", 9001)
                .usePlaintext()
                .build();


        stub = ChatServiceGrpc.newBlockingStub(channel);


        for (int i = 0; i < 300; i++) {
            new Thread(() -> fun(300)).start();
        }
        fun(100);
    }

    static void fun(int n) {
        for (int i = 0; i < n; i++) {
            Chat.Message request = Chat.Message.newBuilder()
                    .setBody("Hello (" + i + ")")
                    .build();

            Chat.Message response = stub.sayHello(request);

            System.out.println("Response : " + response.getBody());
        }
    }
}
