pipeline{
    agent {
        node {
            label 'linux && golang'
        }
    }

    stages {
        stage('Hello') {
            steps {
                echo("Hello Pipeline From Jenkinsfile")
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