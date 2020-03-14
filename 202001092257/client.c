#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define UN_PATH "/var/tmp/test_domain_socket"


int main(int argc, char *argv[])
{
    int sockfd = -1;
    struct sockaddr_un un;
    int len = 0;

    if ((sockfd = socket(AF_UNIX, SOCK_STREAM, 0)) < 0){
        perror("create socket error\n");
        goto _ERROR;
    }

    memset(&un, 0, sizeof(un));
    un.sun_family = AF_UNIX;
    strcpy(un.sun_path, UN_PATH);

    len = sizeof(un.sun_family) + sizeof(un.sun_path);

    if (connect(sockfd, (struct sockaddr *)&un, len) < 0){
        perror("socket connect error\n");
        goto _ERROR;
    }

    char buf[64] = {0};
    int n = read(STDIN_FILENO, buf, 63);
    write(sockfd, buf, n);
    printf("send-msg:%s", buf);
    read(sockfd, buf, 63);
    buf[64] = '\0';
    printf("recv-msg:%s", buf);

    close(sockfd);
    return 0;

_ERROR:
    if (sockfd != -1){
        close(sockfd);
    }

    return -1;
}