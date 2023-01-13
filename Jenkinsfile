pipeline{
    agent none
    tools {
        go 'go1.19.5'
    }
    environment {
        AUTHOR = 'Sumitro Aji Prabowo'
        EMAIL = 'bowo@anakdesa.id'
    }

    stages {
        stage ('Prepare') {
            agent {
                node {
                    label 'golang'
                }
            }
            steps {
                echo ("Author, ${env.AUTHOR}")
                echo ("Email, ${env.EMAIL}")
                echo ("Start Job ${env.JOB_NAME}")
                echo ("Start Build ${env.BUILD_NUMBER}")
                echo ("Start Branch ${env.BRANCH_NAME}")
                sh 'go version'
            }
        }
        stage('Build') {
            agent {
                node {
                    label 'golang'
                }
            }
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
            agent {
                node {
                        label 'golang'
                }
            }
            steps {
                echo("Hello Test 1")
                sh 'go test -v ./...'
                echo("Hello Test 2")
            }
        }

        stage('Deploy') {
            agent {
                node {
                        label 'golang'
                }
            }
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