pipeline {
  agent any
  stages {
    stage('Pre Test') {
      steps {
        echo 'Installing dependencies'
        sh 'go version'
        sh 'sh \'go get -u golang.org/x/lint/golint\''
      }
    }

    stage('Build') {
      steps {
        echo 'Compiling and building'
        sh 'cd cmd/shopping4chow/ && go build .'
      }
    }

    stage('Test') {
      steps {
        withEnv(overrides: ["PATH+GO=${GOPATH}/bin"]) {
          echo 'Running vetting'
          sh 'go vet .'
          echo 'Running linting'
          sh 'golint .'
          echo 'Running test'
          sh 'cd cmd/shopping4chow/dao/ && go test -v'
        }

      }
    }

  }
  tools {
    go 'LocalGo'
  }
  environment {
    GO114MODULE = 'on'
    CGO_ENABLED = 0
    GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
  }
  post {
    always {
      emailext(body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}", recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']], to: "${params.RECIPIENTS}", subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}")
    }

  }
}