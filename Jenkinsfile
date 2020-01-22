// -*- mode: groovy; -*-
pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
  annotations:
    kubernetes.io/target-runtime: kiyot
spec:
  containers:
  - name: golang
    image: elotl/golangbuild:latest
    command:
    - /bin/bash 
    - -c 
    - "sleep 9999"
    tty: true
    resources:
      requests:
        memory: "4Gi"
        cpu: "2000m"
      limits:
        memory: "4Gi"
        cpu: "2000m"
"""
    }
  }
  environment {
    AWS_ZONE = 'us-east-1c'
    AWS_DEFAULT_REGION = 'us-east-1'
    AWS_REGION = 'us-east-1'
    AWS_ACCESS_KEY_ID = credentials('aws-full-access-key-id')
    AWS_SECRET_ACCESS_KEY = credentials('aws-full-access-secret-key')
    IAM_ACCESS_KEY_ID = credentials('aws-ci-iam-key-id')
    IAM_SECRET_ACCESS_KEY = credentials('aws-ci-iam-secret-key')
    BUILD_BUCKET = 'milpa-builds'
    REGISTRY_USER     = 'elotlbuild'
    REGISTRY_PASSWORD = credentials('dockerhub-elotlbuild-password')
  }
  stages {
    stage('Build cloud-instance-provier') {
      steps {
        // Create symlink in GOPATH.
        container('golang') {
          sh 'mkdir -p /go/src/github.com/elotl; ln -s `pwd` /go/src/github.com/elotl/cloud-instance-provider'
        }
        // Run unit tests and build containers
        container('golang') {
          sh 'cd /go/src/github.com/elotl/cloud-instance-provider && ./scripts/run_tests.sh'
        }
      }
    }
  }
  post {
    success {
      slackSend (color: '#00FF00', message: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
    }
    failure {
      slackSend (color: '#FF0000', message: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
    }
  }
}
