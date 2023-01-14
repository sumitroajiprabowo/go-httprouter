pipeline{
    agent none
    tools {
        go 'go1.19.5'
    }
    environment {
        AUTHOR = 'Sumitro Aji Prabowo'
        EMAIL = 'bowo@anakdesa.id'
    }

    triggers {
        //cron('* * * * *')
        pollSCM('* * * * *')
    }

    parameters {
        string(name: 'BRANCH', defaultValue: 'master', description: 'Branch to build')
    }

    options {
        buildDiscarder(logRotator(numToKeepStr: '5'))
        disableConcurrentBuilds()
        timeout(time: 1, unit: 'HOURS')
    }

    stages {
        stage ('Prepare') {
            agent {
                node {
                    label 'golang'
                }
            }
            environment {
                APP = credentials('dockerhub-sumitroajiprabowo')
            }
            steps {
                echo ("Author, ${env.AUTHOR}")
                echo ("Email, ${env.EMAIL}")
                echo ("Docker Hub Username, ${env.APP_USR}")
                echo ("Docker Hub Password, ${env.APP_PSW}")
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
            input {
                message 'Deploy?'
                ok 'Yes, we should.'
                submitter 'sumitroajiprabowo, megadata'
                parameters {
                    choice(name: 'TARGET ENVIRONMENT', choices: ['DEV', 'STAGING', 'PRODUCTION'], description: 'Choose the target environment')
                }
            }
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