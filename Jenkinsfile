pipeline {
    agent any
        tools {
        go 'go-1.14'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh '''cd src
                      go install
                      go build
                '''
            }
        }
        stage('Test') {
            steps {
                sh 'echo "Test"'
                sh '''cd src
                      go test
                '''
            }
        }
    }
    post {
        failure {
            emailext body: 'Failed to build FH', subject: 'Build Failure', to: '$DEFAULT_RECIPIENTS'
        }
    }
}
