pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''echo \'Compiling and building\'


'''
        sh 'cd cmd/shopping4chow/ && go build .'
      }
    }

    stage('Test') {
      steps {
        sh '''                    echo \'Running vetting\'
                    

sh \'go vet .\'
                    echo \'Running linting\'
                    sh \'golint .\'
                    echo \'Running test\'
                    sh \'cd test && go test -v\''''
        sh 'cd cmd/shopping4chow/dao/ && go test -v'
      }
    }

  }
}