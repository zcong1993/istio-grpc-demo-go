workflow "Docker Build and Push" {
  on = "push"
  resolves = ["Docker Push"]
}

action "Docker Login" {
  uses = "actions/docker/login@aea64bb1b97c42fa69b90523667fef56b90d7cff"
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "Docker Build" {
  uses = "actions/docker/cli@aea64bb1b97c42fa69b90523667fef56b90d7cff"
  needs = ["Docker Login"]
  args = "build -t zcong/istio-grpc-demo-go ."
}

action "Docker Push" {
  uses = "actions/docker/cli@aea64bb1b97c42fa69b90523667fef56b90d7cff"
  needs = ["Docker Build"]
  args = "push zcong/istio-grpc-demo-go"
}
