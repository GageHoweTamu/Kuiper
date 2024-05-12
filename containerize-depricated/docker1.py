import docker

def main():
    try:
        client = docker.from_env()
        client.containers.run("my-image",
                                name = "my-container",
                                detach = True,
                                ports = {"8080/tcp": 8080},
        )
        print("Container is running...")
        client.containers.list()
    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()