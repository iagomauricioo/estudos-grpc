import grpc
from concurrent import futures
import logging
import sys
import os

# Adicionar o diretório src/pb ao path para importar os módulos gerados
sys.path.append(os.path.join(os.path.dirname(__file__), 'src', 'pb'))

import calc_service_pb2
import calc_service_pb2_grpc


class CalcServiceServicer(calc_service_pb2_grpc.CalcServiceServicer):
    def Calc(self, request_iterator, context):
        quantity = 0
        total = 0
        
        for input_data in request_iterator:
            quantity += 1
            total += input_data.value
            print(f"input: {input_data}")
        
        if quantity > 0:
            avg = float(total) / quantity
        else:
            avg = 0.0
            
        return calc_service_pb2.Output(
            quantity=quantity,
            average=avg,
            sum=total
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    calc_service_pb2_grpc.add_CalcServiceServicer_to_server(
        CalcServiceServicer(), server
    )
    
    listen_addr = '[::]:9090'
    server.add_insecure_port(listen_addr)
    server.start()
    print("gRPC server started port 9090")
    server.wait_for_termination()


def main():
    print("starting gRPC server")
    serve()


if __name__ == "__main__":
    main()