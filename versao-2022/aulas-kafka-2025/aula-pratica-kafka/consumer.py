import pika
import json
import time

def callback(ch, method, properties, body):
    try:
        message = json.loads(body)
        print(f"[x] Received: {message}")
        
        # Aumenta o tempo de processamento simulado para 5 segundos
        print(f"[x] Processing message for 5 seconds...")
        time.sleep(5)
        
        # Confirma que a mensagem foi processada com sucesso
        ch.basic_ack(delivery_tag=method.delivery_tag)
        print(f"[x] Done processing message")
        
    except Exception as e:
        # Em caso de erro, rejeita a mensagem
        # requeue=True coloca a mensagem de volta na fila
        ch.basic_nack(delivery_tag=method.delivery_tag, requeue=True)
        print(f"[!] Error processing message: {str(e)}")

def main():
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()

        # Declara a fila
        channel.queue_declare(queue='hello')
        
        # Configura para enviar apenas uma mensagem por vez para o consumer
        channel.basic_qos(prefetch_count=1)
        
        # auto_ack=False para confirmação manual
        channel.basic_consume(queue='hello', 
                            on_message_callback=callback,
                            auto_ack=False)

        print('[*] Waiting for messages. To exit press CTRL+C')
        channel.start_consuming()
        
    except KeyboardInterrupt:
        print('[!] Interrupted by user')
        try:
            connection.close()
        except:
            pass
    except Exception as e:
        print(f'[!] Error: {str(e)}')

if __name__ == '__main__':
    main()
