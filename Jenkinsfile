pipeline {
    agent any

    environment {
        TAG = "fantlab/go:go_build_$BUILD_NUMBER"
        IMAGE = ''
    }

    stages {
        stage('Build') {
            steps {
                script {
                    IMAGE = docker.build TAG
                }
            }
        }
        stage('Push') {
            steps {
                script {
                    docker.withRegistry('', 'fldockerhub') {
                        IMAGE.push()
                        IMAGE.push('latest')
                    }
                }
            }
        }
        stage('Clean') {
            steps{
                sh "docker image prune -f"
                sh "docker container prune -f"
            }
        }
        stage('Deploy') {
            steps{
                sh "docker stack deploy --prune --compose-file deploy.yml go-api"
            }
        }
    }
}
