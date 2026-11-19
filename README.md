# go-node-info

A small Go server that serves a basic webpage for practicing with EKS.

## Run locally

With Go 1.25 or newer installed, run:

```sh
go run .
```

## Run with Docker

```sh
docker build -t go-node-info .
docker run --rm -p 8080:8080 go-node-info
```

Open [http://localhost:8080](http://localhost:8080).

## Publish to Amazon ECR

Each push to `main` builds the Docker image and pushes it to a private ECR repository with the full Git commit SHA as its tag.

Create the ECR repository in AWS. Set up a GitHub OIDC IAM role that trusts `repo:Stevendylan01/go-node-info:ref:refs/heads/main` and can push images to that repository. See the [GitHub OIDC setup guide](https://docs.github.com/en/actions/how-tos/secure-your-work/security-harden-deployments/oidc-in-aws) and [ECR push permissions](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-push-iam.html).

In GitHub, add these repository variables under **Settings → Secrets and variables → Actions → Variables**.

| Variable | Value |
| --- | --- |
| `AWS_ROLE_ARN` | ARN of the GitHub OIDC IAM role |
| `AWS_REGION` | AWS region of the ECR repository |
| `ECR_REPOSITORY` | ECR repository name, for example `go-node-info` |
