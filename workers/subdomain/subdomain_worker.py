import requests
import argparse
import os
import pika 
import subprocess

# python3 ./subdomain_worker.py  -r 192.168.0.20 -u guest -p guest -s 192.168.0.20

rabbitmq_host = ""
rabbitmq_user = ""
rabbitmq_pass = ""
api_server = ""

# create a function to consume the queue but not ack the message until after the task is complete
def consume_queue():
    
    print(' [*] Waiting for messages. To exit press CTRL+C')
    
    connection = pika.BlockingConnection(
        pika.ConnectionParameters(host="192.168.0.20", port=5672)
    )
    channel = connection.channel()
    channel.basic_consume(queue='domain', on_message_callback=run_task, consumer_tag='subdomain_worker')
    channel.start_consuming()


def run_task(ch, method, properties, msg):

    body = msg.decode("utf-8")

    out_dir = "/root/app/workers/output/recon/" + str(body) + "/subdomains/"

    # Create output direcetory
    os.makedirs(out_dir, exist_ok=True)

    # Run subfinder
    print("[+] Running subfinder")
    subfinder = subprocess.Popen(["subfinder", "-d", body, "-o", out_dir + "subfinder.txt"])

    # Run assetfinder
    print("[+] Running assetfinder")
    assetfinder = subprocess.Popen(["assetfinder", "--subs-only", body,  " >> ", out_dir + "assetfinder.txt"])

    # Run amass
    print("[+] Running amass")
    amass = subprocess.Popen(["amass", "enum", "-passive", "-d", body, "-o", out_dir + "amass.txt"])

    # Run sublist3r
    print("[+] Running sublist3r")
    sublister = subprocess.Popen(["sublist3r", "-d", body, "-o", out_dir + "sublist3r.txt"])

    subfinder.communicate()
    assetfinder.communicate()
    amass.communicate()
    sublister.communicate()


    unique_results = []

    # Read all of the files in the output directory
    all_files = os.listdir(out_dir)
    for filename in all_files:
        if filename.endswith(".txt"):
            with open(out_dir + filename, 'r') as f:
                lines = f.readlines()
                for line in lines:
                    if line not in unique_results:
                        unique_results.append(line)

                        # Publish the results to the queue
                        ch.basic_publish(exchange='', routing_key='subdomain', body=line)

    # Write the unique results to a file
    with open(out_dir + "unique.txt", 'w') as f:
        for line in unique_results:
            f.write(line)

    # Acknowledge the message
    ch.basic_ack(delivery_tag=method.delivery_tag)

    # Send the results to the API
    r = requests.put('http://192.168.0.20:6123/api/v1/domain/subdomain/create', json={'domains': unique_results})
    
    if r.status_code == 200:
        print("[+] Successfully sent results to API")

def main():
    # Get flags 
    parser = argparse.ArgumentParser()
    parser.add_argument('-r', '--rmq-host', help='RabbitMQ Host', required=True)
    parser.add_argument('-s', '--api-server', help='API Host', required=True)
    parser.add_argument('-u', '--rmq-user', help='RabbitMQ User', required=True)
    parser.add_argument('-p', '--rmq-pass', help='RabbitMQ Password', required=True)
    args = parser.parse_args()

    rabbitmq_host = args.rmq_host
    rabbitmq_user = args.rmq_user
    rabbitmq_pass = args.rmq_pass
    api_server = args.api_server
    
    # while True:
    consume_queue()

if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)