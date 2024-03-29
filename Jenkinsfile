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
        booleanParam(name: 'DEPLOY', defaultValue: false, description: 'Deploy to production')
    }

    options {
        disableConcurrentBuilds()
        timeout(time: 1, unit: 'HOURS')
    }

    stages {

        stage ('Prepare') {
            environment {
                APP = credentials('dockerhub-sumitroajiprabowo')
            }
            parallel {
                stage ('Preparation Golang') {
                    agent {
                        node {
                            label 'golang'
                        }
                    }
                    steps {
                        sh 'go version'
                        sh 'go env'
                        sleep 5
                    }
                }

                stage ('Preparation Environment') {
                    agent {
                        node {
                            label 'golang'
                        }
                    }
                    steps {
                        echo ("Author, ${env.AUTHOR}")
                        echo ("Email, ${env.EMAIL}")
                        echo ("Docker Hub Username, ${env.APP_USR}")
                        echo ("Docker Hub Password, ${env.APP_PSW}")
                        echo ("Start Job ${env.JOB_NAME}")
                        echo ("Start Build ${env.BUILD_NUMBER}")
                        echo ("Start Branch ${env.BRANCH_NAME}")
                        sleep 5
                    }
                }
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
                    choice(name: 'TARGET_ENVIRONMENT', choices: ['DEV', 'STAGING', 'PRODUCTION'], description: 'Choose the target environment')
                }
            }
            agent {
                node {
                        label 'golang'
                }
            }
            steps {
                echo("Deploy to ${TARGET_ENVIRONMENT}")
            }
        }

        stage('Release') {
            when {
                expression {
                    return params.DEPLOY;
                }
            }
            agent {
                node {
                        label 'golang'
                }
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-sumitroajiprabowo', usernameVariable: 'DOCKERHUB_USERNAME', passwordVariable: 'DOCKERHUB_PASSWORD')]) {
                    //sh 'docker login -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD'
                    //sh 'docker build -t sumitroajiprabowo/jenkins-pipeline-golang:latest .'
                    //sh 'docker push sumitroajiprabowo/jenkins-pipeline-golang:latest'
                    sh('echo "Release with -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD"')
                }
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