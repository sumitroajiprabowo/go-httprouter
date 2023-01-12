pipeline{
    agent {
        node {
            label 'golang'
        }
    }
    tools {
            go 'go1.19.5'
        }
        environment {
            GO114MODULE = 'on'
            CGO_ENABLED = 0
            GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        }

    stages {
        stage('Build') {
            steps {
                echo("Start Build ")
                sh 'go version'
                echo("Finish Build")
            }
        }

        stage('Test') {
            steps {
                echo("Hello Test 1")
                sleep 10
                echo("Hello Test 2")
            }
        }

        stage('Deploy') {
            steps {
                echo("Hello Deploy")
                sleep 10
                echo("Hello Deploy 2")
            }
        }
    }

    post {
        always {
            echo("Hello Again Pipeline From Jenkinsfile")
        }
        success {
            echo("Hello Success Pipeline From Jenkinsfile")
        }
        failure {
            echo("Hello Failure Pipeline From Jenkinsfile")
        }
        cleanup {
            echo("Hello Cleanup Pipeline From Jenkinsfile")
        }
    }
}