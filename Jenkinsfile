pipeline{
    agent {
        node {
            label 'linux && golang'
        }
    }

    stages {
        stage('Build') {
            steps {
                echo("Hello Build")
            }
        }

        stage('Test') {
            steps {
                echo("Hello Test")
            }
        }

        stage('Deploy') {
            steps {
                echo("Hello Deploy")
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