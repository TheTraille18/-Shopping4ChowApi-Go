pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''echo \'Compiling and building\'


'''
        sh 'ls'
        sh 'cd cmd'
        sh 'cd shopping4chow'
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
        sh 'cd dao'
        sh 'go test -v'
      }
    }

  }
}