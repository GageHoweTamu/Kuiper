import docker

client = docker.from_env()

print("Creating and running the container...")
container = client.containers.run("ubuntu", "echo hi")

print("Waiting for the container to finish...")
exit_code = container.wait()["StatusCode"] # Wait for the container to finish

print(f"Container exited with code: {exit_code}")
logs = container.logs(stdout=True, stderr=True).decode("utf-8")
print("Container logs:")
print(logs)

container.remove()