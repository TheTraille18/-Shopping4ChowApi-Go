pipeline {
  agent any
  stages {
    stage('Bduil') {
      parallel {
        stage('Bduil') {
          steps {
            sh 'echo \'Compiling and building\''
          }
        }

        stage('') {
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