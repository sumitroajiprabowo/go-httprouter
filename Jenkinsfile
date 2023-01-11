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
}