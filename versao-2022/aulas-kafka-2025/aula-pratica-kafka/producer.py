import pika
import json
from flask import Flask, request, jsonify
from functools import wraps

app = Flask(__name__)

def setup_rabbitmq_connection():
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()
        channel.queue_declare(queue='hello')
        return connection, channel
    except pika.exceptions.AMQPConnectionError:
        return None, None

def require_json(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if not request.is_json:
            return jsonify({"error": "Content-Type must be application/json"}), 400
        return f(*args, **kwargs)
    return decorated_function

@app.route('/health', methods=['GET'])
def health_check():
    """Health check endpoint to verify if the service is running"""
    return jsonify({"status": "healthy"}), 200

@app.route('/send', methods=['POST'])
@require_json
def send():
    """Send a message to RabbitMQ queue"""
    try:
        data = request.json
        connection, channel = setup_rabbitmq_connection()
        
        if not connection or not channel:
            return jsonify({"error": "Failed to connect to RabbitMQ"}), 503

        channel.basic_publish(
            exchange='',
            routing_key='hello',
            body=json.dumps(data)
        )
        connection.close()
        return jsonify({'status': 'sent', 'data': data}), 200
    
    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/send_batch', methods=['POST'])
@require_json
def send_batch():
    """Send multiple messages to RabbitMQ queue"""
    try:
        data = request.json
        if not isinstance(data, list):
            return jsonify({"error": "Request body must be an array"}), 400

        connection, channel = setup_rabbitmq_connection()
        
        if not connection or not channel:
            return jsonify({"error": "Failed to connect to RabbitMQ"}), 503

        for message in data:
            channel.basic_publish(
                exchange='',
                routing_key='hello',
                body=json.dumps(message)
            )
        
        connection.close()
        return jsonify({
            'status': 'sent',
            'message_count': len(data)
        }), 200
    
    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.errorhandler(404)
def not_found(error):
    return jsonify({"error": "Not found"}), 404

@app.errorhandler(500)
def internal_error(error):
    return jsonify({"error": "Internal server error"}), 500

if __name__ == '__main__':
    app.run(debug=True, port=5000) 