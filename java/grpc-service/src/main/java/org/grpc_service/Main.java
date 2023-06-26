package org.grpc_service;

import io.grpc.Server;
import io.grpc.ServerBuilder;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException {

        Server server = ServerBuilder
                .forPort(9001)
                .addService(new ChatService())
                .build();

        server.start();

        System.out.println("Server started on " + server.getPort());

        server.awaitTermination();
    }
}