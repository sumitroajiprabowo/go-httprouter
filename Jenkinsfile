pipeline{
    agent {
        node {
            label 'golang'
        }
    }
    tools {
        go 'go1.19.5'
    }

    stages {
        stage('Build') {
            steps {
            script {
            def exampleJSON = [
                name: 'example',
                version: '1.0.0',
            ]
            def exampleJSONString = exampleJSON.toString()
            writeFile file: 'example.json', text: exampleJSONString
            }
                echo("Start Build ")
                sh "go build -o main"
                echo("Finish Build")
            }
        }

        stage('Test') {
            steps {
                echo("Hello Test 1")
                sh 'go test -v ./...'
                echo("Hello Test 2")
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