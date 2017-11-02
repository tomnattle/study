#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string.h>
#include <signal.h>
#include <sys/epoll.h>
#include <errno.h>

#define SUB_PROCCESS_COUNT 2
#define SERVER_PORT 8000
#define MAX_EVENTS 256

void createSocket(int *_socket_id){
    int socket_id;
    struct sockaddr_in server_addr;
    int sin_size;
    
    memset(&server_addr, 0, sizeof(server_addr));
    // IP 通信
    server_addr.sin_family = AF_INET;
    // 允许接受本地人和主机的请求
    server_addr.sin_addr.s_addr = INADDR_ANY;
    server_addr.sin_port = htons(SERVER_PORT);

    // 创建一个socket文件句柄
    socket_id = socket(AF_INET, SOCK_STREAM, 0);
    // printf("%d\n", socket_id);
    if (socket_id < 0){
        perror("socket create failure.\n");
    }
    printf("create socket handel success.\n");
    int opt = 1;
    //  端口重用
    setsockopt(socket_id, SOL_SOCKET, SO_REUSEADDR,(const void *)&opt , sizeof(opt));
    // 绑定端口
    int bind_r = bind(socket_id, (struct sockaddr *)&server_addr, sizeof(struct sockaddr));
    if (bind_r < 0)
        perror("bind port failure.");
    printf("bind port success.\n");

    if(listen(socket_id, 5) < 1){
        perror("listen failure");
    }
    printf("listen success, len 5.\n");

    sin_size = sizeof(struct sockaddr_in);
    struct sockaddr_in remote_addr;
    int client_id;
    // epoll
    int epoll_fd;
    // 创建一个epoll事件池子
    epoll_fd = epoll_create(MAX_EVENTS);
    if(epoll_fd < 0){
        perror("epoll_create failed");
        exit(EXIT_FAILURE);
    }

    struct epoll_event ev;
    struct epoll_event events[MAX_EVENTS];
    // 当有数据进入时出发
    ev.events = EPOLLIN|EPOLLET;
    ev.data.fd = socket_id;

    epoll_ctl(epoll_fd, EPOLL_CTL_ADD, socket_id, &ev);
    char buf[BUFSIZ];
    int nfds;
    while(1) {
        //等待事件， 如果有，则放在一个列表中
        nfds = epoll_wait(epoll_fd, events, MAX_EVENTS, -1);
        //printf("epoll events triggled.\n");
        int i, len;
        //遍历整个列表
        for(i = 0; i < nfds; i++) {
            // 如果数据的id 是连接id 则是新增的连接
            int fd = events[i].data.fd;
            // printf("epoll_fd%d, socket_id %d\n", fd, socket_id);
            if ((fd == socket_id) &&(events[i].events & EPOLLIN)){
                client_id = accept(socket_id, (struct sockaddr*)&remote_addr, (socklen_t*)&sin_size);
                ev.data.fd = client_id;
                ev.events = EPOLLIN | EPOLLET;
                printf("new connection");
                printf("ip is %s\n", inet_ntoa(remote_addr.sin_addr));
                printf("port is %d\n", htons(remote_addr.sin_port));
                epoll_ctl(epoll_fd, EPOLL_CTL_ADD, client_id, &ev);
                //读取
            // 如果事件是in事件，则是数据接收事件
            }else if(events[i].events & EPOLLIN){
                // printf("accept data\n");
                if ((client_id = events[i].data.fd) < 0)
                    continue;
                if ((len = read(client_id, buf, BUFSIZ)) < 0){
                    // 删除事件
                    epoll_ctl(epoll_fd, EPOLL_CTL_DEL, client_id, &ev);
                } else if( len == 0){
                    printf("client colse.\n");
                    close(client_id);
                    epoll_ctl(epoll_fd, EPOLL_CTL_DEL, client_id, &ev);
                } else {
                    buf[len-1] = '\0';
                    buf[len-2] = '\0';
                    //printf("[%d]. data [%s], len %d, last two letter[%d %d]\n", client_id, buf, len, buf[len-2], buf[len-1]);
                    printf("[%d]. data [%s], len %d\n", client_id, buf, len);
                    if(strcmp(buf, "quit\0\0") == 0){
                        close(client_id);
                        //exit(1);
                        break;
                    }
                    buf[len-1] = '\n';
                    buf[len-2] = '\r';
                    ev.data.fd = client_id;
                    ev.events = EPOLLOUT | EPOLLET;
                    // 修改事件 这个操作会出发上句中的 out句柄
                    epoll_ctl(epoll_fd, EPOLL_CTL_MOD, client_id, &ev);
                }
            }else if(events[i].events & EPOLLOUT){
                //client_id = events[i].data.fd;
                write(client_id, buf, len);
                ev.data.fd = client_id;
                ev.events = EPOLLIN | EPOLLET;
                epoll_ctl(epoll_fd, EPOLL_CTL_MOD, client_id, &ev);
            }
        }
    }
}

void sigHandel(int sig){
    switch(sig) {
        case 2:
            printf("ctr+c press, programe exit.\n");
            exit(0);
        default:
            printf("signal is :%d\n", sig);
    }
}

int main(int argc, char* argv[]){
    signal(SIGINT, sigHandel);
    pid_t pid;
    int socket_id;
    createSocket(&socket_id);
    printf("socket id is %d\n", socket_id);
    //printf("this is parent process, %d, exit.\n", getpid());
    for (int i = 0; i < SUB_PROCCESS_COUNT; ++i)
    {
        pid = fork();
        printf("%d\n", pid);
        // 如果在子进程中 或者是出错的情况下 不继续fork
        if( 0 == pid || -1 == pid )  
            break;
    }

    if(pid == -1) {
        printf("error happend");
        exit(EXIT_FAILURE);
    }

    if (pid != 0){
        printf("this is parent process, %d, exit.\n", getpid());
        exit(0);
    }else{
        if (setsid() == -1)
            exit(EXIT_FAILURE);
        chdir("/");
        for (int i = 0; i < 3; i++){
            close(i);
            open("/dev/null", O_RDWR);
            dup(0);
            dup(0);
        }
        umask(0);
        sleep(1);
        printf("this is child process.\n");
        while(1){
            printf("%d\n", getpid());
            sleep(1);
        }
    }
    
    return 0;
}


    /**
    while(1){
        // 接受数据
        printf("wait for accept client connect\n");
        client_id = accept(socket_id, (struct sockaddr*)&remote_addr, (socklen_t*)&sin_size);
        //printf("%d\n", client_id);
        if (client_id < 0){
            //perror("accept error");
            continue;
        }
        //exit(1);
        printf("ip is %s\n", inet_ntoa(remote_addr.sin_addr));
        printf("port is %d\n", htons(remote_addr.sin_port));
        while(1){
            char buf[BUFSIZ];
            len = recv(client_id, buf, BUFSIZ, 0);
            if (len < 0) {
                perror("recv error");
                break;
            }
            
            buf[len] = '\0';
            if(strcmp(buf, "quit\r\n") == 0){
                printf("quit\n");
                close(client_id);
                break;
            }else{
                printf("[%d]recv data is %s\n", len, buf);
                //printf("%d\n", strcmp(buf, "quit"));
                //printf("%lu\n", strlen(buf));
            }
            //send(client_id, buf, len, 0);
        }
    }
    *_socket_id = socket_id;
    **/
