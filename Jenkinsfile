pipeline {
  agent any
  stages {
    stage('Build') {
      parallel {
        stage('Build') {
          steps {
            sh 'echo \'Compiling and building\''
          }
        }

        stage('error') {
          steps {
            sh '''sh \'go build\'
'''
          }
        }

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
      }
    }

  }
}